package assert

import (
	"testing"
)

func TestEqual(t *testing.T) {
	Equal(t, nil, nil)
	Equal(t, "", "")
	Equal(t, 1, 1)
	Equal(t, uint64(0), uint64(0))
	Equal(t, []string{"1"}, []string{"1"})
}

func TestNotEqual(t *testing.T) {
	NotEqual(t, "123", nil)
	NotEqual(t, 123, uint(123))
	NotEqual(t, "123", 123)
	NotEqual(t, []string{"123"}, &[]string{"123"})
}

func TestEqualVal(t *testing.T) {
	EqualVal(t, nil, nil)
	EqualVal(t, "", "")
	EqualVal(t, 1, 1)
	EqualVal(t, 123, uint(123))
	EqualVal(t, 123, float32(123))
}

func TestNotEqualVal(t *testing.T) {
	NotEqualVal(t, "123", nil)
	NotEqualVal(t, "123", 123)
	NotEqualVal(t, []string{"123"}, &[]string{"123"})
}
