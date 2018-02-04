package conv

import (
	"fmt"
	"reflect"
)

func AsFloat64Slice(values interface{}) []float64 {
	v, _ := toFloat64Slice(values)
	return v
}

func ToFloat64Slice(values interface{}) ([]float64, error) {
	return toFloat64Slice(values)
}

func toFloat64Slice(values interface{}) ([]float64, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind() != reflect.Array && rvalue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]float64, len)
	var err error
	for i := 0; i < len; i++ {
		results[i], err = toFloat64(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
