package assert

import (
	"testing"
)

// Error asserts that the specified error is not nil.
//
//    assert.Error(t, err)
//
func Error(t *testing.T, err interface{}) {
	if err == nil {
		fail(t, "Expected: an error\nActual  : nil")
	}
}

// NoError asserts that the specified error is nil.
//
//    assert.NoError(t, err)
//
func NoError(t *testing.T, err interface{}) {
	if err != nil {
		fail(t, "Expected: no error\nActual  : %#v", err)
	}
}
