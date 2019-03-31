package tests

import (
	"fmt"
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/transaction"
	"testing"
)

func TestUnMarshalString(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(5,tx)
	q:=NewQuery("Node")
	instances:=m.UnMarshal(q)
	fmt.Println(len(instances))
}
