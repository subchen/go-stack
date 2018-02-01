package assert

import (
	"reflect"
	"testing"
)

// Implements asserts that an object is implemented by the specified interface.
//
//    assert.Implements(t, new(MyObject), (*MyInterface)(nil))
//
func Implements(t *testing.T, object interface{}, interfaceObject interface{}) {
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if !reflect.TypeOf(object).Implements(interfaceType) {
		fail(t, "%T must implement %v", object, interfaceType)
	}
}

// SameType asserts that the specified objects are of the same type.
//
//    assert.SameType(t, "1", "1")
//
func SameType(t *testing.T, object interface{}, anotherObject interface{}) {
	if !isEqual(reflect.TypeOf(object), reflect.TypeOf(anotherObject)) {
		fail(t, "Object expected to be of type %v, but was %v", reflect.TypeOf(anotherObject), reflect.TypeOf(object))
	}
}
