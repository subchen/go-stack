package conv

import (
	"fmt"
)

func AsStringStringMap(value interface{}) map[string]string {
	v, _ := toStringStringMap(value)
	return v
}

func ToStringStringMap(value interface{}) (map[string]string, error) {
	return toStringStringMap(value)
}

func toStringStringMap(value interface{}) (map[string]string, error) {
	if value == nil {
		return nil, nil
	}

	switch val := value.(type) {
	case map[string]string:
		return val, nil
	case map[string]interface{}:
		m := make(map[string]string, len(val))
		for k, v := range val {
			m[AsString(k)] = v
		}
		return m, nil
	case map[interface{}]interface{}:
		m := make(map[string]string, len(val))
		for k, v := range val {
			m[AsString(k)] = AsString(v)
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unsupport convert map[string]string from type(%T)", value)
	}
}
