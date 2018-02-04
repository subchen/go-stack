package conv

import (
	"fmt"
	"reflect"
)

func AsUint8Slice(values interface{}) []uint8 {
	v, _ := toUint8Slice(values)
	return v
}

func ToUint8Slice(values interface{}) ([]uint8, error) {
	return toUint8Slice(values)
}

func toUint8Slice(values interface{}) ([]uint8, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind() != reflect.Array && rvalue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]uint8, len)
	var err error
	for i := 0; i < len; i++ {
		results[i], err = toUint8(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
