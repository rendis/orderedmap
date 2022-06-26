package orderedmap

// New creates a new OrderedMap
func New[T any]() OrderedMap[T] {
	return &OMap[T]{
		data: make(map[any]T),
	}
}

// OrderedMap is a map with ordered keys.
type OrderedMap[T any] interface {
	Len() int
	Set(key any, val T)
	Get(key any) (T, bool)
	Delete(key any) (T, bool)
	Keys() []any
	Values() []T
	Exists(key any) bool
	IndexOf(key any) int
	ReplaceKey(oldKey any, newKey any) bool
	SetBefore(presentKey any, newKey any, val T) (int, bool)
	SetAfter(presentKey any, newKey any, val T) (int, bool)
	Iterator() Iterator[T]
}

// OMap is the OrderedMap implementation.
type OMap[T any] struct {
	key  []any
	data map[any]T
}

// Len returns the number of elements in the map.
func (m *OMap[T]) Len() int {
	return len(m.key)
}

// Set sets the value for the given key.
func (m *OMap[T]) Set(key any, val T) {
	if _, ok := m.data[key]; !ok {
		m.key = append(m.key, key)
	}
	m.data[key] = val
}

// Get returns the value for the given key.
// If the key does not exist, the second argument will be false.
func (m *OMap[T]) Get(key any) (T, bool) {
	if v, ok := m.data[key]; ok {
		return v, true
	}
	return *new(T), false
}

// Delete deletes the value for the given key.
// Returns deleted value, true if the key is found.
func (m *OMap[T]) Delete(key any) (T, bool) {
	if v, ok := m.data[key]; ok {
		delete(m.data, key)
		return v, true
	}
	return *new(T), false
}

// Keys returns the keys of the map.
func (m *OMap[T]) Keys() []any {
	return m.key
}

// Values returns the values of the map.
func (m *OMap[T]) Values() []T {
	vals := make([]T, len(m.key))
	for i, k := range m.key {
		vals[i] = m.data[k]
	}
	return vals
}

// Exists returns true if the key exists in the map.
func (m *OMap[T]) Exists(key any) bool {
	_, ok := m.data[key]
	return ok
}

// IndexOf returns the index of the given key.
// If the key does not exist, returns -1.
func (m *OMap[T]) IndexOf(key any) int {
	for i, k := range m.key {
		if k == key {
			return i
		}
	}
	return -1
}

// ReplaceKey replaces the key of the given key with the new key.
// Returns false if the key does not exist.
func (m *OMap[T]) ReplaceKey(oldKey any, newKey any) bool {
	if _, ok := m.data[oldKey]; !ok {
		return false
	}
	if _, ok := m.data[newKey]; ok {
		return false
	}
	i := m.IndexOf(oldKey)
	m.key[i] = newKey
	m.data[newKey] = m.data[oldKey]
	delete(m.data, oldKey)
	return true
}

// SetBefore sets the new (newKey, value) pair before the given presentKey.
// Return false if presentKey does not exist or if newKey already exists.
func (m *OMap[T]) SetBefore(presentKey any, newKey any, value T) (int, bool) {
	return m.setAt(presentKey, newKey, value, 0)
}

// SetAfter sets the new (newKey, value) pair after the given presentKey.
// Return false if presentKey does not exist or if newKey already exists.
func (m *OMap[T]) SetAfter(presentKey any, newKey any, value T) (int, bool) {
	return m.setAt(presentKey, newKey, value, 1)
}

// Iterator returns an OMIterator for the map.
func (m *OMap[T]) Iterator() Iterator[T] {
	values := make([]T, 0, len(m.key))
	for _, key := range m.key {
		values = append(values, m.data[key])
	}
	return NewIterator[T](values)
}

func (m *OMap[T]) setAt(presentKey any, newKey any, val T, delta int) (int, bool) {
	if _, ok := m.data[presentKey]; !ok {
		return 0, false
	}
	if _, ok := m.data[newKey]; ok {
		return 0, false
	}
	m.data[newKey] = val
	pos := -1
	for i, k := range m.key {
		if k == presentKey {
			pos = i + delta
			break
		}
	}
	arr := make([]any, 0, len(m.key)+1)
	arr = append(arr, m.key[:pos]...)
	arr = append(arr, newKey)
	arr = append(arr, m.key[pos:]...)
	m.key = arr
	return pos, true
}
