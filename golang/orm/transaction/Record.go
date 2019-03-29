package transaction

import (
	"bytes"
	. "github.com/saichler/utils/golang/orm/registry"
	"reflect"
	"strconv"
)

type Record struct {
	recordData map[string]reflect.Value
}

func (rec *Record) Set(key string,value reflect.Value) {
	if rec.recordData==nil {
		rec.recordData = make(map[string]reflect.Value)
	}
	rec.recordData[key]=value
}

func (rec *Record) PrimaryIndex(pi *Index) string {
	buff:=bytes.Buffer{}
	buff.WriteString("[")
	buff.WriteString(pi.Table().Name())
	for _,column:=range pi.Columns() {
		buff.WriteString(".")
		val:=rec.recordData[column.Name()]
		sv:=StringValue(val)
		buff.WriteString(column.Name())
		buff.WriteString("=")
		buff.WriteString(sv)
	}
	buff.WriteString("]")
	return buff.String()
}

func StringValue(v reflect.Value) string {
	if v.Kind()==reflect.String {
		return v.String()
	} else if v.Kind()==reflect.Int || v.Kind()==reflect.Int16 || v.Kind()==reflect.Int32 || v.Kind()==reflect.Int64 {
		return strconv.Itoa(int(v.Int()))
	} else {
		panic("Please implemenet String Value for kind:"+v.Kind().String())
	}
}