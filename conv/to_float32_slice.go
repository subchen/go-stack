package conv

import (
	"fmt"
	"reflect"
)

func AsFloat32Slice(values interface{}) []float32 {
	v, _ := toFloat32Slice(values)
	return v
}

func ToFloat32Slice(values interface{}) ([]float32, error) {
	return toFloat32Slice(values)
}

func toFloat32Slice(values interface{}) ([]float32, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]float32, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toFloat32(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
