package tests

import (
	. "github.com/saichler/utils/golang/orm/marshal"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"testing"
)

func TestMarshal(t *testing.T) {
	registry := &OrmRegistry{}
	registry.Register(Node{})
	node:=InitTestModel(1)[0]
	Marshal(node,registry,&Transaction{},nil)
}
