package utils

import (
	"encoding/binary"
)

type ByteSlice struct {
	data []byte
	loc  int
}

func (ba *ByteSlice) AddByteSlice(data []byte) {
	ba.AddUInt32(uint32(len(data)))
	ba.data = append(ba.data, data...)
	ba.loc += len(data)
}

func (ba *ByteSlice) AddByte(data byte) {
	ba.data = append(ba.data, data)
	ba.loc++
}

func (ba *ByteSlice) GetByte() byte {
	if ba.loc == len(ba.data) {
		return 0
	}
	result := ba.data[ba.loc]
	ba.loc++
	return result
}

func (ba *ByteSlice) Add(data []byte) {
	ba.data = append(ba.data, data...)
	ba.loc += len(data)
}

func (ba *ByteSlice) GetByteSlice() []byte {
	size := int(ba.GetUInt32())
	result := ba.data[ba.loc : ba.loc+size]
	ba.loc += size
	return result
}

func (ba *ByteSlice) AddString(str string) {
	data := []byte(str)
	ba.AddUInt32(uint32(len(data)))
	ba.data = append(ba.data, data...)
	ba.loc += len(data)
}

func (ba *ByteSlice) GetString() string {
	size := int(ba.GetUInt32())
	result := ba.data[ba.loc : ba.loc+size]
	ba.loc += size
	return string(result)
}

func (ba *ByteSlice) AddInt64(i64 int64) {
	num := make([]byte, 8)
	binary.LittleEndian.PutUint64(num, uint64(i64))
	ba.data = append(ba.data, num...)
	ba.loc += 8
}

func (ba *ByteSlice) GetInt64() int64 {
	result := int64(binary.LittleEndian.Uint64(ba.data[ba.loc : ba.loc+8]))
	ba.loc += 8;
	return result
}

func (ba *ByteSlice) AddInt(i int) {
	ba.AddInt32(int32(i))
}

func (ba *ByteSlice) GetInt() int {
	return int(ba.GetInt32())
}

func (ba *ByteSlice) AddInt32(i32 int32) {
	num := make([]byte, 4)
	binary.LittleEndian.PutUint32(num, uint32(i32))
	ba.data = append(ba.data, num...)
	ba.loc += 4
}

func (ba *ByteSlice) GetInt32() int32 {
	result := int32(binary.LittleEndian.Uint32(ba.data[ba.loc : ba.loc+4]))
	ba.loc += 4;
	return result
}

func (ba *ByteSlice) AddUInt16(i16 uint16) {
	num := make([]byte, 2)
	binary.LittleEndian.PutUint16(num, i16)
	ba.data = append(ba.data, num...)
	ba.loc += 2
}

func (ba *ByteSlice) GetUInt16() uint16 {
	result := binary.LittleEndian.Uint16(ba.data[ba.loc : ba.loc+2])
	ba.loc += 2;
	return result
}

func (ba *ByteSlice) AddUInt32(i32 uint32) {
	num := make([]byte, 4)
	binary.LittleEndian.PutUint32(num, i32)
	ba.data = append(ba.data, num...)
	ba.loc += 4
}

func (ba *ByteSlice) GetUInt32() uint32 {
	result := binary.LittleEndian.Uint32(ba.data[ba.loc : ba.loc+4])
	ba.loc += 4;
	return result
}

func (ba *ByteSlice) AddBool(b bool) {
	sb := byte(0)
	if b {
		sb = 1
	}
	ba.data = append(ba.data, sb)
	ba.loc += 1
}

func (ba *ByteSlice) GetBool() bool {
	var result bool = false
	if ba.data[ba.loc] == 1 {
		result = true
	}
	ba.loc += 1
	return result
}

func NewByteSlice() *ByteSlice {
	ba := &ByteSlice{}
	ba.data = make([]byte, 0)
	ba.loc = 0
	return ba
}

func NewByteSliceWithData(data []byte, loc int) *ByteSlice {
	ba := &ByteSlice{}
	ba.data = data
	ba.loc = loc
	return ba
}

func (ba *ByteSlice) Data() []byte {
	return ba.data[0:ba.loc]
}

func (ba *ByteSlice) Loc() int {
	return ba.loc
}

func (ba *ByteSlice) Put(key, value []byte) {
	ba.AddByteSlice(key)
	ba.AddByteSlice(value)
}

func (ba *ByteSlice) Get() ([]byte, []byte) {
	key := ba.GetByteSlice()
	value := ba.GetByteSlice()
	return key, value
}

func (ba *ByteSlice) IsEOF() bool {
	return ba.loc == len(ba.data)
}

func Encode2BoolAndUInt6(b1, b2 bool, i int) byte {
	if b1 && b2 {
		return (1 << 7) + (1 << 6) + byte(i)
	} else if b1 && !b2 {
		return (1 << 7) + byte(i)
	} else if !b1 && b2 {
		return (1 << 6) + byte(i)
	}
	return byte(i)
}

func Decode2BoolAndUInt6(byt byte) (bool, bool, int) {
	b1 := byt >> 7
	b2 := (byt - b1<<7) >> 6
	i := int(byt - (b1<<7 + b2<<6))
	return b1 == 1, b2 == 1, i
}
