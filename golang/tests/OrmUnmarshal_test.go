package tests

import (
	. "github.com/saichler/utils/golang"
	. "github.com/saichler/utils/golang/orm/common"
	. "github.com/saichler/utils/golang/orm/transaction"
	"strconv"
	"testing"
)

func TestUnMarshalCountOnlyTopLevel(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
}

func TestUnMarshalCountAllLevels(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",false)
	instances:=m.UnMarshal(q)
	if len(instances)!=30 {
		t.Fail()
		Error("Expected:30 but got "+strconv.Itoa(len(instances)))
	}
}

func TestUnMarshalString(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		expected:="String-"+strconv.Itoa(i)
		found:=false
		for _,n:=range instances {
			node:=n.(*Node)
			if node.String==expected {
				found=true
				break
			}
		}
		if !found {
			t.Fail()
			Error("Failed to find String "+expected)
		}
	}
}

func TestUnMarshalInt(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		expected:=-101+i
		found:=false
		for _,n:=range instances {
			node:=n.(*Node)
			if node.Int==expected {
				found=true
				break
			}
		}
		if !found {
			t.Fail()
			Error("Failed to find Int "+strconv.Itoa(expected))
		}
	}
}

func TestUnMarshalIntSlice(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		expected:=104
		found:=false
		for _,n:=range instances {
			node:=n.(*Node)
			if node.SliceInt==nil {
				t.Fail()
				Error("Expected int slice to exist")
			} else if len(node.SliceInt)!=5 {
				t.Fail()
				Error("Expected int slice of size 4 but got "+strconv.Itoa(len(node.SliceInt)))
			} else if  node.SliceInt[3]==expected {
				found=true
			}
		}
		if !found {
			t.Fail()
			Error("Failed to find Int in slice "+strconv.Itoa(expected))
		}
	}
}

func TestUnMarshalStringSlice(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		expected:="303"
		found:=false
		for _,n:=range instances {
			node:=n.(*Node)
			if node.SliceInt==nil {
				t.Fail()
				Error("Expected int slice to exist")
			} else if len(node.SliceInt)!=5 {
				t.Fail()
				Error("Expected int slice of size 4 but got "+strconv.Itoa(len(node.SliceInt)))
			} else if  node.SliceString[3]==expected {
				found=true
			}
		}
		if !found {
			t.Fail()
			Error("Failed to find string in slice "+expected)
		}
	}
}

func TestUnMarshalPtrKey(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for i:=0;i<size;i++ {
		for _,n:=range instances {
			node:=n.(*Node)
			if node.Ptr==nil {
				t.Fail()
				Error("Expected ptr to exist")
			} else if node.Ptr.String=="" {
				t.Fail()
				Error("Expected ptr name not to be blank ")
			}
		}
	}
}

func TestUnMarshalPtrSliceNoKey(t *testing.T) {
	tx:=&Transaction{}
	m:=initMarshaler(size,tx)
	q:=NewQuery("Node",true)
	instances:=m.UnMarshal(q)
	if len(instances)!=size {
		t.Fail()
		Error("Expected:"+strconv.Itoa(size)+" but got "+strconv.Itoa(len(instances)))
	}
	for _,n:=range instances {
		node:=n.(*Node)
		if node.SubNode1Slice==nil {
			t.Fail()
			Error("Expected ptr slice to exist")
		} else if len(node.SubNode1Slice)!=3 {
			t.Fail()
			Error("Expected int slice of size 4 but got "+strconv.Itoa(len(node.SliceInt)))
		} else {
			for _,sn:=range node.SubNode1Slice {
				if sn==nil {
					t.Fail()
					Error("Nil Entry in slice")
				} else if sn.String=="" {
					t.Fail()
					Error("Expected String to not be blank")
				}
			}
		}
	}
}