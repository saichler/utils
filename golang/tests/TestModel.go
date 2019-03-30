package tests

import "strconv"

type Node struct {
	String string `Title=Hello Size=5 PrimaryKey=name:0`
	String2 string `Ignore=true`
	String3 string `Mask=true`
	String4 string `UniqueKey=uk1:1`
	String5 string `UniqueKey=uk2:0`
	String6 string `NonUniqueKey=nuk:0`
	IntKey int `UniqueKey=uk1:0`
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
	PtrNoKey *SubNode1
	SliceOfPtr []*Node
	NilSliceOfPtr []*Node
	SliceInt []int
	SliceString []string
	MapStringPtr map[string]*SubNode4
	MapStringPtrNil map[string]*Node
	MapIntString map[int]string
	SubNode1Slice []*SubNode1
	SubNode2Slice []*SubNode2
	SlicePrimary []*SubNode5
	MapPrimary map[string]*SubNode6
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

type SubNode4 struct {
	String string
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*Node
	MaoOfPtrNil map[string]*Node
	MapIntToStr map[int]string
}

type SubNode5 struct {
	String string `PrimaryKey=name:0`
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*Node
	MaoOfPtrNil map[string]*Node
	MapIntToStr map[int]string
}

type SubNode6 struct {
	String string `PrimaryKey=name:0`
	IntSlice []int
	IntSliceNil []int
	StringSlice []string
	MapOfPtr map[string]*Node
	MaoOfPtrNil map[string]*Node
	MapIntToStr map[int]string
}

func createSubChild(loc int) []*Node {
	result:=make([]*Node,4)
	for i:=0;i<4;i++ {
		n:=&Node{}
		n.String = strconv.Itoa(loc)+"-Sub-Child-"+strconv.Itoa(i)
		n.SliceInt = make([]int,3)
		n.SliceInt[1]=544
		n.SliceString = make([]string,3)
		n.SliceString[1]="S1"
		n.MapStringPtr = make(map[string]*SubNode4)
		n.MapStringPtr["B"]=&SubNode4{}
		n.MapStringPtr["B"].String = "str"
		result[i] = n
	}
	return result
}

func createSubNodes1(loc int) []*SubNode1 {
	result:=make([]*SubNode1,3)
	for i:=0;i<3;i++ {
		result[i] = &SubNode1{}
		result[i].String = "SubNode1-"+strconv.Itoa(i)+"-"+strconv.Itoa(loc)
	}
	return result
}

func InitTestModel(size int) []*Node {
	result:=make([]*Node,size)
	for i:=0;i<size;i++ {
		n:=&Node{}
		n.String = "String-"+strconv.Itoa(i)
		n.Int = -101+i
		n.Int32 = -102+int32(i)
		n.Bool = true
		n.Int64 = -103+int64(i)
		n.Uint = 104+uint(i)
		n.Uint32 = 105
		n.Uint64 = 106
		n.Float32 = -107.23
		n.Float64 = -108.25
		n.Ptr = &Node{}
		n.Ptr.String = "OnlyChild-"+n.String
		n.PtrNoKey = &SubNode1{}
		n.PtrNoKey.String = "NoKey-"+n.String
		n.SliceOfPtr = createSubChild(i)
		n.SliceInt = make([]int,5)
		n.SliceInt[3] = 104
		n.SliceString = make([]string,7)
		n.SliceString[3]="303"
		n.MapStringPtr = make(map[string]*SubNode4)
		for _,child:=range n.SliceOfPtr {
			n.MapStringPtr[child.String]=&SubNode4{}
		}
		n.MapIntString=make(map[int]string)
		n.MapIntString[3] = "3"

		n.SubNode1Slice = createSubNodes1(i)

		result[i] = n
	}
	return result
}