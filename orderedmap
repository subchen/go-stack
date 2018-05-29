package orderedmap

// OrderedMap is a map where the keys keep the order that they're added.
// It also can be de/serialized from/to JSON
type OrderedMap struct {
	keys   []string
	values map[string]interface{}
}

// New creates a new OrderedMap instrance
func New() *OrderedMap {
	return &OrderedMap{
		keys:   make([]string, 0, 8),
		values: make(map[string]interface{}, 8),
	}
}

// Get gets the value associated with the given key
func (o *OrderedMap) Get(key string) (interface{}, bool) {
	val, exists := o.values[key]
	return val, exists
}

// Set sets the associated value with given key
func (o *OrderedMap) Set(key string, value interface{}) {
	_, exists := o.values[key]
	if !exists {
		o.keys = append(o.keys, key)
	}
	o.values[key] = value
}

// Del deletes the values associated with key.
func (o *OrderedMap) Del(key string) {
	// check key is in use
	_, ok := o.values[key]
	if !ok {
		return
	}
	// remove from keys
	for i, k := range o.keys {
		if k == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
			break
		}
	}
	// remove from values
	delete(o.values, key)
}

// Keys returns ordered keys
func (o *OrderedMap) Keys() []string {
	return o.keys
}

// Len returns the number of map
func (o *OrderedMap) Len() int {
	return len(o.keys)
}
