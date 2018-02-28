package query

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/subchen/go-stack/conv"
	"github.com/subchen/go-stack/ss"
)

type Query struct {
	data interface{}
}

var emptyValue = &conv.Value{nil}

func NewQuery(data interface{}) *Query {
	return &Query{data}
}

func (q *Query) Query(expr string) *conv.Value {
	if q.data == nil {
		return emptyValue
	}

	if expr == "." {
		return &conv.Value{q.data}
	}

	ctx := q.data
	multi := false

	paths := ss.SplitWithQuotes(expr, ".", `',"`, false)
	for _, path := range paths {
		if multi {
			matchedCtx := make([]interface{}, 0)
			for _, c := range ctx.([]interface{}) {
				d, m, err := getAttr(c, path)
				if err != nil {
					return emptyValue
				}
				if d != nil {
					if m {
						matchedCtx = append(matchedCtx, d.([]interface{})...)
					} else {
						matchedCtx = append(matchedCtx, d)
					}
				}
			}
			ctx = matchedCtx
		} else {
			var err error
			ctx, multi, err = getAttr(ctx, path)
			if err != nil {
				return emptyValue
			}
			if ctx == nil {
				return emptyValue
			}
		}
	}

	return &conv.Value{ctx}
}

func getAttr(data interface{}, attr string) (value interface{}, multi bool, err error) {
	if len(attr) >= 3 && strings.HasPrefix(attr, "[") && strings.HasSuffix(attr, "]") {
		// array
		v, ok := data.([]interface{})
		if !ok {
			panic(fmt.Errorf("%s is not an array, its type is %T", attr, v))
		}

		if strings.ContainsRune(attr, '=') {
			// map[key=value]
			kv := strings.SplitN(attr, "=", 2)

			matchedList := make([]interface{}, 0)
			for _, c := range v {
				if cm, ok := c.(map[string]interface{}); ok {
					if val, ok := cm[kv[0]]; ok {
						if fmt.Sprintf("%v", val) == kv[1] {
							matchedList = append(matchedList, cm)
						}
					}
				} else if cm, ok := c.(map[interface{}]interface{}); ok {
					if val, ok := cm[kv[0]]; ok {
						if fmt.Sprintf("%v", val) == kv[1] {
							matchedList = append(matchedList, cm)
						}
					}
				} else {
					panic(fmt.Errorf("%s is not an array, its type is %T", attr, c))
				}
			}
			return matchedList, true, nil
		} else if strings.ContainsRune(attr, ':') {
			// slice[low:high]
			low := 0
			high := len(v)
			var err error
			indexes := strings.SplitN(attr, ":", 2)
			if indexes[0] != "" {
				low, err = strconv.Atoi(indexes[0])
				if err != nil {
					panic(err)
				}
			}
			if indexes[1] != "" {
				high, err = strconv.Atoi(indexes[1])
				if err != nil {
					panic(err)
				}
			}
			if low < 0 || low >= len(v) {
				return nil, true, fmt.Errorf("%s: low index out of range", attr)
			}
			if high <= 0 {
				high = len(v) + high
			}
			if high < 0 || high > len(v) {
				return nil, true, fmt.Errorf("%s: high index out of range", attr)
			}
			if high <= low {
				return nil, true, fmt.Errorf("%s: high index must be large than low index", attr)
			}
			return v[low:high], true, nil
		} else {
			// array[i]
			index, err := strconv.Atoi(attr[1 : len(attr)-1])
			if err != nil {
				panic(err)
			}
			if index < 0 || index >= len(v) {
				return nil, false, fmt.Errorf("%s: index out of range", attr)
			}
			return v[index], false, nil
		}
	} else {
		// map
		if v, ok := data.(map[string]interface{}); ok {
			if val, ok := v[attr]; ok {
				return val, false, nil
			}
			return nil, false, nil
		} else if v, ok := data.(map[interface{}]interface{}); ok {
			if val, ok := v[attr]; ok {
				return val, false, nil
			}
			return nil, false, nil
		} else {
			panic(fmt.Errorf("%s is not an array, its type is %T", attr, v))
		}
	}
}
