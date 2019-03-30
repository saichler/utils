package tests

import (
	. "github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/marshal"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"strconv"
	"testing"
)

var size = 5;

func initTest(numOfNodes int) *Transaction {
	registry := &OrmRegistry{}
	registry.Register(Node{})
	nodes:=InitTestModel(numOfNodes)
	tx:=&Transaction{}
	Marshal(nodes,registry,tx,nil)
	return tx
}

func findNodeRecords(records []*Record, id int) *Record {
	for _,nr:=range records {
		if nr.Map()["String"].String()=="String-"+strconv.Itoa(id) {
			return nr
		}
	}
	return nil
}

func TestMarshalString(t *testing.T) {
	tx:=initTest(5)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
	}
}

func TestMarshalInt(t *testing.T) {
	tx:=initTest(size)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
		expected:=int64(-101+i)
		val:=rec.Map()["Int"].Int()
		if val!= expected{
			t.Fail()
			Error("Expected "+strconv.Itoa(int(expected))+" but got:"+strconv.Itoa(int(val)))
		}
	}
}

func TestMarshalInt32(t *testing.T) {
	tx:=initTest(size)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
		expected:=int64(-102+i)
		val:=rec.Map()["Int32"].Int()
		if val!= expected{
			t.Fail()
			Error("Expected "+strconv.Itoa(int(expected))+" but got:"+strconv.Itoa(int(val)))
		}
	}
}

func TestMarshalInt64(t *testing.T) {
	tx:=initTest(size)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
		expected:=int64(-103+i)
		val:=rec.Map()["Int64"].Int()
		if val!= expected{
			t.Fail()
			Error("Expected "+strconv.Itoa(int(expected))+" but got:"+strconv.Itoa(int(val)))
		}
	}
}

func TestMarshalBool(t *testing.T) {
	tx:=initTest(size)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
		expected:=true
		val:=rec.Map()["Bool"].Bool()
		if val!= expected{
			t.Fail()
			Error("Expected true but got false")
		}
	}
}

func TestMarshalPtrKey(t *testing.T) {
	tx:=initTest(size)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
		expected:="OnlyChild-String-"+strconv.Itoa(i)
		val:=rec.Map()["Ptr"].String()
		if val!=expected {
			t.Fail()
			Error("Expected:"+expected+" got:"+val)
		}
	}
}

func TestMarshalPtrNoKey(t *testing.T) {
	tx:=initTest(size)
	nodeRecords:=tx.Records()["Node"]
	for i:=0;i<size;i++ {
		rec := findNodeRecords(nodeRecords, i)
		if rec == nil {
			t.Fail()
			Error("No Recrod was found with id:"+strconv.Itoa(i))
		}
		val:=rec.Map()["PtrNoKey"]
		if val.IsValid() {
			t.Fail()
			Error("Expected not valid but got valid")
		}
	}
}

func TestMarshalNumberOfRecords(t *testing.T) {
	tx:=initTest(5)
	nodeRecords:=tx.Records()["Node"]
	if len(nodeRecords)!=30 {
		t.Fail()
		Error("Expected 30 but got:"+strconv.Itoa(len(nodeRecords)))
		return
	}
	nodeRecords=tx.Records()["SubNode1"]
	if len(nodeRecords)!=20 {
		t.Fail()
		Error("Expected 20 but got:"+strconv.Itoa(len(nodeRecords)))
		return
	}
}
