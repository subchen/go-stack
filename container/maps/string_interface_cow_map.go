package maps

import (
	"sync"
	"sync/atomic"
)

type (
	// StringInterfaceCOWMap is a synchronous copy on write map. Reads are cheap. Writes are expensive.
	StringInterfaceCOWMap struct {
		data  atomic.Value
		mutex sync.Mutex // used only by writers
	}
)

// NewStringInterfaceCOWMap initializes a new empty map.
// Use of nil to empty the StringInterfaceCOWMap is okay.
func NewStringInterfaceCOWMap() *StringInterfaceCOWMap {
	m := new(StringInterfaceCOWMap)
	m.data.Store(make(map[string]interface{}))
	return m
}

// Get retrieves the value associated with the key.
func (m *StringInterfaceCOWMap) Get(key string) interface{} {
	data := m.data.Load().(map[string]interface{})
	return data[key]
}

// Load retrieves the value associated with the key.
func (m *StringInterfaceCOWMap) Load(key string) (value interface{}, ok bool) {
	data := m.data.Load().(map[string]interface{})
	value, ok = data[key]
	return
}

// Set inserts a key-value pair.
func (m *StringInterfaceCOWMap) Set(key string, value interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[string]interface{})
	dst := m.dup(data)
	dst[key] = value
	m.data.Store(dst)
	m.mutex.Unlock()
}

// SetIfNotPresent inserts a key-value pair if not exists.
func (m *StringInterfaceCOWMap) SetIfNotPresent(key string, value interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[string]interface{})
	if _, ok := data[key]; !ok {
		dst := m.dup(data)
		dst[key] = value
		m.data.Store(dst)
	}
	m.mutex.Unlock()
}

// CopyFrom efficiently inserts all the key-value pairs.
func (m *StringInterfaceCOWMap) CopyFrom(src map[string]interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[string]interface{})
	dst := m.dup(data)
	for k, v := range src {
		dst[k] = v
	}
	m.data.Store(dst)
	m.mutex.Unlock()
}

// Remove removes key from the StringInterfaceCOWMap.
func (m *StringInterfaceCOWMap) Remove(key string) {
	m.mutex.Lock()
	data := m.data.Load().(map[string]interface{})
	dst := m.dup(data)
	delete(dst, key)
	m.data.Store(dst)
	m.mutex.Unlock()
}

// Clear removes all keys from the StringInterfaceCOWMap.
func (m *StringInterfaceCOWMap) Clear() {
	m.mutex.Lock()
	m.data.Store(make(map[string]interface{}))
	m.mutex.Unlock()
}

// Keys returns all keys
func (m *StringInterfaceCOWMap) Keys() []string {
	data := m.data.Load().(map[string]interface{})

	var keys []string
	for k, _ := range data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values
func (m *StringInterfaceCOWMap) Values() []interface{} {
	data := m.data.Load().(map[string]interface{})

	var values []interface{}
	for _, v := range data {
		values = append(values, v)
	}
	return values
}

// Clone returns a shallow clone of map
func (m *StringInterfaceCOWMap) Clone() map[string]interface{} {
	data := m.data.Load().(map[string]interface{})
	return m.dup(data)
}

func (m *StringInterfaceCOWMap) dup(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
