package dsutils

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	STATEMENT_TYPE_INSERT = 1
	STATEMENT_TYPE_UPDATE = 2
	STATEMENT_TYPE_DELETE = 3
	STATEMENT_TYPE_SELECT = 4;
	PARAM_CHAR="#"
	VALUE_CHAR="%"
	STATEMENT_CHAR="$"
	CACHE_ENTRY_TIMEOUT=30
	CACHE_CLEAN_TIMEOUT=300
)

type SqlST struct {
	statementType int
	tableName     string
	parameters    map[int]string
	values        []interface{}
	cValues       []interface{}
	cParameters   map[int]string
	statement     *sql.Stmt
	criteria      *StringBuilder
}

type TxCacheEntry struct {
	timeStamp int64
	statements map[string]*sql.Stmt
}

var lastCacheSwipe = time.Now().Unix()
var txCache = make(map[*sql.Tx]*TxCacheEntry)
var txMutex = &sync.Mutex{}

func CreateInsertStatement(tableName string) *SqlST {
	s := createStatement(tableName,"")
	s.statementType = STATEMENT_TYPE_INSERT
	log.Debug("Created an Insert SqlST for table ", tableName)
	return s
}

func CreateUpdateStatement(tableName string, criteria string) *SqlST {
	s := createStatement(tableName,criteria)
	s.statementType = STATEMENT_TYPE_UPDATE
	log.Debug("Created an Update SqlST for table ", tableName)
	return s
}

func CreateDeleteStatement(tableName string, criteria string) *SqlST {
	s := createStatement(tableName,criteria)
	s.statementType = STATEMENT_TYPE_DELETE
	log.Debug("Created a Delete SqlST for table ", tableName)
	return s
}

func CreateSelectStatement(tableName string, criteria string) *SqlST {
	s := createStatement(tableName,criteria)
	s.statementType = STATEMENT_TYPE_SELECT
	log.Debug("Created a Select SqlST for table ", tableName)
	return s
}

func createStatement(tableName string, criteria string) *SqlST {
	s := SqlST{}
	s.tableName = tableName
	s.values = make([]interface{},0)
	s.parameters = make(map[int]string)
	s.cValues = make([]interface{},0)
	s.cParameters = make(map[int]string)
	s.criteria = NewStringBuilder(criteria)
	return &s
}

func (s *SqlST) AddColumn(name string, value interface{}){
	index := len(s.values)
	s.values = append(s.values, value)
	s.parameters[index] = name
}

func (s *SqlST) AddCriteriaColumn(name string, value interface{}){
	index := len(s.cValues)
	s.cValues = append(s.cValues, value)
	s.cParameters[index] = name
}

func putStatementInCache(tx *sql.Tx,stmt *sql.Stmt,md5key string) {
	txEntry:=txCache[tx]
	txEntry.statements[md5key] = stmt
}

func cleanCache(){
	now:=time.Now().Unix()
	if now-lastCacheSwipe>= CACHE_CLEAN_TIMEOUT {
		log.Info("***** SqlST Cache being cleaned...")
		lastCacheSwipe = now
		entriesToRemove := make([]*sql.Tx, 0)
		for k, v := range txCache {
			if now-v.timeStamp >= CACHE_ENTRY_TIMEOUT {
				entriesToRemove = append(entriesToRemove, k)
			}
		}
		for _, k := range entriesToRemove {
			delete(txCache, k)
		}
		log.Info("***** SqlST Cache was cleaned, "+strconv.Itoa(len(entriesToRemove))+" were removed.")
	}
}

func getCachedStatement(tx *sql.Tx,sqlQuery string) (*sql.Stmt,string) {
	cleanCache()
	md5:=md5.Sum([]byte(sqlQuery))
	md5Key:= base64.StdEncoding.EncodeToString(md5[:])
	txEntry:=txCache[tx]
	if txEntry==nil {
		txEntry = &TxCacheEntry{}
		txEntry.statements = make(map[string]*sql.Stmt)
		txCache[tx]=txEntry
	}
	return txEntry.statements[md5Key],md5Key
}

