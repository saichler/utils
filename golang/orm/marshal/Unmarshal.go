package marshal

import (
	"github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"reflect"
)

var setters = make(map[reflect.Kind]func(reflect.Value,*Column,*Record,*RecordID))

func initSetters(){
	if len(setters)==0 {
		setters[reflect.Ptr]=setPtr
		setters[reflect.String]=setDefault
		setters[reflect.Float32]=setDefault
		setters[reflect.Float64]=setDefault
		setters[reflect.Uint]=setDefault
		setters[reflect.Uint16]=setDefault
		setters[reflect.Uint32]=setDefault
		setters[reflect.Uint64]=setDefault
		setters[reflect.Int]=setDefault
		setters[reflect.Int16]=setDefault
		setters[reflect.Int32]=setDefault
		setters[reflect.Int64]=setDefault
		setters[reflect.Bool]=setDefault
		setters[reflect.Struct]=setStruct
		setters[reflect.Map]=setMap
		setters[reflect.Slice]=setSlice
	}
}

func (m *Marshaler) UnMarshal(ormQuery *Query) []interface{} {
	initSetters()
	instances:=unmarshal(ormQuery,m.tx,m.ormRegistry,&RecordID{})
	result:=make([]interface{},len(instances))
	for i:=0;i<len(result);i++ {
		result[i] = instances[i].Interface()
	}
	return result
}

func unmarshal(query *Query, tx *Transaction,ormRegistry *OrmRegistry,id *RecordID) []reflect.Value {
	table:=ormRegistry.Table(query.TableName())
	if table==nil {
		panic("Unknown table "+query.TableName())
	}
	result:=make([]reflect.Value,0)

	records:=tx.Records()[query.TableName()]

	for _,record:=range records {
		if record.Get(RECORD_LEVEL).Int()==0 || !query.OnlyTopLevel() {
			instance := table.NewInstance()
			result = append(result, instance)
			for _, column := range table.Columns() {
				var field reflect.Value
				if instance.Kind() == reflect.Ptr {
					field = instance.Elem().FieldByName(column.Name())
				} else {
					field = instance.FieldByName(column.Name())
				}
				id.Add(table.Name(),column.Name(),record.PrimaryIndex(table.Indexes().PrimaryIndex()))
				set(field, column, record,id)
				id.Del()
			}
		}
	}
	return result
}

func set(field reflect.Value,column *Column,record *Record,id *RecordID) {
	setter:=setters[field.Kind()]
	if setter==nil {
		panic("No Setter for kind:"+field.Kind().String())
	}
	setter(field,column,record,id)
}

func setPtr(field reflect.Value,column *Column,record *Record,id *RecordID) {
	ptrKind:=field.Type().Elem().Kind()
	if ptrKind==reflect.Struct {
		newPtr := reflect.New(field.Type().Elem())
		field.Set(newPtr)
	} else if ptrKind==reflect.Slice {
		newSlice := reflect.MakeSlice(reflect.SliceOf(field.Type().Elem()), 0, 0)
		field.Elem().Set(newSlice.Elem())
	} else if ptrKind==reflect.Map {
		newMap:=reflect.MakeMapWithSize(field.Type(), 0)
		field.Elem().Set(newMap.Elem())
	} else {
		panic("No Ptr Handle of:"+ptrKind.String())
	}
	set(field.Elem(),column,record,id)
}

func setDefault(field reflect.Value,column *Column,record *Record,id *RecordID) {
	field.Set(record.Get(column.Name()))
}

func setStruct(field reflect.Value,column *Column,record *Record,id *RecordID) {
}

func setMap(field reflect.Value,column *Column,record *Record,id *RecordID) {
}

func setSlice(field reflect.Value,column *Column,record *Record,id *RecordID) {
	value:=record.Get(column.Name())
	vString:=value.String()
	//This slice is NOT a slice of pointers to a non key table
	if value.IsValid() && vString!="" {
		if column.MetaData().ColumnTableName()=="" {
			v:=utils.FromString(vString,column.Type())
			field.Set(v)
		} else {

		}
	}
}