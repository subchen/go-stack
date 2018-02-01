package maps

import (
	"sync"
)

type (
	// StringInterfaceMap is a synchronous map.
	StringInterfaceMap struct {
		data  map[string]interface{}
		mutex sync.RWMutex // used only by writers
	}
)

// NewStringInterfaceMap initializes a new empty map.
// Use of nil to empty the StringInterfaceMap is okay.
func NewStringInterfaceMap() *StringInterfaceMap {
	m := new(StringInterfaceMap)
	m.data = make(map[string]interface{})
	return m
}

// Get retrieves the value associated with the key.
func (m *StringInterfaceMap) Get(key string) interface{} {
	m.mutex.RLock()
	value := m.data[key]
	m.mutex.RUnlock()
	return value
}

// Load retrieves the value associated with the key.
func (m *StringInterfaceMap) Load(key string) (value interface{}, ok bool) {
	m.mutex.RLock()
	value, ok = m.data[key]
	m.mutex.RUnlock()
	return
}

// Set inserts a key-value pair.
func (m *StringInterfaceMap) Set(key string, value interface{}) {
	m.mutex.Lock()
	m.data[key] = value
	m.mutex.Unlock()
}

// SetIfNotPresent inserts a key-value pair if not exists.
func (m *StringInterfaceMap) SetIfNotPresent(key string, value interface{}) {
	m.mutex.Lock()
	if _, ok := m.data[key]; !ok {
		m.data[key] = value
	}
	m.mutex.Unlock()
}

// CopyFrom efficiently inserts all the key-value pairs.
func (m *StringInterfaceMap) CopyFrom(src map[string]interface{}) {
	m.mutex.Lock()
	for k, v := range src {
		m.data[k] = v
	}
	m.mutex.Unlock()
}

// Remove removes key from the StringInterfaceMap.
func (m *StringInterfaceMap) Remove(key string) {
	m.mutex.Lock()
	delete(m.data, key)
	m.mutex.Unlock()
}

// Clear removes all keys from the StringInterfaceMap.
func (m *StringInterfaceMap) Clear() {
	m.mutex.Lock()
	m.data = make(map[string]interface{})
	m.mutex.Unlock()
}

// Keys returns all keys
func (m *StringInterfaceMap) Keys() []string {
	var keys []string
	m.mutex.RLock()
	for k, _ := range m.data {
		keys = append(keys, k)
	}
	m.mutex.RUnlock()
	return keys
}

// Values returns all values
func (m *StringInterfaceMap) Values() []interface{} {
	var values []interface{}
	m.mutex.RLock()
	for _, v := range m.data {
		values = append(values, v)
	}
	m.mutex.RUnlock()
	return values
}

// Clone returns a shallow clone of map
func (m *StringInterfaceMap) Clone() map[string]interface{} {
	dst := make(map[string]interface{}, len(m.data))
	m.mutex.RLock()
	for k, v := range m.data {
		dst[k] = v
	}
	m.mutex.RUnlock()
	return dst
}
