package assert

import (
	"reflect"
	"testing"
)

// Nil asserts that the specified object is nil.
//
//    assert.Nil(t, err)
//
func Nil(t *testing.T, object interface{}) {
	if !isNil(object) {
		fail(t, "Expected: nil\nActual  : %#v", object)
	}
}

// NotNil asserts that the specified object is not nil.
//
//    assert.NotNil(t, err)
//
func NotNil(t *testing.T, object interface{}) {
	if isNil(object) {
		fail(t, "Expected: NOT nil\nActual  : %#v", object)
	}
}

// isNil checks if a specified object is nil or not, without Failing.
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}
