package marshal

import (
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"reflect"
)

func (m *Marshaler) UnMarshal(ormQuery *Query) []interface{} {
	instances:=unmarshal(ormQuery.TableName(),m.tx,m.ormRegistry)
	result:=make([]interface{},len(instances))
	for i:=0;i<len(result);i++ {
		result[i] = instances[i].Interface()
	}
	return result
}

func unmarshal(tablename string, tx *Transaction,ormRegistry *OrmRegistry) []reflect.Value {
	table:=ormRegistry.Table(tablename)
	if table==nil {
		panic("Unknown table "+tablename)
	}
	result:=make([]reflect.Value,0)

	records:=tx.Records()[tablename]

	for _,record:=range records {
		instance:=table.NewInstance()
		result = append(result,instance)
		for _,column:=range table.Columns() {
			set(instance,column,record)
		}
	}
	return result
}

func set(instance reflect.Value,column *Column,record *Record) {
	if instance.Kind()==reflect.Ptr {
		set(instance.Elem(),column,record)
		return
	}
	fld:=instance.FieldByName(column.Name())
	if fld.Kind()==reflect.String {
		fld.SetString(record.Map()[column.Name()].String())
	}
}