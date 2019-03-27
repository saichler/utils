package tests

import (
	. "github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/registry"
	"testing"
)

func TestOrmRegistryMainTable(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("Node")
	if nodeTable == nil {
		t.Fail()
		Error("Failed to register node")
		return
	}
}

func TestOrmRegistryTitleTag(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("Node")
	if nodeTable.Columns()["String"].MetaData().Title() != "Hello" {
		t.Fail()
		Error("Title tag does not work")
		return
	}
}

func TestOrmRegistrySizeTag(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("Node")
	if nodeTable.Columns()["String"].MetaData().Size() != 5 {
		t.Fail()
		Error("Size tag does not work")
		return
	}
}

func TestOrmRegistryIgnoreTag(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("Node")
	if !nodeTable.Columns()["String"].MetaData().Ignore(){
		t.Fail()
		Error("Ignore tag does not work")
		return
	}
}

func TestOrmRegistryMaskTag(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("Node")
	if !nodeTable.Columns()["String"].MetaData().Mask(){
		t.Fail()
		Error("Mask tag does not work")
		return
	}
}

func TestOrmRegistryColumns(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("Node")
	if len(nodeTable.Columns()) == 0 {
		t.Fail()
		Error("No columns in registry")
		return
	}
}

func TestOrmRegistrySubNode1(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable := registry.Table("SubNode1")
	if nodeTable == nil {
		t.Fail()
		Error("Did not find sub table 1")
		return
	}
}

func TestOrmRegistrySubNode4(t *testing.T) {
	registry := OrmRegistry{}
	registry.Register(Node{})
	nodeTable:=registry.Table("SubNode4")
	if nodeTable==nil {
		t.Fail()
		Error("Did not find sub table 4")
		return
	}
}
