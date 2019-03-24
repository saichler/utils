package tests

import "strconv"

type TestModel struct {
	Name string
	Val1  int
	Val2  int32
	Val3  bool
	Val4 int64
	Val5 uint
	Val6 uint32
	Node *TestNode
	Node2Nil *TestNode
	Nodes []*TestNode
	Nodes2Nil []*TestNode
}

type TestNode struct {
	NodeName string
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*TestNode
	MaoOfPtrNil map[string]*TestNode
	MapIntToStr map[int]string
}

func InitTestModel() *TestModel {
	tm:=&TestModel{}
	tm.Name = "Model Test"
	tm.Val1 = 153
	tm.Val2 = 32
	tn:=&TestNode{}
	tm.Node = tn
	tm.Nodes = buildNodes(10)
	tn.NodeName = "node name"
	return tm
}

func buildNode(loc int) *TestNode {
	node:=&TestNode{}
	node.NodeName = "Node #"+strconv.Itoa(loc)
	node.IntSlice = make([]int,2)
	node.IntSlice[1] = loc
	node.StringSlice = make([]string,2)
	node.StringSlice[1] = "str="+strconv.Itoa(loc)
	node.MapOfPtr = make(map[string]*TestNode,3)
	node.MapOfPtr["B"]=&TestNode{}
	node.MapOfPtr["B"].NodeName = "Map="+strconv.Itoa(loc)
	node.MapIntToStr = make(map[int]string)
	node.MapIntToStr[44] = "44"
	return node
}

func buildNodes(size int) []*TestNode {
	nodes:=make([]*TestNode,size)
	for i:=0;i<size/2;i++ {
		nodes[i] = buildNode(i)
	}
	return nodes
}