package assert

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
)

func TestSameType(t *testing.T) {
	SameType(t, 1, 0)
	SameType(t, errors.New("test"), errors.New("test"))
	SameType(t, "x", "y")
	SameType(t, func(int) {}, func(int) {})
}

func TestImplements(t *testing.T) {
	Implements(t, errors.New("test"), (*error)(nil))
	Implements(t, bytes.NewBufferString(""), (*fmt.Stringer)(nil))
}