func (s *SqlST) prepareStatement(db_tx *sql.Tx, sqlQuery string) error{

	defer txMutex.Unlock()
	txMutex.Lock()

	statement,md5key:=getCachedStatement(db_tx,sqlQuery)
	if statement!=nil {
		s.statement = statement
		return nil
	}

	log.Debug("Going to execute prepare statement for the following sql:",sqlQuery)
	newStatement, err := db_tx.Prepare(sqlQuery)
	s.statement = newStatement

	if err==nil {
		putStatementInCache(db_tx,newStatement,md5key)
	}
	return err
}

func (s *SqlST) executeStatement(db_tx *sql.Tx) (sql.Result, error) {
	m := reflect.ValueOf(s.statement).MethodByName("Exec")
	args:=make([]reflect.Value,0)
	subStatements:=make([]*SqlST,0)
	for i:=0;i<len(s.values);i++ {
		args = append(args,reflect.ValueOf(s.values[i]))
	}

	for i:=0;i<len(s.cValues);i++ {
		cst, isStatement := s.cValues[i].(*SqlST)
		if isStatement {
			subStatements = append(subStatements,cst)
		} else {
			args = append(args,reflect.ValueOf(s.cValues[i]))
		}
	}

	for i:=0;i<len(subStatements);i++ {
		for j:=0;j<len(subStatements[i].cValues);j++ {
			args = append(args,reflect.ValueOf(subStatements[i].cValues[j]))
		}
	}

	value :=  m.Call(args)
	if !value[1].IsNil() {
		return nil, value[1].Interface().(error)
	}
	return value[0].Interface().(sql.Result),nil
}

func (s *SqlST) getSubSql(subStatement *SqlST) *StringBuilder {
	sql := subStatement.BuildSelectStatement().String()
	subArgsSize := len(subStatement.cValues)
	offset := 0
	if s.statementType==STATEMENT_TYPE_SELECT {
		offset = len(s.cValues)-1
	} else {
		offset = len(s.cValues)+len(s.values)-1
	}
	for i:=1;i<=subArgsSize;i++ {
		dollarParam := formatParam(STATEMENT_CHAR,i)
		dollarIndex := strings.Index(sql,dollarParam)
		if dollarIndex!=-1 {
			buff := bytes.Buffer{}
			buff.WriteString(sql[0:dollarIndex])
			buff.WriteString(STATEMENT_CHAR)
			buff.WriteString(strconv.Itoa(i+offset))
			buff.WriteString(sql[dollarIndex+2:])
			sql = buff.String()
		}
	}
	return NewStringBuilder(sql)
}

func (s *SqlST) queryStatement(db_tx *sql.Tx) (*sql.Rows, error) {
	m := reflect.ValueOf(s.statement).MethodByName("Query")
	args:=make([]reflect.Value,0)
	subStatements := make([]*SqlST,0)
	for i:=0;i<len(s.cValues);i++ {
		cst, isStatement := interface{}(s.cValues[i]).(*SqlST)
		if isStatement {
			subStatements = append(subStatements,cst)
		} else {
			args = append(args,reflect.ValueOf(s.cValues[i]))
		}
	}

	for i:=0;i<len(subStatements);i++ {
		for j:=0;j<len(subStatements[i].cValues);j++ {
			args = append(args,reflect.ValueOf(subStatements[i].cValues[j]))
		}
	}

	value :=  m.Call(args)
	if !value[1].IsNil() {
		return nil, value[1].Interface().(error)
	}
	return value[0].Interface().(*sql.Rows),nil
}

