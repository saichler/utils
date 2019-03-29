package utils

import (
	"reflect"
	"strconv"
)

var tostrings = make(map[reflect.Kind]func(reflect.Value)string)
var tostringinit = false
func initToStrings() {
	if !tostringinit {
		tostringinit = true
		tostrings[reflect.String] = stringToString
		tostrings[reflect.Int] = intToString
		tostrings[reflect.Int16] = intToString
		tostrings[reflect.Int32] = intToString
		tostrings[reflect.Int64] = intToString
		tostrings[reflect.Uint] = uintToString
		tostrings[reflect.Uint16] = uintToString
		tostrings[reflect.Uint32] = uintToString
		tostrings[reflect.Uint64] = uintToString
		tostrings[reflect.Float32] = floatToString
		tostrings[reflect.Float64] = floatToString
		tostrings[reflect.Bool] = boolToString
		tostrings[reflect.Ptr] = ptrToString
		tostrings[reflect.Slice] = sliceToString
		tostrings[reflect.Map] = mapToString
		tostrings[reflect.Interface] = interfaceToString
	}
}

func ToString(value reflect.Value) string {
	if !value.IsValid(){
		return ""
	}
	initToStrings()
	tostring:=tostrings[value.Kind()]
	if tostring==nil {
		panic("No ToString for kind:"+value.Kind().String())
	}
	return tostring(value)
}

func stringToString(value reflect.Value) string {
	return value.String()
}

func intToString(value reflect.Value) string {
	return strconv.Itoa(int(value.Int()))
}

func uintToString(value reflect.Value) string {
	return strconv.Itoa(int(value.Uint()))
}

func floatToString(value reflect.Value) string {
	return strconv.FormatFloat(float64(value.Float()),'f', -1, 64)
}

func boolToString(value reflect.Value) string {
	if value.Bool() {
		return "true"
	} else {
		return "false"
	}
}

func ptrToString(value reflect.Value) string {
	if value.IsNil(){
		return ""
	}
	return ToString(value.Elem())
}

func sliceToString(value reflect.Value) string {
	if value.Len()==0 {
		return "[]"
	}
	result:=NewStringBuilder("[")
	for i:=0;i<value.Len();i++ {
		if i!=0 {
			result.Append(",")
		}
		elem:=value.Index(i)
		result.Append(ToString(elem))
	}
	result.Append("]")
	return result.String()
}

func mapToString(value reflect.Value) string {
	mapkeys:=value.MapKeys()
	if len(mapkeys)==0 {
		return "[]"
	}
	result:=NewStringBuilder("[")
	for i,key:=range mapkeys {
		if i!=0 {
			result.Append(",")
		}
		val:=value.MapIndex(key)
		result.Append(ToString(key))
		result.Append("=")
		result.Append(ToString(val))
	}
	result.Append("]")
	return result.String()
}

func interfaceToString(value reflect.Value) string {
	return ToString(value.Elem())
}