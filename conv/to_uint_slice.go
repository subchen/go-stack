package conv

import (
	"fmt"
	"reflect"
)

func AsUintSlice(values interface{}) []uint {
	v, _ := toUintSlice(values)
	return v
}

func ToUintSlice(values interface{}) ([]uint, error) {
	return toUintSlice(values)
}

func toUintSlice(values interface{}) ([]uint, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind() != reflect.Array && rvalue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]uint, len)
	var err error
	for i := 0; i < len; i++ {
		results[i], err = toUint(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
