package maps

import (
	"sync"
	"sync/atomic"
)

type (
	// InterfaceInterfaceCOWMap is a synchronous copy on write map. Reads are cheap. Writes are expensive.
	InterfaceInterfaceCOWMap struct {
		data  atomic.Value
		mutex sync.Mutex // used only by writers
	}
)

// NewInterfaceInterfaceCOWMap initializes a new empty map.
// Use of nil to empty the InterfaceInterfaceCOWMap is okay.
func NewInterfaceInterfaceCOWMap() *InterfaceInterfaceCOWMap {
	m := new(InterfaceInterfaceCOWMap)
	m.data.Store(make(map[interface{}]interface{}))
	return m
}

// Get retrieves the value associated with the key.
func (m *InterfaceInterfaceCOWMap) Get(key interface{}) interface{} {
	data := m.data.Load().(map[interface{}]interface{})
	return data[key]
}

// Load retrieves the value associated with the key.
func (m *InterfaceInterfaceCOWMap) Load(key interface{}) (value interface{}, ok bool) {
	data := m.data.Load().(map[interface{}]interface{})
	value, ok = data[key]
	return
}

// Set inserts a key-value pair.
func (m *InterfaceInterfaceCOWMap) Set(key interface{}, value interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[interface{}]interface{})
	dst := m.dup(data)
	dst[key] = value
	m.data.Store(dst)
	m.mutex.Unlock()
}

// SetIfNotPresent inserts a key-value pair if not exists.
func (m *InterfaceInterfaceCOWMap) SetIfNotPresent(key interface{}, value interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[interface{}]interface{})
	if _, ok := data[key]; !ok {
		dst := m.dup(data)
		dst[key] = value
		m.data.Store(dst)
	}
	m.mutex.Unlock()
}

// CopyFrom efficiently inserts all the key-value pairs.
func (m *InterfaceInterfaceCOWMap) CopyFrom(src map[interface{}]interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[interface{}]interface{})
	dst := m.dup(data)
	for k, v := range src {
		dst[k] = v
	}
	m.data.Store(dst)
	m.mutex.Unlock()
}

// Remove removes key from the InterfaceInterfaceCOWMap.
func (m *InterfaceInterfaceCOWMap) Remove(key interface{}) {
	m.mutex.Lock()
	data := m.data.Load().(map[interface{}]interface{})
	dst := m.dup(data)
	delete(dst, key)
	m.data.Store(dst)
	m.mutex.Unlock()
}

// Clear removes all keys from the InterfaceInterfaceCOWMap.
func (m *InterfaceInterfaceCOWMap) Clear() {
	m.mutex.Lock()
	m.data.Store(make(map[interface{}]interface{}))
	m.mutex.Unlock()
}

// Keys returns all keys
func (m *InterfaceInterfaceCOWMap) Keys() []interface{} {
	data := m.data.Load().(map[interface{}]interface{})

	var keys []interface{}
	for k, _ := range data {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values
func (m *InterfaceInterfaceCOWMap) Values() []interface{} {
	data := m.data.Load().(map[interface{}]interface{})

	var values []interface{}
	for _, v := range data {
		values = append(values, v)
	}
	return values
}

// Clone returns a shallow clone of map
func (m *InterfaceInterfaceCOWMap) Clone() map[interface{}]interface{} {
	data := m.data.Load().(map[interface{}]interface{})
	return m.dup(data)
}

func (m *InterfaceInterfaceCOWMap) dup(src map[interface{}]interface{}) map[interface{}]interface{} {
	dst := make(map[interface{}]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
