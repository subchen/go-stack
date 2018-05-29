package stringbuilder

import (
	"bytes"
	"fmt"
	"io"
)

// StringBuilder is a buffer for string
type StringBuilder struct {
	buf bytes.Buffer
}

// New create a new StringBuilder instance
func New() *StringBuilder {
	return new(StringBuilder)
}

// Grow grows the buffer's capacity
func (sb *StringBuilder) Grow(n int) {
	sb.buf.Grow(n)
}

// Write appends the contents to the buffer
func (sb *StringBuilder) Write(s string) {
	sb.buf.WriteString(s)
}

// Writeln appends the contents to the buffer, following a new line
func (sb *StringBuilder) Writeln(s string) {
	sb.buf.WriteString(s)
	sb.buf.WriteByte('\n')
}

// Writef appends the formatted contents to the buffer
func (sb *StringBuilder) Writef(format string, arg ...interface{}) {
	sb.buf.WriteString(fmt.Sprintf(format, arg...))
}

// Write writes the contents into the buffer.
func (sb *StringBuilder) WriteByte(b byte) {
	sb.buf.WriteByte(b)
}

// WriteRune writes the contents into the buffer.
func (sb *StringBuilder) WriteRune(r rune) {
	sb.buf.WriteRune(r)
}

// WriteInt writes the contents into the buffer.
func (sb *StringBuilder) WriteInt(n int) {
	sb.buf.WriteString(fmt.Sprintf("%d", n))
}

// WriteInt64 writes the contents into the buffer.
func (sb *StringBuilder) WriteInt64(n int64) {
	sb.buf.WriteString(fmt.Sprintf("%d", n))
}

// WriteUint writes the contents into the buffer.
func (sb *StringBuilder) WriteUint(n uint) {
	sb.buf.WriteString(fmt.Sprintf("%d", n))
}

// WriteUint64 writes the contents into the buffer.
func (sb *StringBuilder) WriteUint64(n uint64) {
	sb.buf.WriteString(fmt.Sprintf("%d", n))
}

// WriteTo implements the io.WriterTo interface.
func (sb *StringBuilder) WriteTo(w io.Writer) (n int64, err error) {
	return sb.buf.WriteTo(w)
}

// Reset discards any buffered data
func (sb *StringBuilder) Reset() {
	sb.buf.Reset()
}

// Len returns the number of bytes in buffer
func (sb *StringBuilder) Len() int {
	return sb.buf.Len()
}

// String returns the contents of buffer
func (sb *StringBuilder) String() string {
	return sb.buf.String()
}

// Bytes returns a slice of length sb.Len()
func (sb *StringBuilder) Bytes() []byte {
	return sb.buf.Bytes()
}
