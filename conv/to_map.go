package conv

import (
	"reflect"
)

func convertToMap(rvalue reflect.Value, rtype reflect.Type) (interface{}, error) {
	if rvalue.Type() == rtype {
		// same type, nothing to convert.
		return rvalue.Interface(), nil
	}

	// src map value
	if rvalue.Kind() != reflect.Map {
		return nil, fmt.Errorf("src value type is not a map, its type is %s", rvalue.Kind().String())
	}
	// dest map type
	if rtype.Kind() != reflect.Map {
		return nil, fmt.Errorf("dest type is not a map, its type is %s", rtype.Kind().String())
	}

	m := reflect.MakeMapWithSize(rtype, rvalue.Len())

	for _, key := range rvalue.MapKeys() {
		value := rvalue.MapIndex(key)

		newkey, err := convertTo(key, rtype.Key())
		if err != nil {
			return nil, err
		}
		newvalue, err := convertTo(value, rtype.Elem())
		if err != nil {
			return nil, err
		}

		m.SetMapIndex(reflect.ValueOf(newkey), reflect.ValueOf(newvalue))
	}

	return m.Interface(), nil
}
