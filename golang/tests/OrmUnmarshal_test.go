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