package conv

import (
	"fmt"
	"reflect"
)

func AsUint32Slice(values interface{}) []uint32 {
	v, _ := toUint32Slice(values)
	return v
}

func ToUint32Slice(values interface{}) ([]uint32, error) {
	return toUint32Slice(values)
}

func toUint32Slice(values interface{}) ([]uint32, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]uint32, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toUint32(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