func (s *SqlST) queryRowStatement(db_tx *sql.Tx) (*sql.Row) {
	m := reflect.ValueOf(s.statement).MethodByName("QueryRow")
	args:=make([]reflect.Value,len(s.cValues))
	for i:=0;i<len(s.cValues);i++ {
		args[i] = reflect.ValueOf(s.cValues[i])
	}
	value :=  m.Call(args)
	return value[0].Interface().(*sql.Row)
}

func (s *SqlST) BuildInsertStatement() *StringBuilder {
	sql:=NewStringBuilder("INSERT INTO ").Append(s.tableName).Append(" ")

	fields:=NewStringBuilder("(")
	values:=NewStringBuilder(" VALUES (")

	for i:=0;i<len(s.values);i++ {
		if i!=0 {
			fields.Append(",")
			values.Append(",")
		}
		fields.Append(s.parameters[i])
		values.Append(STATEMENT_CHAR)
		values.Append(strconv.Itoa(i+1))
	}
	fields.Append(")")
	values.Append(")")
	sql.AppendSB(fields)
	sql.AppendSB(values)
	sql.Append(";")

	return sql
}

func (s *SqlST) pos(i int) int {
	if s.statementType==STATEMENT_TYPE_SELECT {
		return i+1
	}
	return i+1+len(s.values)
}

func (s *SqlST) posStr(i int) string {
	return strconv.Itoa(s.pos(i))
}

func formatParam(char string,index int) string {
	param := bytes.Buffer{}
	param.WriteString(char)
	param.WriteString(strconv.Itoa(index))
	return param.String()
}

func (s *SqlST) insertCriteria() *StringBuilder {
	where := NewStringBuilder("")
	if !s.criteria.Empty() {
		where.Append(" WHERE ")
		cr := s.criteria.String()

		for i:=0;i<len(s.cValues);i++ {
			dollarParam := formatParam(PARAM_CHAR,i+1)
			percentParam := formatParam(VALUE_CHAR,i+1)
			dollarIndex := strings.Index(cr,dollarParam)
			percentIndex := strings.Index(cr,percentParam)
			if dollarIndex!=-1 {
				if percentIndex == -1 {
					buff:=NewStringBuilder(cr[0:dollarIndex])
					buff.Append(s.cParameters[i]).Append("=").Append(STATEMENT_CHAR)
					buff.Append(s.posStr(i)).Append(cr[dollarIndex+2:])
					cr =  buff.String()
				} else {
					buff:=NewStringBuilder(cr[0:dollarIndex])
					buff.Append(s.cParameters[i])
					buff.Append(cr[dollarIndex+2:])
					cr =  buff.String()
					percentIndex := strings.Index(cr,percentParam)
					cst, isStatement := interface{}(s.cValues[i]).(*SqlST)
					if isStatement {
						buff:=NewStringBuilder(cr[0:percentIndex])
						buff.AppendSB(s.getSubSql(cst))
						buff.Append(cr[percentIndex+2:])
						cr = buff.String()
					} else {
						buff:=NewStringBuilder(cr[0:percentIndex])
						buff.Append(STATEMENT_CHAR)
						buff.Append(s.posStr(i))
						buff.Append(cr[percentIndex+2:])
						cr = buff.String()
					}
				}
			}
		}
		where.Append(cr)
	}
	return where
}

func (s *SqlST) BuildUpdateStatement() *StringBuilder {
	sql:=NewStringBuilder("UPDATE ")
	sql.Append(s.tableName).Append(" SET ")
	for i:=0;i<len(s.values);i++ {
		if i!=0 {
			sql.Append(",")
		}
		sql.Append(s.parameters[i]).Append("=").Append(STATEMENT_CHAR).Append(strconv.Itoa(i+1))
	}
	sql.AppendSB(s.insertCriteria())
	return sql
}

func (s *SqlST) BuildDeleteStatement() *StringBuilder {
	sql:=NewStringBuilder("DELETE FROM ").Append(s.tableName).Append(" ")
	sql.AppendSB(s.insertCriteria())
	return sql
}

