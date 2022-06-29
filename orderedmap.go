package orderedmap

// New return a new OrderedMap
func New[T any]() OrderedMap[T] {
	return &om[T]{
		data: make(map[any]T),
	}
}

// OrderedMap is a map with ordered keys.
type OrderedMap[T any] interface {
	// Len returns the number of elements in the map.
	Len() int

	// Set sets the value for the given key.
	Set(key any, val T)

	// Get returns the value for the given key.
	// If the key does not exist, the second argument will be false.
	Get(key any) (T, bool)

	// Delete deletes the value for the given key.
	// Returns deleted value, true if the key is found.
	Delete(key any) (T, bool)

	// Keys returns the keys of the map.
	Keys() []any

	// Values returns the values of the map.
	Values() []T

	// Exists returns true if the key exists in the map.
	Exists(key any) bool

	// IndexOf returns the index of the given key.
	// If the key does not exist, returns -1.
	IndexOf(key any) int

	// ReplaceKey replaces the key of the given key with the new key.
	// Returns false if the key does not exist.
	ReplaceKey(oldKey any, newKey any) bool

	// SetBefore sets the new (newKey, value) pair before the given presentKey.
	// Return false if presentKey does not exist or if newKey already exists.
	SetBefore(presentKey any, newKey any, val T) (int, bool)

	// SetAfter sets the new (newKey, value) pair after the given presentKey.
	// Return false if presentKey does not exist or if newKey already exists.
	SetAfter(presentKey any, newKey any, val T) (int, bool)

	// Iterator returns an omIter for the map.
	Iterator() Iterator[T]
}

// om is the OrderedMap implementation.
type om[T any] struct {
	key  []any
	data map[any]T
}

func (m *om[T]) Len() int {
	return len(m.key)
}

func (m *om[T]) Set(key any, val T) {
	if _, ok := m.data[key]; !ok {
		m.key = append(m.key, key)
	}
	m.data[key] = val
}

func (m *om[T]) Get(key any) (T, bool) {
	if v, ok := m.data[key]; ok {
		return v, true
	}
	return *new(T), false
}

func (m *om[T]) Delete(key any) (T, bool) {
	if v, ok := m.data[key]; ok {
		delete(m.data, key)
		return v, true
	}
	return *new(T), false
}

func (m *om[T]) Keys() []any {
	return m.key
}

func (m *om[T]) Values() []T {
	vals := make([]T, len(m.key))
	for i, k := range m.key {
		vals[i] = m.data[k]
	}
	return vals
}

func (m *om[T]) Exists(key any) bool {
	_, ok := m.data[key]
	return ok
}

func (m *om[T]) IndexOf(key any) int {
	for i, k := range m.key {
		if k == key {
			return i
		}
	}
	return -1
}

func (m *om[T]) ReplaceKey(oldKey any, newKey any) bool {
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

func (m *om[T]) SetBefore(presentKey any, newKey any, value T) (int, bool) {
	return m.setAt(presentKey, newKey, value, 0)
}

func (m *om[T]) SetAfter(presentKey any, newKey any, value T) (int, bool) {
	return m.setAt(presentKey, newKey, value, 1)
}

func (m *om[T]) Iterator() Iterator[T] {
	values := make([]T, 0, len(m.key))
	for _, key := range m.key {
		values = append(values, m.data[key])
	}
	return NewIterator[T](values)
}

func (m *om[T]) setAt(presentKey any, newKey any, val T, delta int) (int, bool) {
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
