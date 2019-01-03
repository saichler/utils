package utils

const (
	size=600
)

type ByteArray struct {
	data [size]byte
	loc int
}

func (ba *ByteArray) Init(data []byte) {
	copy(ba.data[0:],data[0:])
	ba.loc=0
}

func (ba *ByteArray) Reset() {
	ba.loc=0
}

func (ba *ByteArray) AddByteArray(data []byte){
	ba.AddUInt32(uint32(len(data)))
	for i:=0;i<len(data);i++ {
		ba.data[ba.loc+i]=data[i]
	}
	ba.loc+=len(data)
}

func (ba *ByteArray) GetByteArray() []byte {
	size := int(ba.GetUInt32())
	result := ba.data[ba.loc:ba.loc+size]
	ba.loc+=size
	return result
}

func (ba *ByteArray) AddString(str string){
	data := []byte(str)
	ba.AddUInt32(uint32(len(data)))
	for i:=0;i<len(data);i++ {
		ba.data[ba.loc+i]=data[i]
	}
	ba.loc+=len(data)
}

func (ba *ByteArray) GetString() string {
	size := int(ba.GetUInt32())
	result := ba.data[ba.loc:ba.loc+size]
	ba.loc+=size
	return string(result)
}

func (ba *ByteArray) AddInt64(i64 int64) {
	ba.addInt64(i64)
	ba.loc+=8
}

func (ba *ByteArray) GetInt64() int64 {
	result := ba.getInt64()
	ba.loc+=8;
	return result
}

func (ba *ByteArray) AddInt(i int) {
	ba.AddInt32(int32(i))
}

func (ba *ByteArray) GetInt() int {
	return int(ba.GetInt32())
}

func (ba *ByteArray) AddInt32(i32 int32) {
	ba.addInt32(i32)
	ba.loc+=4
}

func (ba *ByteArray) GetInt32() int32 {
	result := ba.getInt32()
	ba.loc+=4;
	return result
}

func (ba *ByteArray) AddUInt16(i16 uint16) {
	ba.addUInt16(i16)
	ba.loc+=2
}

func (ba *ByteArray) GetUInt16() uint16 {
	result := ba.getUInt16()
	ba.loc+=2;
	return result
}

func (ba *ByteArray) AddUInt32(i32 uint32) {
	ba.addInt32(int32(i32))
	ba.loc+=4
}

func (ba *ByteArray) GetUInt32() uint32 {
	result := uint32(ba.getInt32())
	ba.loc+=4;
	return result
}

func (ba *ByteArray) AddBool(b bool){
	sb := byte(0)
	if b {
		sb = 1
	}
	ba.data[ba.loc]=sb
	ba.loc+=1
}

func (ba *ByteArray) GetBool() bool {
	var result bool = false
	if ba.data[ba.loc] == 1 {
		result = true
	}
	ba.loc+=1
	return result
}

func NewByteArray() *ByteArray {
	ba := &ByteArray{}
	ba.loc = 0
	return ba
}

func NewByteArrayWithData(data []byte,loc int) *ByteArray {
	ba := &ByteArray{}
	copy(ba.data[0:],data)
	ba.loc = loc
	return ba
}

func (ba *ByteArray) Data()[]byte {
	return ba.data[0:ba.loc]
}

func (ba *ByteArray) Loc() int {
	return ba.loc
}

func (ba *ByteArray) Put(key, value []byte) {
	ba.AddByteArray(key)
	ba.AddByteArray(value)
}

func (ba *ByteArray) Get() ([]byte,[]byte) {
	key := ba.GetByteArray()
	value := ba.GetByteArray()
	return key, value
}

func (ba *ByteArray) IsEOF() bool {
	return ba.loc==len(ba.data)
}

func (ba *ByteArray) addInt64(num int64) {
	ba.data[ba.loc+0] = byte(num)
	ba.data[ba.loc+1] = byte(num >> 8)
	ba.data[ba.loc+2] = byte(num >> 16)
	ba.data[ba.loc+3] = byte(num >> 24)
	ba.data[ba.loc+4] = byte(num >> 32)
	ba.data[ba.loc+5] = byte(num >> 40)
	ba.data[ba.loc+6] = byte(num >> 48)
	ba.data[ba.loc+7] = byte(num >> 56)
}

func (ba *ByteArray) getInt64() int64 {
	return int64(ba.data[ba.loc+0]) |
		   int64(ba.data[ba.loc+1])<<8 |
		   int64(ba.data[ba.loc+2])<<16 |
		   int64(ba.data[ba.loc+3])<<24 |
		   int64(ba.data[ba.loc+4])<<32 |
		   int64(ba.data[ba.loc+5])<<40 |
		   int64(ba.data[ba.loc+6])<<48 |
		   int64(ba.data[ba.loc+7])<<56
}

func (ba *ByteArray) addInt32(num int32) {
	ba.data[ba.loc+0] = byte(num)
	ba.data[ba.loc+1] = byte(num >> 8)
	ba.data[ba.loc+2] = byte(num >> 16)
	ba.data[ba.loc+3] = byte(num >> 24)
}

func (ba *ByteArray) getInt32() int32 {
	return int32(ba.data[ba.loc+0]) |
		   int32(ba.data[ba.loc+1])<<8 |
		   int32(ba.data[ba.loc+2])<<16 |
		   int32(ba.data[ba.loc+3])<<24
}

func (ba *ByteArray) addUInt16(num uint16) {
	ba.data[ba.loc+0] = byte(num)
	ba.data[ba.loc+1] = byte(num >> 8)
}

func (ba *ByteArray) getUInt16() uint16 {
	return uint16(ba.data[ba.loc+0]) |
		   uint16(ba.data[ba.loc+1])<<8
}