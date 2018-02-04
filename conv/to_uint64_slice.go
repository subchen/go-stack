package conv

import (
	"fmt"
	"reflect"
)

func AsUint64Slice(values interface{}) []uint64 {
	v, _ := toUint64Slice(values)
	return v
}

func ToUint64Slice(values interface{}) ([]uint64, error) {
	return toUint64Slice(values)
}

func toUint64Slice(values interface{}) ([]uint64, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind() != reflect.Array && rvalue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]uint64, len)
	var err error
	for i := 0; i < len; i++ {
		results[i], err = toUint64(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
