package utils

type Serializer interface {
	ToBytes() []byte
	FromBytes([]byte)
	Write(*ByteSlice)
	Read(*ByteSlice)
}
