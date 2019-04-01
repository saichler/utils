package tests

import (
	. "github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/transaction"
	"strconv"
	"testing"
)

func TestUnMarshalCountOnlyTopLevel(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
}

func TestUnMarshalCountAllLevels(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",false)
	instances:=m.UnMarshal(q)
	if len(instances)!=30 {
		t.Fail()
		Error("Expected:30 but got "+strconv.Itoa(len(instances)))
	}
}

func TestUnMarshalString(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		expected:="String-"+strconv.Itoa(i)
		found:=false
		for _,n:=range instances {
			node:=n.(*Node)
			if node.String==expected {
				found=true
				break
			}
		}
		if !found {
			t.Fail()
			Error("Failed to find String "+expected)
		}
	}
}

func TestUnMarshalInt(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		expected:=-101+i
		found:=false
		for _,n:=range instances {
			node:=n.(*Node)
			if node.Int==expected {
				found=true
				break
			}
		}
		if !found {
			t.Fail()
			Error("Failed to find Int "+strconv.Itoa(expected))
		}
	}
}