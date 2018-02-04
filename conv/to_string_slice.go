package conv

import (
	"fmt"
	"reflect"
)

func AsStringSlice(values interface{}) []string {
	v, _ := toStringSlice(values)
	return v
}

func ToStringSlice(values interface{}) ([]string, error) {
	return toStringSlice(values)
}

func toStringSlice(values interface{}) ([]string, error) {
	if values == nil {
		return nil, nil
	}

	rvalue := reflect.ValueOf(values)
	if rvalue.Kind != reflect.Array && rvalue.Kind != reflect.Slice {
		return nil, fmt.Errorf("values is not an array or slice, its type is %T", values)
	}

	len := rvalue.Len()
	results := make([]string, len)
	var err error
	for i:=0; i<len; i++ {
		results[i], err = toString(rvalue.Index(i))
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}
