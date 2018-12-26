# utils
Utilities Library For the Habitat Project, it includes:

# Yaml Parser

A Yaml file Parser and Yaml File creator that creates an easy to use tree model out of .yaml file and allows you to create a tree model of yaml nodes that is persisted to .yaml file.

## Yaml Parser Example

    package examples
    
    import (
    	. "github.com/saichler/utils/golang"
    	"io/ioutil"
    	"os"
    )

    func Parse_yaml_file(yamlfile string) (*YamlNode, error){
        _,e := os.Stat(yamlfile)
        if e!=nil {
            return nil,e
        }

        buff,e := ioutil.ReadFile(yamlfile)
        if e!=nil {
            return nil,e
        }

        data := string(buff)
        modelRoot:=NewYamlRoot()
        Parse(data,yamlfile,modelRoot)
        return modelRoot,nil
    }

    func Create_yaml_file(yamlfile string) {
        modelRoot:=NewYamlRoot()
        child:=&YamlNode{}
        child.Init("tag","value",modelRoot.GetLvl()+1)
        modelRoot.AddChild(child)
        ioutil.WriteFile(yamlfile,[]byte(modelRoot.String()),0777)
    }

# SqlST

An slq.ST wrapper that makes it easy to create sql parepare statements to avoid sql injection via concatenating sql queries.

## SqlST Example

    package examples

    import (
        "database/sql"
        "fmt"
        "github.com/saichler/utils/golang"
        log "github.com/sirupsen/logrus"
    )

    func insertToTableExample(tx *sql.Tx) {
        st:=dsutils.CreateInsertStatement("MyTable")
        st.AddColumn("Name","MyName")
        st.AddColumn("FamilyName","My Family Name")
        st.AddColumn("Age",5)
        _,err:=st.Exec(tx)
        if err!=nil{
            log.Error(err)
        }
    }

    func selectFromTableExample(tx *sql.Tx) {
        st:=dsutils.CreateSelectStatement("MyTable","#1 AND #2")
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
    
    # LineReader
    # Syncronized Queue
    # ByteAray



