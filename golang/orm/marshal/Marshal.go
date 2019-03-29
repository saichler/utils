package marshal

import (
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"reflect"
)

var marshalers = make(map[reflect.Kind]func(reflect.Value,*OrmRegistry,*Transaction,Persistency)(reflect.Value,error))
func initMarshalers() {
	if len(marshalers)==0 {
		marshalers[reflect.Ptr]=ptrMarshal
		marshalers[reflect.Struct]=structMarshal
		marshalers[reflect.Map]=mapMarshal
		marshalers[reflect.Slice]=sliceMarshal
		marshalers[reflect.String]=defaultMarshal
		marshalers[reflect.Int]=defaultMarshal
		marshalers[reflect.Int32]=defaultMarshal
		marshalers[reflect.Int64]=defaultMarshal
		marshalers[reflect.Uint]=defaultMarshal
		marshalers[reflect.Uint32]=defaultMarshal
		marshalers[reflect.Uint64]=defaultMarshal
		marshalers[reflect.Float64]=defaultMarshal
		marshalers[reflect.Float32]=defaultMarshal
		marshalers[reflect.Bool]=defaultMarshal
	}
}

func Marshal(any interface{},r *OrmRegistry,tx *Transaction, pr Persistency) error {
	initMarshalers()
	if any==nil {
		return nil
	}
	value:=reflect.ValueOf(any)
	value,err:=marshal(value,r,tx,pr)
	return err
}

func marshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency) (reflect.Value,error) {
	marshaler:=marshalers[value.Kind()]
	if marshaler==nil {
		panic("No Marshaler for kind "+value.Kind().String())
	}
	return marshaler(value,r,tx,pr)
}

func ptrMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency) (reflect.Value,error) {
	if value.IsNil() {
		return value,nil
	}
	v:=value.Elem()
	return marshal(v,r,tx,pr)
}

func structMarshal(value reflect.Value,r *OrmRegistry,tx *Transaction,pr Persistency) (reflect.Value,error) {
	tableName:=value.Type().Name()
	//No need to do anything, nameless struct
	if tableName=="" {
		return value,nil
	}

	table:=r.Table(tableName)
	if table==nil {
		panic("Table:"+tableName+" was not registered!")
	}

	rec:=tx.AddRecord(tableName)
	for fieldName,column:=range table.Columns() {
		fieldValue:=value.FieldByName(fieldName)
		marshalValue,err:=marshal(fieldValue,r,tx,pr)
		if err!=nil {
			panic(err)
		}
		if column.MetaData().ColumnTableName()=="" {
			rec.Set(fieldName,marshalValue)
		} else {
			Process sub table marshaling here
		}
	}
	return reflect.ValueOf(rec),nil
}

func sliceMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency) (reflect.Value,error) {
	if value.IsNil() {
		return value,nil
	}
	list:=make([]reflect.Value,value.Len())
	for i:=0;i<value.Len();i++ {
		v,e:=marshal(value.Index(i),r,tx,pr)
		if e!=nil {
			return v,e
		}
		list = append(list,v)
	}
	return reflect.ValueOf(list),nil
}

func mapMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency) (reflect.Value,error) {
	return value,nil
}

func defaultMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency) (reflect.Value,error) {
	return value,nil
}