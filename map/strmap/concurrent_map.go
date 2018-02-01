package strmap

import (
	"sync"
)

type (
	// ConcurrentMap is a synchronous map.
	ConcurrentMap struct {
		data  map[string]interface{}
		mutex sync.RWMutex // used only by writers
	}
)

// NewConcurrentMap initializes a new empty map.
// Use of nil to empty the ConcurrentMap is okay.
func NewConcurrentMap() *ConcurrentMap {
	m := new(ConcurrentMap)
	m.data = empty()
	return m
}

// Get retrieves the value associated with the key.
func (m *ConcurrentMap) Get(key string) interface{} {
	m.mutex.RLock()
	value := m.data[key]
	m.mutex.RUnlock()
	return value
}

// GetOK retrieves the value associated with the key.
func (m *ConcurrentMap) GetOK(key string) (value interface{}, ok bool) {
	m.mutex.RLock()
	value, ok = m.data[key]
	m.mutex.RUnlock()
	return
}

// Set inserts a key-value pair.
func (m *ConcurrentMap) Set(key string, value interface{}) {
	m.mutex.Lock()
	m.data[key] = value
	m.mutex.Unlock()
}

// Copy efficiently inserts all the key-value pairs.
func (m *ConcurrentMap) Copy(src map[string]interface{}) {
	m.mutex.Lock()
	cp(m.data, src)
	m.mutex.Unlock()
}

// Remove removes key from the ConcurrentMap.
func (m *ConcurrentMap) Remove(key string) {
	m.mutex.Lock()
	delete(m.data, key)
	m.mutex.Unlock()
}

// Clear removes all keys from the ConcurrentMap.
func (m *ConcurrentMap) Clear() {
	m.mutex.Lock()
	m.data = empty()
	m.mutex.Unlock()
}

func (m *ConcurrentMap) Keys() []string {
	m.mutex.RLock()
	keys := make([]string, len(m.data))
	for k, _ := range m.data {
		keys = append(keys, k)
	}
	m.mutex.RUnlock()
	return keys
}

func (m *ConcurrentMap) Values() []interface{} {
	m.mutex.RLock()
	values := make([]interface{}, len(m.data))
	for _, v := range m.data {
		values = append(values, v)
	}
	m.mutex.RUnlock()
	return values
}

func (m *ConcurrentMap) RawMap() map[string]interface{} {
	m.mutex.RLock()
	dst := dup(m.data)
	m.mutex.RUnlock()
	return dst
}
