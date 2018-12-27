package examples

import (
	"database/sql"
	"fmt"
	"github.com/saichler/utils/golang"
	log "github.com/sirupsen/logrus"
)

func insertToTableExample(tx *sql.Tx) {
	st:= utils.CreateInsertStatement("MyTable")
	st.AddColumn("Name","MyName")
	st.AddColumn("FamilyName","My Family Name")
	st.AddColumn("Age",5)
	_,err:=st.Exec(tx)
	if err!=nil{
		log.Error(err)
	}
}

func selectFromTableExample(tx *sql.Tx) {
	st:= utils.CreateSelectStatement("MyTable","#1 AND #2")
	st.AddColumn("Name","")
	st.AddCriteriaColumn("FamilyName","My Family Name")
	st.AddCriteriaColumn("Age",5)
	rows,err:=st.Query(tx)
	if err!=nil {
		log.Error(err)
	}
	for ;rows.Next(); {
		name:=""
		rows.Scan("Name",name)
		fmt.Println(name)
	}
}
