package assert

import (
	"reflect"
	"testing"
	"time"
)

// Empty asserts that the specified object is empty.  I.e. nil, "", false, 0 or either
// a slice or a channel with len == 0.
//
//  assert.Empty(t, obj)
//
func Empty(t *testing.T, object interface{}) {
	if !isEmpty(object) {
		fail(t, "Expected: empty\nActual  : %#v", object)
	}
}

// array, slice, map, channel, string with len == 0.
//
//   assert.Equal(t, "two", obj[1])
//
func NotEmpty(t *testing.T, object interface{}) {
	if isEmpty(object) {
		fail(t, "Expected: NOT empty\nActual  : %v", object)
	}
}

var numericZeros = []interface{}{
	int(0),
	int8(0),
	int16(0),
	int32(0),
	int64(0),
	uint(0),
	uint8(0),
	uint16(0),
	uint32(0),
	uint64(0),
	float32(0),
	float64(0),
}

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object interface{}) bool {
	if object == nil {
		return true
	} else if object == "" {
		return true
	} else if object == false {
		return true
	}

	for _, v := range numericZeros {
		if object == v {
			return true
		}
	}

	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		{
			return objValue.Len() == 0
		}
	case reflect.Struct:
		switch object.(type) {
		case time.Time:
			return object.(time.Time).IsZero()
		}
	case reflect.Ptr:
		{
			if objValue.IsNil() {
				return true
			}
			switch object.(type) {
			case *time.Time:
				return object.(*time.Time).IsZero()
			default:
				return false
			}
		}
	}
	return false
}
