package tests

import "strconv"

type Node struct {
	String string
	Int  int
	Int32  int32
	Bool  bool
	Int64 int64
	Uint uint
	Uint32 uint32
	Uint64 uint64
	Float32 float32
	Float64 float64
	Ptr *Node
	NilPtr *Node
	SliceOfPtr []*Node
	NilSliceOfPtr []*Node
	SliceInt []int
	SliceString []string
	MapStringPtr map[string]*Node
	MapStringPtrNil map[string]*Node
	MapIntString map[int]string
	SubNode1Slice []*SubNode1
	SubNode2Slice []*SubNode2
}

type SubNode1 struct {
	String string
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*Node
	MaoOfPtrNil map[string]*Node
	MapIntToStr map[int]string
}

type SubNode2 struct {
	String string
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*Node
	MaoOfPtrNil map[string]*Node
	MapIntToStr map[int]string
	SliceInSlice []*SubNode3
}

type SubNode3 struct {
	String string
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*Node
	MaoOfPtrNil map[string]*Node
	MapIntToStr map[int]string
}

func InitTestModel() *Node {
	n:=&Node{}
	n.String = "String"
	n.Int = -101
	n.Int32 = -102
	n.Bool = true
	n.Int64 = -103
	n.Uint = 104
	n.Uint32 = 105
	n.Uint64 = 106
	n.Float32 = -107.23
	n.Float64 = -108.25

	Ptr *Node
	NilPtr *Node
	SliceOfPtr []*Node
	NilSliceOfPtr []*Node
	SliceInt []int
	SliceString []string
	MapStringPtr map[string]*Node
	MapStringPtrNil map[string]*Node
	MapIntString map[int]string
	SubNode1Slice []*SubNode1
	SubNode2Slice []*SubNode2



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