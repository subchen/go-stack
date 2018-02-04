package conv

import (
	"fmt"
	"reflect"
)

func AsInt64Slice(values interface{}) []int64 {
	v, _ := toInt64Slice(values)
	return v
}

func ToInt64Slice(values interface{}) ([]int64, error) {
	return toInt64Slice(values)
}

func toInt64Slice(values interface{}) ([]int64, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]int64, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toInt64(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
