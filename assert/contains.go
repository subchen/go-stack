package assert

import (
	"reflect"
	"strings"
	"testing"
)

// Support array, slice, map, chan, string
func Contains(t *testing.T, list interface{}, element interface{}) {
	ok, found := containsElement(list, element)
	if !ok {
		fail(t, "Could not be applied Contains(list: %#v, element: %#v)", list, element)
	}
	if !found {
		fail(t, "List (%#v) does not contain element %#v", list, element)
	}
}

// Support array, slice, map, chan, string
func NotContains(t *testing.T, list interface{}, element interface{}) {
	ok, found := containsElement(list, element)
	if !ok {
		fail(t, "Could not be applied Contains(list: %#v, element: %#v)", list, element)
	}
	if found {
		fail(t, "List (%#v) contains element %#v", list, element)
	}
}

// containsElement try loop over the list check if the list includes the element.
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func containsElement(list interface{}, element interface{}) (ok, found bool) {
	listValue := reflect.ValueOf(list)
	elementValue := reflect.ValueOf(element)

	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()

	if reflect.TypeOf(list).Kind() == reflect.String {
		return true, strings.Contains(listValue.String(), elementValue.String())
	}

	if reflect.TypeOf(list).Kind() == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := 0; i < len(mapKeys); i++ {
			if isEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := 0; i < listValue.Len(); i++ {
		if isEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false
}
