package assert

import (
	"reflect"
	"testing"
)

// Support array, slice, map, chan, string
func Len(t *testing.T, object interface{}, length int) {
	l, ok := getLen(object)
	if !ok {
		fail(t, "Could not be applied Len(): %#v", object)
	}
	if l != length {
		fail(t, "%#v should have %d item(s), but has %d", object, length, l)
	}
}

// getLen try to get length of object.
// return (false, 0) if impossible.
func getLen(x interface{}) (length int, ok bool) {
	v := reflect.ValueOf(x)
	defer func() {
		if e := recover(); e != nil {
			ok = false
		}
	}()
	return v.Len(), true
}
