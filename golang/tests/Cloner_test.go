package tests

import (
	"fmt"
	"strconv"
	"testing"
)
import . "github.com/saichler/utils/golang"

func TestCloner(t *testing.T) {
	tm:=InitTestModel()
	tmClone:=Clone(tm).(*TestModel)
	if tmClone.Name!="Model Test" {
		t.Fail()
		fmt.Println("Fail @1")
		return
	}
	if tmClone.Node.NodeName!="node name" {
		t.Fail()
		fmt.Println("Fail @2")
		return
	}
	if tmClone.Val1!=153 {
		t.Fail()
		fmt.Println("Fail @3")
		return
	}
	if tmClone.Nodes[1].NodeName!="Node #1" {
		t.Fail()
		fmt.Println("Fail @4")
		return
	}
	if tmClone.Nodes[1].IntSlice[1]!=1 {
		t.Fail()
		fmt.Println("Fail @5")
		return
	}
	if tmClone.Nodes[1].StringSlice[1]!="str=1" {
		t.Fail()
		fmt.Println("Fail @6")
		return
	}
	if tmClone.Nodes[1].MapOfPtr["B"].NodeName!="Map=1" {
		t.Fail()
		fmt.Println("Fail @7")
		return
	}
	if tmClone.Nodes[1].MapIntToStr[44]!="44" {
		t.Fail()
		fmt.Println("Fail @8")
		return
	}
	fmt.Println(strconv.Itoa(tmClone.Val1))
	fmt.Println(tmClone.Name)
	fmt.Println(tmClone.Node.NodeName)
	fmt.Println(tmClone.Nodes[1].NodeName)
	fmt.Println(tmClone.Nodes[1].IntSlice[1])
	fmt.Println(tmClone.Nodes[1].StringSlice[1])
	fmt.Println(tmClone.Nodes[1].MapOfPtr["B"].NodeName)
	fmt.Println(tmClone.Nodes[1].MapIntToStr[44])
}
