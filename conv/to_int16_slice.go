package conv

import (
	"fmt"
	"reflect"
)

func AsInt16Slice(values interface{}) []int16 {
	v, _ := toInt16Slice(values)
	return v
}

func ToInt16Slice(values interface{}) ([]int16, error) {
	return toInt16Slice(values)
}

func toInt16Slice(values interface{}) ([]int16, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]int16, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toInt16(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
