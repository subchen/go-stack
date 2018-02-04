package conv

import (
	"fmt"
	"reflect"
)

func AsBoolSlice(values interface{}) []bool {
	v, _ := toBoolSlice(values)
	return v
}

func ToBoolSlice(values interface{}) ([]bool, error) {
	return toBoolSlice(values)
}

func toBoolSlice(values interface{}) ([]bool, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind() != reflect.Array && rvalue.Kind() != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]bool, len)
	var err error
	for i := 0; i < len; i++ {
		results[i], err = toBool(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