func (s *SqlST) BuildSelectStatement() *StringBuilder {
	sql:=NewStringBuilder("SELECT ")
	for i:=0;i<len(s.values);i++ {
		if i==0 {
			sql.Append(s.parameters[i])
		} else {
			sql.Append(",")
			sql.Append(s.parameters[i])
		}
	}
	sql.Append(" FROM ")
	sql.Append(s.tableName)
	sql.AppendSB(s.insertCriteria())
	return sql
}

func (st *SqlST)Exec(db_tx *sql.Tx) (sql.Result, error) {
	var sql *StringBuilder
	if st.statementType==STATEMENT_TYPE_INSERT {
		sql = st.BuildInsertStatement()
	} else if st.statementType==STATEMENT_TYPE_UPDATE {
		sql = st.BuildUpdateStatement()
	} else if st.statementType==STATEMENT_TYPE_DELETE {
		sql = st.BuildDeleteStatement()
	}

	err := st.prepareStatement(db_tx, sql.String())
	if err!=nil {
		log.Error("Failed to prepare statement due to the following error:", err)
		return nil,err
	}
	log.Debug("Prepare statement was successful, going to execute the statement")
	result,err := st.executeStatement(db_tx)
	if err!=nil {
		log.Error("Failed to execute statement due to the following error:", err)
		panic("p")
		return nil,err
	}
	log.Debug("SqlST was executed successfully!")
	return result,nil
}

func (st *SqlST)Query(db_tx *sql.Tx) (*sql.Rows, error) {
	sql := st.BuildSelectStatement()
	err := st.prepareStatement(db_tx, sql.String())
	if err!=nil {
		log.Error("Failed to prepare statement due to the following error:", err)
		return nil,err
	}
	log.Debug("Prepare statement was successful, going to execute the statement")
	result,err := st.queryStatement(db_tx)
	if err!=nil {
		log.Error("Failed to query statement due to the following error:", err)
		return nil,err
	}
	log.Debug("SqlST was executed successfully!")
	return result,nil
}

func (st *SqlST)QueryRow(db_tx *sql.Tx) (*sql.Row) {
	sql := st.BuildSelectStatement()
	err := st.prepareStatement(db_tx, sql.String())
	if err!=nil {
		log.Error("Failed to prepare statement due to the following error:", err)
		return nil
	}
	log.Debug("Prepare statement was successful, going to execute the statement")
	return st.queryRowStatement(db_tx)
}

func (st *SqlST) GetNextCriteriaParamIndex() *StringBuilder {
	result:=NewStringBuilder(" ")
	result.Append(STATEMENT_CHAR).Append(strconv.Itoa(len(st.cValues))).Append(" ")
	return result
}

func (st *SqlST) GetNextCriteriaValueIndex() *StringBuilder {
	result:=NewStringBuilder(" ")
	result.Append(VALUE_CHAR).Append(strconv.Itoa(len(st.cValues))).Append(" ")
	return result
}

func (st *SqlST) AppendCriteria(name string,value interface{}, ctr string) {
	index := len(st.cValues)+1
	st.AddCriteriaColumn(name, value)
	param := strings.Index(ctr,PARAM_CHAR+"1")
	if param!=-1 {
		st.criteria.Append(ctr[0:param]).Append(PARAM_CHAR).Append(strconv.Itoa(index)).Append(ctr[param+2:])
	}
	param = strings.Index(ctr,VALUE_CHAR+"1")
	if param!=-1 {
		st.criteria.Append(ctr[0:param]).Append(VALUE_CHAR).Append(strconv.Itoa(index)).Append(ctr[param+2:])
	}
}

func (st *SqlST) AppendCriteriaNoArgs(ctr *StringBuilder) {
	st.criteria.AppendSB(ctr)
}

func (st *SqlST) MarshalCriteria(size int){
	cr:=st.criteria.String()
	st.criteria = NewStringBuilder(cr[0:len(cr)-size])
}