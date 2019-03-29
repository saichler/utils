package tests

import (
	"fmt"
	"github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/marshal"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"testing"
)

func TestMarshal(t *testing.T) {
	registry := &OrmRegistry{}
	registry.Register(Node{})
	node:=InitTestModel(1)[0]
	tx:=&Transaction{}
	Marshal(node,registry,tx,nil)
	records:=tx.Records()

	nodeRecords:=records["Node"]
	fmt.Println(len(nodeRecords))
	for _,rec:=range nodeRecords {
		if rec.Map()["String"].String()=="String-0" {
			for k, v := range rec.Map() {
				utils.Info(k + "=" + utils.ToString(v))
			}
		}
	}
}
