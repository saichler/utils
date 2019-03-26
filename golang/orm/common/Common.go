package common

import "strings"

const (
	TITLE = "Title"
	SIZE = "Size"
	MASK = "MASK"
	IGNORE = "Ignore"
	PRIMARY_KEY = "PrimaryKey"
	UNIQUE_KEY = "UniqueKey"
	NON_UNIQUE_KEY = "NonUniqueKey"
)

func GetTag(tag,tags string) string {
	index:= strings.Index(tags,tag+":")
	if index==-1 {
		return ""
	}
	subset:=tags[index+len(tag)+2:]
	index = strings.Index(subset,"\"")
	return subset[0:index]
}
