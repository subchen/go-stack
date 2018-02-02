package maps

import (
	"sync"
)

type (
	// InterfaceInterfaceMap is a synchronous map.
	InterfaceInterfaceMap struct {
		data  map[interface{}]interface{}
		mutex sync.RWMutex // used only by writers
	}
)

// NewInterfaceInterfaceMap initializes a new empty map.
// Use of nil to empty the InterfaceInterfaceMap is okay.
func NewInterfaceInterfaceMap() *InterfaceInterfaceMap {
	m := new(InterfaceInterfaceMap)
	m.data = make(map[interface{}]interface{})
	return m
}

// Get retrieves the value associated with the key.
func (m *InterfaceInterfaceMap) Get(key interface{}) interface{} {
	m.mutex.RLock()
	value := m.data[key]
	m.mutex.RUnlock()
	return value
}

// Load retrieves the value associated with the key.
func (m *InterfaceInterfaceMap) Load(key interface{}) (value interface{}, ok bool) {
	m.mutex.RLock()
	value, ok = m.data[key]
	m.mutex.RUnlock()
	return
}

// Set inserts a key-value pair.
func (m *InterfaceInterfaceMap) Set(key interface{}, value interface{}) {
	m.mutex.Lock()
	m.data[key] = value
	m.mutex.Unlock()
}

// SetIfNotPresent inserts a key-value pair if not exists.
func (m *InterfaceInterfaceMap) SetIfNotPresent(key interface{}, value interface{}) {
	m.mutex.Lock()
	if _, ok := m.data[key]; !ok {
		m.data[key] = value
	}
	m.mutex.Unlock()
}

// CopyFrom efficiently inserts all the key-value pairs.
func (m *InterfaceInterfaceMap) CopyFrom(src map[interface{}]interface{}) {
	m.mutex.Lock()
	for k, v := range src {
		m.data[k] = v
	}
	m.mutex.Unlock()
}

// Remove removes key from the InterfaceInterfaceMap.
func (m *InterfaceInterfaceMap) Remove(key interface{}) {
	m.mutex.Lock()
	delete(m.data, key)
	m.mutex.Unlock()
}

// Clear removes all keys from the InterfaceInterfaceMap.
func (m *InterfaceInterfaceMap) Clear() {
	m.mutex.Lock()
	m.data = make(map[interface{}]interface{})
	m.mutex.Unlock()
}

// Size returns map size
func (m *StringInterfaceMap) Size() int {
	m.mutex.RLock()
	size := len(m.data)
	m.mutex.RUnlock()
	return size
}

// IsEmpty returns true if map is empty
func (m *StringInterfaceMap) IsEmpty() bool {
	return m.Size() == 0
}

// Keys returns all keys
func (m *InterfaceInterfaceMap) Keys() []interface{} {
	var keys []interface{}
	m.mutex.RLock()
	for k, _ := range m.data {
		keys = append(keys, k)
	}
	m.mutex.RUnlock()
	return keys
}

// Values returns all values
func (m *InterfaceInterfaceMap) Values() []interface{} {
	var values []interface{}
	m.mutex.RLock()
	for _, v := range m.data {
		values = append(values, v)
	}
	m.mutex.RUnlock()
	return values
}

// Clone returns a shallow clone of map
func (m *InterfaceInterfaceMap) Clone() map[interface{}]interface{} {
	dst := make(map[interface{}]interface{}, len(m.data))
	m.mutex.RLock()
	for k, v := range m.data {
		dst[k] = v
	}
	m.mutex.RUnlock()
	return dst
}
