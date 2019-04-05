package utils

import (
	"reflect"
	"strconv"
	"strings"
)

var fromStrings = make(map[reflect.Kind]func(string,reflect.Type) reflect.Value)
var fromstringinit = false

func initFromStrings() {
	if !fromstringinit {
		fromStrings[reflect.Int] = intFromString
		fromStrings[reflect.Int32] = int32FromString
		fromStrings[reflect.Ptr] = ptrFromString
		fromStrings[reflect.String] = stringFromString
		fromStrings[reflect.Slice] = sliceFromString
		fromStrings[reflect.Map] = mapFromString
		fromStrings[reflect.Bool] = boolFromString
		fromStrings[reflect.Int64] = int64FromString
		fromStrings[reflect.Uint] = uintFromString
		fromStrings[reflect.Uint32] = uint32FromString
		fromStrings[reflect.Uint64] = uint64FromString
		fromStrings[reflect.Float32] = float32FromString
		fromStrings[reflect.Float64] = float64FromString
	}
}

func FromString(str string,typ reflect.Type) reflect.Value {
	initFromStrings()
	fromString:=fromStrings[typ.Kind()]
	if fromString==nil {
		panic("Cannot find from string for kind:"+typ.Kind().String())
	}
	return fromString(str,typ)
}

func sliceFromString(str string,typ reflect.Type) reflect.Value {
	if str=="" {
		return reflect.ValueOf(nil)
	}
	str=str[1:len(str)-1]
	elements:=strings.Split(str,",")
	newSlice := reflect.MakeSlice(reflect.SliceOf(typ.Elem()), len(elements), len(elements))
	for i,e:=range elements {
		newSlice.Index(i).Set(FromString(e,typ.Elem()))
	}
	return newSlice
}

func ptrFromString(str string,typ reflect.Type) reflect.Value {
	newPtr := reflect.New(typ.Elem())
	elem:=FromString(str,typ.Elem())
	newPtr.Elem().Set(elem)
	return newPtr
}

func mapFromString(str string,typ reflect.Type) reflect.Value {
	str=strings.Trim(str," ")
	if str=="" {
		return reflect.ValueOf(nil)
	}
	str=str[1:len(str)-1]
	pairs:=strings.Split(str,",")
	mapClone := reflect.MakeMapWithSize(typ, len(pairs))
	for _, pair := range pairs {
		index:=strings.Index(pair,"=")
		k:=pair[0:index]
		v:=pair[index+1:]
		key:=FromString(k,typ.Key())
		val:=FromString(v,typ.Elem())
		mapClone.SetMapIndex(key, val)
	}
	return mapClone
}

func intFromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.Atoi(str)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(int(i))
}

func uintFromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.Atoi(str)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(uint(i))
}

func uint32FromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.Atoi(str)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(uint32(i))
}

func uint64FromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.Atoi(str)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(uint64(i))
}

func float32FromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.ParseFloat(str,64)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(float32(i))
}

func float64FromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.ParseFloat(str,64)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(float64(i))
}

func boolFromString(str string,typ reflect.Type) reflect.Value {
	if str=="true" {
		return reflect.ValueOf(true)
	}
	return reflect.ValueOf(false)
}

func int32FromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.Atoi(str)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(int32(i))
}

func int64FromString(str string,typ reflect.Type) reflect.Value {
	i,e:=strconv.Atoi(str)
	if e!=nil {
		panic(e)
	}
	return reflect.ValueOf(int64(i))
}

func stringFromString(str string,typ reflect.Type) reflect.Value {
	return reflect.ValueOf(str)
}

