package marshal

import (
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"reflect"
)

var marshalers = make(map[reflect.Kind]func(reflect.Value,*OrmRegistry,*Transaction,Persistency,*KeyPath)(reflect.Value,error))
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
	value,err:=marshal(value,r,tx,pr,newKeyPath())
	return err
}

func marshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency,kp *KeyPath) (reflect.Value,error) {
	marshaler:=marshalers[value.Kind()]
	if marshaler==nil {
		panic("No Marshaler for kind "+value.Kind().String())
	}
	return marshaler(value,r,tx,pr,kp)
}

func ptrMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency,kp *KeyPath) (reflect.Value,error) {
	if value.IsNil() {
		return value,nil
	}
	v:=value.Elem()
	return marshal(v,r,tx,pr,kp)
}

func structMarshal(value reflect.Value,r *OrmRegistry,tx *Transaction,pr Persistency,kp *KeyPath) (reflect.Value,error) {
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
	subTables:=make([]*Column,0)
	for fieldName,column:=range table.Columns() {
		if column.MetaData().ColumnTableName()=="" {
			fieldValue:=value.FieldByName(fieldName)
			marshalValue,err:=marshal(fieldValue,r,tx,pr,kp)
			if err!=nil {
				panic(err)
			}
			rec.Set(fieldName,marshalValue)
		} else {
			subTables = append(subTables,column)
		}
	}

	if table.Indexes().PrimaryIndex()!=nil {
		kp.add(rec.PrimaryIndex(table.Indexes().PrimaryIndex()))

		for _,sbColumn:=range subTables {
			fieldValue:=value.FieldByName(sbColumn.Name())
			sbValue,err:=marshal(fieldValue,r,tx,pr,kp)
			if err!=nil {
				return reflect.ValueOf(rec),err
			}
			sbTable:=r.Table(sbColumn.MetaData().ColumnTableName())
			if sbTable.Indexes().PrimaryIndex()!=nil{
				rec.Set(sbColumn.Name(),reflect.ValueOf(StringValue(sbValue)))
			}
		}
		kp.del()
	}

	return reflect.ValueOf(rec),nil
}

func sliceMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency,kp *KeyPath) (reflect.Value,error) {
	if value.IsNil() {
		return value,nil
	}
	list:=make([]reflect.Value,value.Len())
	for i:=0;i<value.Len();i++ {
		v,e:=marshal(value.Index(i),r,tx,pr,kp)
		if e!=nil {
			return v,e
		}
		list = append(list,v)
	}
	return reflect.ValueOf(list),nil
}

func mapMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency,kp *KeyPath) (reflect.Value,error) {
	return value,nil
}

func defaultMarshal(value reflect.Value,r *OrmRegistry, tx *Transaction,pr Persistency,kp *KeyPath) (reflect.Value,error) {
	return value,nil
}