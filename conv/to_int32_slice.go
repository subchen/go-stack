package conv

import (
	"fmt"
	"reflect"
)

func AsInt32Slice(values interface{}) []int32 {
	v, _ := toInt32Slice(values)
	return v
}

func ToInt32Slice(values interface{}) ([]int32, error) {
	return toInt32Slice(values)
}

func toInt32Slice(values interface{}) ([]int32, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]int32, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toInt32(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
