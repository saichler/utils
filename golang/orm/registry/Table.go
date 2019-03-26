package registry

import "reflect"

type Table struct {
	ormRegistry *OrmRegistry
	structType reflect.Type
	columns map[string]*Column
}

func (t *Table) inspect() {
	if t.columns==nil {
		t.columns = make(map[string]*Column)
	}
	for i:=0;i<t.structType.NumField();i++ {
		field:=t.structType.Field(i)
		c:=t.columns[field.Name]
		if c==nil {
			c = &Column{}
			c.field = field
			c.table = t
			t.columns[field.Name]=c
			c.inspect()
		}
	}
}

func (t *Table) Columns() map[string]*Column {
	return t.columns
}
