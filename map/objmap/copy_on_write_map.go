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
func (m *CopyOnWriteMap) Get(key Key) Value {
	data := m.data.Load().(map[Key]Value)
	return data[key]
}

// GetOK retrieves the value associated with the key.
func (m *CopyOnWriteMap) GetOK(key Key) (value Value, ok bool) {
	data := m.data.Load().(map[Key]Value)
	value, ok = data[key]
	return
}

// Set inserts a key-value pair.
func (m *CopyOnWriteMap) Set(key Key, value Value) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data := m.data.Load().(map[Key]Value)
	dst := dup(data)
	dst[key] = value
	m.data.Store(dst)
}

// Copy efficiently inserts all the key-value pairs.
func (m *CopyOnWriteMap) Copy(src map[Key]Value) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data := m.data.Load().(map[Key]Value)
	dst := dup(data)
	cp(dst, src)
	m.data.Store(dst)
}

// Remove removes key from the CopyOnWriteMap.
func (m *CopyOnWriteMap) Remove(key Key) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	data := m.data.Load().(map[Key]Value)
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

func (m *CopyOnWriteMap) Keys() []Key {
	data := m.data.Load().(map[Key]Value)

	var keys []Key
	for k, _ := range data {
		keys = append(keys, k)
	}
	return keys
}

func (m *CopyOnWriteMap) Values() []Value {
	data := m.data.Load().(map[Key]Value)

	var values []Value
	for _, v := range data {
		values = append(values, v)
	}
	return values
}

func (m *CopyOnWriteMap) RawMap() map[Key]Value {
	data := m.data.Load().(map[Key]Value)
	return dup(data)
}
