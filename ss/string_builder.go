package ss

import (
	"bytes"
	"fmt"
	"io"
)

type StringBuilder struct {
	buf bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
	return new(StringBuilder)
}

func (sb *StringBuilder) Write(s string) {
	sb.buf.WriteString(s)
}

func (sb *StringBuilder) Writeln(s string) {
	sb.buf.WriteString(s)
	sb.buf.WriteByte('\n')
}

func (sb *StringBuilder) Writef(format string, arg ...interface{}) {
	sb.buf.WriteString(fmt.Sprintf(format, arg...))
}

func (sb *StringBuilder) WriteByte(b byte) {
	sb.buf.WriteByte(b)
}

func (sb *StringBuilder) WriteRune(r rune) {
	sb.buf.WriteRune(r)
}

func (sb *StringBuilder) WriteInt(n int) {
	sb.buf.WriteString(fmt.Sprintf("%d", n))
}

func (sb *StringBuilder) WriteInt64(n int64) {
	sb.buf.WriteString(fmt.Sprintf("%d", n))
}

func (sb *StringBuilder) WriteTo(w io.Writer) (n int64, err error) {
	return sb.buf.WriteTo(w)
}

func (sb *StringBuilder) Reset() {
	sb.buf.Reset()
}

func (sb *StringBuilder) Len() int {
	return sb.buf.Len()
}

func (sb *StringBuilder) String() string {
	return sb.buf.String()
}

func (sb *StringBuilder) Bytes() []byte {
	return sb.buf.Bytes()
}
