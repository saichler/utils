package dsutils

import (
	"bytes"
	"strconv"
	"strings"
)

const (
	TAB = "    ";
)

var nextUntaggedIndex = 0

func NewYamlRoot() *YamlNode {
	root:=&YamlNode{}
	root.Init("root","",-1)
	return root
}

func getNextUntagged() string {
	nextUntaggedIndex++
	return strconv.Itoa(nextUntaggedIndex)
}

func removeQuotes(str string) string {
	result := Trim(str)
	if result[0:1]=="\"" || result[0:1]=="'" {
		return result[1:len(result)-1]
	}
	return result
}

func Trim(str string) string {
	begin:=true
	result:=""
	subset:=""
	for i:=0;i<len(str);i++ {
		if begin && (str[i:i+1]==" " ||
			str[i:i+1]=="\n" ||
			str[i:i+1]=="\r" ||
			str[i:i+1]=="\t") {
				continue
		} else if str[i:i+1]==" " ||
			str[i:i+1]=="\n" ||
			str[i:i+1]=="\r" ||
			str[i:i+1]=="\t" {
			subset+=str[i:i+1]
			continue
		}
		result+=subset+str[i:i+1]
		begin = false
		subset=""
	}
	return result
}

func parseTag(line string) string {
	line = Trim(line)
	if "-" == line {
		return "List-" + getNextUntagged()
	}
	index:=strings.Index(line,":")
	if index == -1 {
		return "Tag-" + getNextUntagged();
	}
	return removeQuotes(line[0:index]);
}

func parsetValue(line string) string {
	line = Trim(line)
	index := strings.Index(line, ":")
	index2 := strings.Index(line, "#")

	if index == len(line)-1 {
		return ""
	}

	if index == -1 && index2 == -1 {
		return Trim(line)
	}

	if index == -1 && index2 != -1 {
		return removeQuotes(line[0:index2])
	}

	if index2 == -1 {
		return removeQuotes(line[index+1:])
	}
	return removeQuotes(line[index+1:index2])
}

func parseLevel(line string) int {
	index := 0
	for ;line[index:index+1]==" "; {
		index++
	}
	if index == 0 {
		return 0
	}
	result := index / len(TAB)
	if index % len(TAB)!=0 {
		result++;
	}
	return result;
}

func getTab(lvl int) string {
	result:=bytes.Buffer{}
	for i := 0; i < lvl; i++ {
		result.WriteString(TAB)
	}
	return result.String()
}

func Parse(data,tag string, root *YamlNode) *YamlNode {
	lr:=NewLineReader(data)
	line:=lr.NextLine()
	for; Trim(line)=="" || Trim(line)[0:1]=="#"; {
		line = lr.NextLine()
	}
	tagRoot:=&YamlNode{}
	tagRoot.Init(tag,"",-1)
	root.AddChild(tagRoot)
	parent:=tagRoot
	for ;line!=EOF; {
		node:= &YamlNode{}
		node.Parse(line)
		if parent.lvl==node.lvl {
			parent = parent.parent
		} else if parent.lvl > node.lvl {
			for ; parent.lvl > node.lvl; {
				parent = parent.parent
			}
			parent = parent.parent;
		}
		node.parent = parent
		parent.AddChild(node)
		parent = node
		line = lr.NextLine()
		for; Trim(line)=="" || Trim(line)[0:1]=="#"; {
			line = lr.NextLine()
		}

	}
	return tagRoot
}


