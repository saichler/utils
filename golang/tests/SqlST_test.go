package dbutil

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	. "github.com/saichler/utils/golang"
	"strings"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "test"
	password = "test"
	dbname   = "test"
	table    = "test_table"
	col1     = "column1"
	col2     = "column2"
)

func startTransaction() *sql.Tx {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	CheckError(err)
	dbcreate(db)
	tx, err := db.Begin()
	return tx
}

func count() int {
	tx := startTransaction()
	st := CreateSelectStatement(table,"")
	st.AddColumn("count(*) as cnt","")
	row :=st.QueryRow(tx)
	var cnt int
	err := row.Scan(&cnt)
	if err!=nil {
		fmt.Println(err)
	}
	return cnt
}

func assert(x int, m string, t *testing.T){
	cnt := count()
	if cnt==x {
		t.Logf("Test of "+m+" passed")
	} else {
		t.Errorf("Test of "+m+" failed")
	}
}

func deleteAll(){
	tx:=startTransaction()
	st := CreateDeleteStatement(table,"")
	st.Exec(tx)
	tx.Commit()
}

func insertInto(c1,c2 string) {
	tx:=startTransaction()
	st := CreateInsertStatement(table)
	st.AddColumn(col1,c1)
	st.AddColumn(col2, c2)
	st.Exec(tx)
	tx.Commit()
}

func TestDeleteAllStatement (t *testing.T) {
	deleteAll()
	assert(0,"DELETE ALL",t)
}

func TestInsertStatement(t *testing.T) {
	deleteAll()
	insertInto("Hello","World")
	assert(1,"INSERT INTO",t)
}

func TestQueryRaw(t *testing.T) {
	deleteAll()
	val := "Hello"
	insertInto(val,"World")
	s:=CreateSelectStatement(table,"#1")
	s.AddColumn(col1,"")
	s.AddCriteriaColumn(col1,val)
	tx := startTransaction()
	row := s.QueryRow(tx)
	name := ""
	row.Scan(&name)
	if strings.Trim(name," ")==val {
		t.Logf("Test of Query Row passed")
	} else {
		t.Errorf("Test of Query Row failed")
	}
}

func TestQueryRaw2(t *testing.T) {
	deleteAll()
	val := "Hello"
	insertInto(val,"World")
	s:=CreateSelectStatement(table,"#1=%1")
	s.AddColumn(col1,"")
	s.AddCriteriaColumn(col1,val)
	tx := startTransaction()
	row := s.QueryRow(tx)
	name := ""
	row.Scan(&name)
	if strings.Trim(name," ")==val {
		t.Logf("Test of Query Row passed")
	} else {
		t.Errorf("Test of Query Row failed")
	}
}

func TestSelect(t *testing.T) {
	deleteAll()
	val := "Hello"
	insertInto(val,"World")
	insertInto("1","1")
	insertInto("2","2")
	s:=CreateSelectStatement(table,"#1")
	s.AddColumn(col1,"")
	s.AddCriteriaColumn(col1,val)
	tx := startTransaction()
	rows, err := s.Query(tx)
	if err!=nil {
		t.Errorf("Test of Select failed")
	}
	rows.Next()
	name := ""
	rows.Scan(&name)
	if strings.Trim(name," ")==val {
		t.Logf("Test of Select Row passed")
	} else {
		t.Errorf("Test of Select Row failed")
	}
}

func TestSelect2(t *testing.T) {
	deleteAll()
	val := "Hello"
	insertInto(val,"World")
	insertInto("1","1")
	insertInto("2","2")
	s:=CreateSelectStatement(table,"#1=%1")
	s.AddColumn(col1,"")
	s.AddCriteriaColumn(col1,val)
	tx := startTransaction()
	rows, err := s.Query(tx)
	if err!=nil {
		t.Errorf("Test of Select failed")
	}
	rows.Next()
	name := ""
	rows.Scan(&name)
	if strings.Trim(name," ")==val {
		t.Logf("Test of Select Row passed")
	} else {
		t.Errorf("Test of Select Row failed")
	}
}

