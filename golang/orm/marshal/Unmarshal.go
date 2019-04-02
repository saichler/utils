package marshal

import (
	"fmt"
	"github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/registry"
	. "github.com/saichler/utils/golang/orm/transaction"
	"reflect"
	"strconv"
)

var setters = make(map[reflect.Kind]func(reflect.Value,*Column,*Record,*RecordID,*Transaction))

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
	instances:=unmarshal(ormQuery,m.tx,m.ormRegistry,NewRecordID())
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

	records:=tx.AllRecords(query.TableName())
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
				key:=record.PrimaryIndex(table.Indexes().PrimaryIndex())
				id.Add(table.Name(),column.Name(),key)
				set(field, column, record,id,tx)
				id.Del()
			}
		}
	}
	return result
}

func set(field reflect.Value,column *Column,record *Record,id *RecordID,tx *Transaction) {
	setter:=setters[field.Kind()]
	if setter==nil {
		panic("No Setter for kind:"+field.Kind().String())
	}
	setter(field,column,record,id,tx)
}

func setPtr(field reflect.Value,column *Column,record *Record,id *RecordID,tx *Transaction) {
	ptrKind:=field.Type().Elem().Kind()
	if ptrKind==reflect.Struct {
		table:=column.Table().OrmRegistry().Table(column.MetaData().ColumnTableName())
		if table==nil {
			panic("Cannot find table with name: "+column.MetaData().ColumnTableName())
		}
		if table.Indexes().PrimaryIndex()!=nil {
			key:=record.Get(column.Name()).String()
			if key=="" {
				return
			}
			rec:=tx.Records(table.Name(),key)
			please continue here
			panic("column="+column.Name()+"key="+key+":"+strconv.Itoa(len(rec)))
		}
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
	set(field.Elem(),column,record,id,tx)
}

func setDefault(field reflect.Value,column *Column,record *Record,id *RecordID,tx *Transaction) {
	field.Set(record.Get(column.Name()))
}

func setStruct(field reflect.Value,column *Column,record *Record,id *RecordID,tx *Transaction) {
	table:=column.Table().OrmRegistry().Table(field.Type().Name())
	if table==nil {
		panic("Cannot find table name:"+field.Type().Name())
	}
	if table.Name()=="Node" {
		panic(column.Name())
	}
	fmt.Println("Struct:"+table.Name()+":"+column.Name())
	instance := table.NewInstance()
	for _, c := range table.Columns() {
		var fld reflect.Value
		if instance.Kind() == reflect.Ptr {
			fld = instance.Elem().FieldByName(c.Name())
		} else {
			fld = instance.FieldByName(c.Name())
		}
		var key="0"
		if table.Indexes().PrimaryIndex()!=nil {
			key = record.PrimaryIndex(table.Indexes().PrimaryIndex())
		}
		id.Add(table.Name(),c.Name(),key)
		set(fld, c, record,id,tx)
		id.Del()
	}
}

func setMap(field reflect.Value,column *Column,record *Record,id *RecordID,tx *Transaction) {
}

func setSlice(field reflect.Value,column *Column,record *Record,id *RecordID,tx *Transaction) {
	value:=record.Get(column.Name())
	vString:=value.String()
	if value.IsValid() {
		if vString==""{
			return
		}
		if column.MetaData().ColumnTableName()=="" {
			v:=utils.FromString(vString,column.Type())
			field.Set(v)
		} else {
			//Keyed ptr slice
			/*
			table:=column.Table().OrmRegistry().Table(column.MetaData().ColumnTableName())
			if table==nil {
				panic("No Table was found with name:"+column.MetaData().ColumnTableName())
			}
			//recs:=tx.Records(table.Name(),id.String())
			fmt.Println(column.Name()+":"+vString)
			fmt.Println(id.String())
			*/
		}
	} else if column.MetaData().ColumnTableName()!="" {

		table:=column.Table().OrmRegistry().Table(column.MetaData().ColumnTableName())
		if table==nil {
			panic("No Table was found with name:"+column.MetaData().ColumnTableName())
		}
		fmt.Println("Slice "+table.Name()+":"+column.Name())
		recs:=tx.Records(table.Name(),id.String())
		newSlice:=reflect.MakeSlice(column.Type(),len(recs),len(recs))
		for _,rec:=range recs {
			instance:=table.NewInstance()
			set(instance.Elem(),column,rec,id,tx)
		}
		field.Set(newSlice)
	}
}