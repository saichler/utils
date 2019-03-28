package utils

import "bytes"

type StringBuilder struct {
	buff *bytes.Buffer
}

func NewStringBuilder(str string) *StringBuilder {
	sb := &StringBuilder{}
	sb.buff = &bytes.Buffer{}
	sb.buff.WriteString(str)
	return sb
}

func (sb *StringBuilder) Append(str string) *StringBuilder {
	sb.buff.WriteString(str)
	return sb
}

func (sb *StringBuilder) AppendSB(other *StringBuilder) *StringBuilder {
	sb.buff.Write(other.buff.Bytes())
	return sb
}

func (sb *StringBuilder) String() string {
	return sb.buff.String()
}

func (sb *StringBuilder) Empty() bool {
	return sb.buff.Len() == 0
}
