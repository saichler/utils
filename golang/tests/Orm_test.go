package tests

import (
	. "github.com/saichler/utils/golang"
	"testing"
	. "github.com/saichler/utils/golang/orm/registry"
	)

func TestOrmRegistry(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable:=registry.Table("Node")
	if nodeTable==nil {
		t.Fail()
		Error("Failed to register node")
		return
	}
	if len(nodeTable.Columns())==0 {
		t.Fail()
		Error("No columns in registry")
		return
	}
	nodeTable=registry.Table("SubNode1")
	if nodeTable==nil {
		t.Fail()
		Error("Did not find sub table 1")
		return
	}
	nodeTable=registry.Table("SubNode4")
	if nodeTable==nil {
		t.Fail()
		Error("Did not find sub table 4")
		return
	}
}
