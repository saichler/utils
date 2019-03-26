package registry

import (
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang"
	"reflect"
	"strconv"
	"strings"
)

type Column struct {
	table *Table
	field reflect.StructField

	title string
	size int
	ignore bool
	mask bool
	primaryKey string
	uniqueKey[] string
	nonUniqueKey[] string
}

func (c *Column) inspect() {
	c.parseTags()
	if c.IsStruct() {
		strct:=getStruct(c.field.Type)
		c.table.ormRegistry.register(strct)
	}
}

func (c *Column) IsStruct() bool {
	return isStruct(c.field.Type)
}

func isStruct(typ reflect.Type) bool {
	if typ.Kind()==reflect.Struct {
		return true
	} else if typ.Kind()==reflect.Ptr {
		return isStruct(typ.Elem())
	} else if typ.Kind()==reflect.Slice {
		return isStruct(typ.Elem())
	} else if typ.Kind()==reflect.Map {
		return isStruct(typ.Elem())
	}
	return false
}

func getStruct(typ reflect.Type) reflect.Type {
	if typ.Kind()==reflect.Struct {
		return typ
	} else if typ.Kind()==reflect.Ptr {
		return getStruct(typ.Elem())
	} else if typ.Kind()==reflect.Slice {
		return getStruct(typ.Elem())
	} else if typ.Kind()==reflect.Map {
		return getStruct(typ.Elem())
	}
	return nil
}

func (c *Column) parseTags() {
	tags:=string(c.field.Tag)
	if tags=="" {
		return
	}
	splits:=strings.Split(tags, " ")
	for _,tag:=range splits {
		c.parseTag(tag)
	}
}

func (c *Column) parseTag(tag string) {
	c.title = c.field.Name
	c.size = 128
	if strings.Trim(tag, " ") != "" {
		return
	}
	index := strings.Index(tag, ":")
	if index == -1 {
		return
	}
	name := tag[0:index]
	value := tag[index+1 : len(tag)-1]
	if name == TITLE {
		c.title = value
	} else if name == SIZE {
		val,err:=strconv.Atoi(value)
		if err!=nil {
			Error("Unable to parse field size from:"+value+" in field:"+c.field.Name)
		} else {
			c.size = val
		}
	} else if name == MASK {
		c.mask = true
	} else if name == IGNORE {
		c.ignore = true
	} else if name == PRIMARY_KEY {
		c.primaryKey = value
	} else if name == UNIQUE_KEY {
		if c.uniqueKey==nil {
			c.uniqueKey = make([]string,0)
		}
		c.uniqueKey = append(c.uniqueKey,value)
	} else if name == NON_UNIQUE_KEY {
		if c.nonUniqueKey==nil {
			c.nonUniqueKey = make([]string,0)
		}
		c.nonUniqueKey = append(c.nonUniqueKey,value)
	}
}
