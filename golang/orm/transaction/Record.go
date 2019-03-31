package transaction

import (
	"github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/registry"
	"reflect"
)

type Record struct {
	recordData map[string]reflect.Value
}

func (rec *Record) init(){
	if rec.recordData==nil {
		rec.recordData = make(map[string]reflect.Value)
	}
}

func (rec *Record) Set(key string,value reflect.Value) {
	rec.init()
	rec.recordData[key]=value
}

func (rec *Record) PrimaryIndex(pi *Index) string {
	result:=utils.NewStringBuilder("")
	for _,column:=range pi.Columns() {
		val:=rec.recordData[column.Name()]
		sv:=utils.ToString(val)
		result.Append(sv)
	}
	return result.String()
}

func (rec *Record) Map() map[string]reflect.Value {
	rec.init()
	return rec.recordData
}
