package conv

import (
	"fmt"
	"reflect"
)

func AsInt8Slice(values interface{}) []int8 {
	v, _ := toInt8Slice(values)
	return v
}

func ToInt8Slice(values interface{}) ([]int8, error) {
	return toInt8Slice(values)
}

func toInt8Slice(values interface{}) ([]int8, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]int8, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toInt8(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