func TestUpdate(t *testing.T) {
	deleteAll()
	oldVal := "Hello"
	newVal := "NewValue"
	insertInto(oldVal,"World")
	insertInto("1","1")
	insertInto("2","2")
	s:=CreateUpdateStatement(table,"#1")
	s.AddColumn(col1,newVal)
	s.AddCriteriaColumn(col1,oldVal)
	tx := startTransaction()
	_,err := s.Exec(tx)
	if err!=nil {
		t.Errorf("Test of Update table failed")
	}
	tx.Commit()

	s=CreateSelectStatement(table,"#1")
	s.AddColumn(col1,"")
	s.AddCriteriaColumn(col1,newVal)
	tx = startTransaction()
	row := s.QueryRow(tx)

	name := ""
	row.Scan(&name)
	if strings.Trim(name," ")==newVal {
		t.Logf("Test of update Row passed")
	} else {
		t.Errorf("Test of update Row failed")
	}
}

func TestUpdate2(t *testing.T) {
	deleteAll()
	oldVal := "Hello"
	newVal := "NewValue"
	insertInto(oldVal,"World")
	insertInto("1","1")
	insertInto("2","2")
	s:=CreateUpdateStatement(table,"#1=%1")
	s.AddColumn(col1,newVal)
	s.AddCriteriaColumn(col1,oldVal)
	tx := startTransaction()
	_,err := s.Exec(tx)
	if err!=nil {
		t.Errorf("Test of Update table failed")
	}
	tx.Commit()

	s=CreateSelectStatement(table,"#1")
	s.AddColumn(col1,"")
	s.AddCriteriaColumn(col1,newVal)
	tx = startTransaction()
	row := s.QueryRow(tx)

	name := ""
	row.Scan(&name)
	if strings.Trim(name," ")==newVal {
		t.Logf("Test of update Row passed")
	} else {
		t.Errorf("Test of update Row failed")
	}
}

func TestSelectInsideSelect(t *testing.T) {
	deleteAll()
	val := "Hello"
	insertInto(val,"World")
	insertInto("1","1")
	insertInto("2","2")
	s1:=CreateSelectStatement(table,"#1=%1")
	s1.AddColumn(col1,"")
	s1.AddCriteriaColumn(col1,val)

	s2:=CreateSelectStatement(table,"#1 AND #2 IN (%2)")
	s2.AddColumn(col1,"")
	s2.AddCriteriaColumn(col2,"World")
	s2.AddCriteriaColumn(col1,s1)

	tx := startTransaction()
	rows, err := s2.Query(tx)
	if err!=nil {
		t.Errorf("Test of Select failed")
	}
	rows.Next()
	name := ""
	rows.Scan(&name)
	if strings.Trim(name," ")==val {
		t.Logf("Test of Select inside select Row passed")
	} else {
		t.Errorf("Test of Select inside select Row failed")
	}
}

func TestSelectInsideUpdate(t *testing.T) {
	deleteAll()
	val := "Hello"
	newVal := "NewValue"
	insertInto(val,"World")
	insertInto("1","1")
	insertInto("2","2")

	s1:=CreateSelectStatement(table,"#1")
	s1.AddColumn(col1,"")
	s1.AddCriteriaColumn(col1,val)

	s2:=CreateUpdateStatement(table,"#1 AND #2 IN (%2)")
	s2.AddColumn(col1,newVal)
	s2.AddCriteriaColumn(col2,"World")
	s2.AddCriteriaColumn(col1,s1)

	tx := startTransaction()
	_, err := s2.Exec(tx)
	if err!=nil {
		t.Errorf("Test of Update failed")
	}
	tx.Commit()

	s3:=CreateSelectStatement(table,"#1")
	s3.AddColumn(col1,"")
	s3.AddCriteriaColumn(col1,newVal)
	tx = startTransaction()
	row := s3.QueryRow(tx)

	name := ""
	row.Scan(&name)
	if strings.Trim(name," ")==newVal {
		t.Logf("Test of Select inside Update Row passed")
	} else {
		t.Errorf("Test of Select inside Update Row failed")
	}
}

func dbcreate(db *sql.DB){
	db.Exec("CREATE TABLE if NOT EXISTS "+table+" ("+col1+" char(15),"+col2+" char(15))")
}

func CheckError(err error){
	if err!=nil{
		fmt.Println("Error: ", err)
	}
}

