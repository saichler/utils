package tests

import (
	"fmt"
	. "github.com/saichler/utils/golang/orm/updater"
	"testing"
)

func TestUpdater(t *testing.T) {
	nodes:=InitTestModel(2)
	old:=nodes[0]
	new:=nodes[1]
	fmt.Println(old.String)
	fmt.Println(old.SliceString)
	fmt.Println(old.SliceOfPtr[0].String)
	Update(old,new)
	fmt.Println(old.String)
	fmt.Println(old.SliceString)
	fmt.Println(old.SliceOfPtr[0].String)
}
