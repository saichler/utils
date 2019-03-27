package registry

import (
	. "github.com/saichler/utils/golang"
	"strconv"
	"strings"
)

type Indexes struct {
	indexes map[string]*Index
	primaryIndex string
}

type Index struct {
	name string
	columns []*Column
}

func (indexes *Indexes) AddColumn(column *Column) {
	primaryName:=indexes.updateIndex(column.metaData.primaryKey,column)
	if primaryName!="" {
		indexes.primaryIndex = primaryName
	}
	indexes.updateIndex(column.metaData.uniqueKeys,column)
	indexes.updateIndex(column.metaData.nonUniqueKeys,column)
}

func (indexes *Indexes) updateIndex(data string,column *Column) string {
	indexName:=""
	if data!="" {
		im:=getIndexMap(data)
		for k,v:=range im {
			indexName = k
			indexes.primaryIndex = k
			index := indexes.indexes[k]
			if index==nil {
				index = &Index{}
				index.name = k
				index.columns = make([]*Column,0)
				indexes.indexes[k]=index
			}
			index.columns[v]=column
		}
	}
	return indexName
}

func getIndexMap(indexStr string) map[string]int {
	result:=make(map[string]int)
	splits:=strings.Split(indexStr,",")
	for _,indexDef:=range splits {
		i:=strings.Index(indexDef,":")
		if i!=-1 {
			indexName:=indexDef[0:i]
			loc:=indexDef[i+1:]
			indexLoc,err:=strconv.Atoi(loc)
			if err!=nil {
				Error(err)
			} else {
				result[indexName]=indexLoc
			}
		}
	}
	return result
}


