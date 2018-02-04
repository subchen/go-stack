package conv

import (
	"fmt"
	"reflect"
)

func AsUint16Slice(values interface{}) []uint16 {
	v, _ := toUint16Slice(values)
	return v
}

func ToUint16Slice(values interface{}) ([]uint16, error) {
	return toUint16Slice(values)
}

func toUint16Slice(values interface{}) ([]uint16, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind() != reflect.Array && rvalue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]uint16, len)
	var err error
	for i := 0; i < len; i++ {
		results[i], err = toUint16(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
