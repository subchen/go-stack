package conv

import (
	"fmt"
)

func AsInterfaceInterfaceMap(value interface{}) map[interface{}]interface{} {
	v, _ := toInterfaceInterfaceMap(value)
	return v
}

func ToInterfaceInterfaceMap(value interface{}) (map[interface{}]interface{}, error) {
	return toInterfaceInterfaceMap(value)
}

func toInterfaceInterfaceMap(value interface{}) (map[interface{}]interface{}, error) {
	if value == nil {
		return nil, nil
	}

	switch val := value.(type) {
	case map[interface{}]interface{}:
		return val, nil
	case map[string]string:
		m := make(map[interface{}]interface{}, len(val))
		for k, v := range val {
			m[k] = v
		}
		return m, nil
	case map[string]interface{}:
		m := make(map[interface{}]interface{}, len(val))
		for k, v := range val {
			m[k] = v
		}
		return m, nil
	default:
		return nil, fmt.Errorf("unsupport convert map[interface{}]interface{} from type(%T)", value)
	}
}
