package assert

import (
	"reflect"
	"testing"
)

// Equal asserts that two objects are equal.
//
//    assert.Equal(t, 123, 456)
//
func Equal(t *testing.T, actual, expected interface{}) {
	if !isEqual(actual, expected) {
		fail(t, "Expected: %#v\nActual  : %#v", expected, actual)
	}
}

// NotEqual asserts that the specified values are NOT equal.
//
//    assert.NotEqual(t, obj1, obj2)
//
func NotEqual(t *testing.T, actual, expected interface{}) {
	if isEqual(actual, expected) {
		fail(t, "Expected: NOT equal %#v\nActual  : %#v", expected, actual)
	}
}

// EqualVal asserts that two objects are equal (convertable to the same types and equal).
//
//    assert.EqualVal(t, uint32(123), int32(123))
//
func EqualVal(t *testing.T, actual, expected interface{}) {
	if !isEqualVal(actual, expected) {
		fail(t, "Expected: %#v\nActual  : %#v", expected, actual)
	}
}

// NotEqualVal asserts that two objects are NOT equal(convertable to the same types and equal).
//
//    assert.NotEqualVal(t, uint32(123), int32(123))
//
func NotEqualVal(t *testing.T, actual, expected interface{}) {
	if isEqualVal(actual, expected) {
		fail(t, "Expected: NOT equal %#v\nActual  : %#v", expected, actual)
	}
}

// This function does no assertion of any kind.
func isEqual(actual, expected interface{}) bool {
	if actual == nil || expected == nil {
		return actual == expected
	}

	return reflect.DeepEqual(actual, expected)
}

// ObjectsAreEqualVal gets whether two objects are equal, or if their
// values are equal.
func isEqualVal(actual, expected interface{}) bool {
	if isEqual(actual, expected) {
		return true
	}

	actualType := reflect.TypeOf(actual)
	if actualType == nil {
		return false
	}
	expectedValue := reflect.ValueOf(expected)
	if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
		// Attempt comparison after type conversion
		return reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
	}

	return false
}
