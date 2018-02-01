package assert

import (
	"reflect"
	"testing"
)

// Zero asserts that object is the zero value for its type and returns the truth.
func Zero(t *testing.T, object interface{}) {
	if object != nil && !reflect.DeepEqual(object, reflect.Zero(reflect.TypeOf(object)).Interface()) {
		fail(t, "Should be zero, but was %#v", object)
	}
}

// NotZero asserts that object is not the zero value for its type and returns the truth.
func NotZero(t *testing.T, object interface{}) {
	if object == nil || reflect.DeepEqual(object, reflect.Zero(reflect.TypeOf(object)).Interface()) {
		fail(t, "Should not be zero, but was %#v", object)
	}
}
