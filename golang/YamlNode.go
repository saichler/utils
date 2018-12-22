package dsutils

import "bytes"

type YamlNode struct {
	tag string
	value string
	lvl int
	childrenMap map[string]*YamlNode
	childrenList []*YamlNode
	parent *YamlNode
}

func (yn *YamlNode) Parse(line string) {
	yn.Init(parseTag(line),parsetValue(line),parseLevel(line))
}

func (yn *YamlNode) Init(tag, value string, lvl int) {
	yn.tag = tag
	yn.value = value
	yn.lvl = lvl
}

func (yn *YamlNode) GetTag() string {
	return yn.tag
}

func (yn *YamlNode) GetValue() string {
	return yn.value
}

func (yn *YamlNode) GetLvl() int {
	return yn.lvl
}

func (yn *YamlNode) GetParent() *YamlNode {
	return yn.parent
}

func (yn *YamlNode) GetChildByTag(tag string) *YamlNode {
	if yn.childrenMap==nil {
		return nil
	}
	return yn.childrenMap[tag]
}

func (yn *YamlNode) GetChildrenList() []*YamlNode {
	return yn.childrenList
}

func (yn *YamlNode) GetKey() string {
	buff:=&bytes.Buffer{}
	yn.getKey(buff)
	return buff.String()
}

func (yn *YamlNode) getKey(buff *bytes.Buffer) {
	if yn.parent == nil {
		buff.WriteString("/")
		buff.WriteString(yn.tag)
		return
	}
	yn.parent.getKey(buff)
	buff.WriteString("/")
	buff.WriteString(yn.tag)
}

func (yn *YamlNode) AddChild(node *YamlNode) {
	if yn.childrenList==nil {
		yn.childrenList = make([]*YamlNode,0)
		yn.childrenMap = make(map[string]*YamlNode)
	}
	yn.childrenMap[node.tag]=node
	yn.childrenList = append(yn.childrenList,node)
	node.parent = yn
}

func (yn *YamlNode) String() string {
	buff:=&bytes.Buffer{}
	yn.string(buff)
	return buff.String()
}

func (yn *YamlNode) string(buff *bytes.Buffer) {
	if yn.lvl>=0 {
		buff.WriteString(getTab(yn.lvl))
		buff.WriteString(yn.tag)
		buff.WriteString(": ")
		buff.WriteString(yn.value)
		buff.WriteString("\n")
	}
	if yn.childrenList!=nil {
		for _,c:=range yn.childrenList {
			c.string(buff)
		}
	}
}
