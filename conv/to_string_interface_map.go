package conv

import (
	"fmt"
)

func AsStringInterfaceMap(value interface{}) map[string]interface{} {
	v, _ := toStringInterfaceMap(value)
	return v
}

func ToStringInterfaceMap(value interface{}) (map[string]interface{}, error) {
	return toStringInterfaceMap(value)
}

func toStringInterfaceMap(value interface{}) (map[string]interface{}, error) {
	if value == nil {
		return nil, nil
	}

	switch val := value.(type) {
		case map[string]interface{}:
			return val, nil
		case map[interface{}]interface{}:
			m := make(map[string]interface{}, len(val))
			for k, v := range val {
				m[AsString(k)] = v
			}
			return m
		default:
			return nil, fmt.Errorf("unsupport convert map[string]interface{} from type(%T)", value)
	}
}
