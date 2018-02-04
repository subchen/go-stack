package conv

import (
	"fmt"
	"reflect"
)

func AsIntSlice(values interface{}) []int {
	v, _ := toIntSlice(values)
	return v
}

func ToIntSlice(values interface{}) ([]int, error) {
	return toIntSlice(values)
}

func toIntSlice(values interface{}) ([]int, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]int, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toInt(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
