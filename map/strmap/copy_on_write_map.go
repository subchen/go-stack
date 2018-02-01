package strmap

import (
	"sync"
	"sync/atomic"
)

type (
	// CopyOnWriteMap is a synchronous copy on write map. Reads are cheap. Writes are expensive.
	CopyOnWriteMap struct {
		data  atomic.Value
		mutex sync.Mutex // used only by writers
	}
)

// NewCopyOnWriteMap initializes a new empty map.
// Use of nil to empty the CopyOnWriteMap is okay.
func NewCopyOnWriteMap() *CopyOnWriteMap {
	m := new(CopyOnWriteMap)
	m.data.Store(empty())
	return m
}

// Get retrieves the value associated with the key.
func (m *CopyOnWriteMap) Get(key string) interface{} {
	data := m.data.Load().(map[string]interface{})
	return data[key]
}

// GetOK retrieves the value associated with the key.
func (m *CopyOnWriteMap) GetOK(key string) (value interface{}, ok bool) {
	data := m.data.Load().(map[string]interface{})
	value, ok = data[key]
	return
}

// Set inserts a key-value pair.
func (m *CopyOnWriteMap) Set(key string, value interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data := m.data.Load().(map[string]interface{})
	dst := dup(data)
	dst[key] = value
	m.data.Store(dst)
}

// Copy efficiently inserts all the key-value pairs.
func (m *CopyOnWriteMap) Copy(src map[string]interface{}) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data := m.data.Load().(map[string]interface{})
	dst := dup(data)
	cp(dst, src)
	m.data.Store(dst)
}

// Remove removes key from the CopyOnWriteMap.
func (m *CopyOnWriteMap) Remove(key string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data := m.data.Load().(map[string]interface{})
	dst := dup(data)
	delete(dst, key)
	m.data.Store(dst)
}

// Clear removes all keys from the CopyOnWriteMap.
func (m *CopyOnWriteMap) Clear() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.data.Store(empty())
}

func (m *CopyOnWriteMap) Keys() []string {
	data := m.data.Load().(map[string]interface{})

	var keys []string
	for k, _ := range data {
		keys = append(keys, k)
	}
	return keys
}

func (m *CopyOnWriteMap) Values() []interface{} {
	data := m.data.Load().(map[string]interface{})

	var values []interface{}
	for _, v := range data {
		values = append(values, v)
	}
	return values
}

func (m *CopyOnWriteMap) RawMap() map[string]interface{} {
	data := m.data.Load().(map[string]interface{})
	return dup(data)
}
