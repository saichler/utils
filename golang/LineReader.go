package utils

type LineReader struct {
	data       string
	pos        int
	lineNumber int
}

const (
	EOF = "eof;"
)

func NewLineReader(data string) *LineReader {
	l := &LineReader{}
	l.data = data
	l.pos = 0
	l.lineNumber = 0
	return l
}

func (l *LineReader) NextLine() string {
	if l.pos >= len(l.data) {
		return EOF
	}
	line := ""
	for ; l.pos < len(l.data) && l.data[l.pos:l.pos+1] != "\n"; {
		line += l.data[l.pos : l.pos+1]
		l.pos++
	}
	l.pos++
	l.lineNumber++
	return line
}

func (l *LineReader) LineNumber() int {
	return l.lineNumber
}
