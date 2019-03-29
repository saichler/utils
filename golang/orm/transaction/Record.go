package transaction

import "reflect"

type Record struct {
	recordData map[string]reflect.Value
}

func (rec *Record) Set(key string,value reflect.Value) {
	if rec.recordData==nil {
		rec.recordData = make(map[string]reflect.Value)
	}
	rec.recordData[key]=value
}