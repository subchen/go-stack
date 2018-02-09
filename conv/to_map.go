package conv

import (
	"reflect"
)

func convertToMap(smValue reflect.Value, dmType reflect.Type) (interface{}, error) {
	// src map value
	if smValue.Kind() != reflect.Map {
		return nil, fmt.Errorf("src value type is not a map, its type is %s", smValue.Kind().String())
	}
	// dest map type
	if dmType.Kind() != reflect.Map {
		return nil, fmt.Errorf("dest type is not a map, its type is %T", dmType.Kind().String())
	}

	dmValue := reflect.MakeMapWithSize(dmType, smValue.Len())

	for _, skValue := range smValue.MapKeys() {
		svValue := smValue.MapIndex(skValue)
		
		dkValue, err := convertTo(skValue, dmType.Key()) // key
		if err != nil {
			return nil, err
		}
		dvValue, err := convertTo(svValue, dmType.Elem()) // value
		if err != nil {
			return nil, err
		}

		dmValue.SetMapIndex(dkValue, dvValue)
	}

	return dmValue.Interface(), nil
}
