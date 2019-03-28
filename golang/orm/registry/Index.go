package registry

import (
	. "github.com/saichler/utils/golang"
	"strconv"
	"strings"
)

type Indexes struct {
	primaryIndex     *Index
	uniqueIndexes    map[string]*Index
	nonUniqueIndexes map[string]*Index
}

type Index struct {
	name    string
	columns []*Column
	unique  bool
}

func (indexes *Indexes) AddColumn(column *Column) {
	indexes.updateIndex(column.metaData.primaryKey, column, true, true)
	indexes.updateIndex(column.metaData.uniqueKeys, column, false, true)
	indexes.updateIndex(column.metaData.nonUniqueKeys, column, false, false)
}

func newIndex(name string, unique bool) *Index {
	index := &Index{}
	index.unique = unique
	index.name = name
	index.columns = make([]*Column, 0)
	return index
}

func (indexes *Indexes) updateIndex(data string, column *Column, primary, unique bool) {
	if data != "" {
		im := getIndexMap(data)
		for indexName, columnPos := range im {
			var index *Index
			if primary {
				if indexes.primaryIndex == nil {
					indexes.primaryIndex = newIndex(indexName, true)
				}
				index = indexes.primaryIndex
			} else if unique {
				if indexes.uniqueIndexes == nil {
					indexes.uniqueIndexes = make(map[string]*Index)
				}
				index = indexes.uniqueIndexes[indexName]
				if index == nil {
					index = newIndex(indexName, true)
					indexes.uniqueIndexes[indexName] = index
				}
			} else {
				if indexes.nonUniqueIndexes == nil {
					indexes.nonUniqueIndexes = make(map[string]*Index)
				}
				index = indexes.nonUniqueIndexes[indexName]
				if index == nil {
					index = newIndex(indexName, false)
					indexes.nonUniqueIndexes[indexName] = index
				}
			}
			if len(index.columns) <= columnPos {
				for i := len(index.columns); i <= columnPos; i++ {
					index.columns = append(index.columns, nil)
				}
			}
			index.columns[columnPos] = column
		}
	}
}

func getIndexMap(indexStr string) map[string]int {
	result := make(map[string]int)
	splits := strings.Split(indexStr, ",")
	for _, indexDef := range splits {
		i := strings.Index(indexDef, ":")
		if i != -1 {
			indexName := indexDef[0:i]
			loc := indexDef[i+1:]
			indexLoc, err := strconv.Atoi(loc)
			if err != nil {
				Error(err)
			} else {
				result[indexName] = indexLoc
			}
		}
	}
	return result
}

func (indxs *Indexes) PrimaryIndex() *Index {
	return indxs.primaryIndex
}

func (indxs *Indexes) UniqueIndexes() map[string]*Index {
	return indxs.uniqueIndexes
}

func (indxs *Indexes) NonUniqueIndexes() map[string]*Index {
	return indxs.nonUniqueIndexes
}

func (index *Index) Columns() []*Column {
	return index.columns
}
