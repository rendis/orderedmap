package orderedmap

// IIterator is an iterator over a IOrderedMap.
type IIterator[T any] interface {
	HasNext() bool
	Next() bool
	GetNext() (T, bool)
	GetCurrent() T
	GetCurrentV() (T, int, bool)
}

// Iterator is an IIterator implementation.
type Iterator[T any] struct {
	index   int
	values  []T
	current T
}

// HasNext returns true if the Iterator has a next element.
func (i *Iterator[T]) HasNext() bool {
	if i.index < len(i.values) {
		return true
	}
	return false
}

// Next moves the Iterator to the next element and returns true if there was a next element in the Iterator.
func (i *Iterator[T]) Next() bool {
	if i.HasNext() {
		i.current = i.values[i.index]
		i.index++
		return true
	}
	return false
}

// GetNext returns the next element in the Iterator.
func (i *Iterator[T]) GetNext() (T, bool) {
	if i.HasNext() {
		t := i.values[i.index]
		i.current = t
		i.index++
		return t, true
	}
	return *new(T), false
}

// GetCurrent returns the current element in the Iterator.
// Prefer use with HasNext() or with Next() to avoid false positives.
// Alternatively, use GetCurrentV(), which is more verbose.
func (i *Iterator[T]) GetCurrent() T {
	if i.HasNext() {
		return i.current
	}
	return *new(T)
}

// GetCurrentV returns the current element, the index and true if there was a current element in the Iterator.
func (i *Iterator[T]) GetCurrentV() (T, int, bool) {
	if i.HasNext() {
		return i.current, i.index, true
	}
	return *new(T), -1, false
}
