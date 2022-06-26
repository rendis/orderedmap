package orderedmap

// Iterator is an iterator over a OrderedMap.
type Iterator[T any] interface {
	HasNext() bool
	Next() bool
	GetNext() (T, bool)
	GetCurrent() (T, bool)
	GetCurrentV() (T, int, bool)
}

func NewIterator[T any](values []T) Iterator[T] {
	return &OMIterator[T]{
		values: values,
		index:  -1,
		vlen:   len(values) - 1,
	}
}

// OMIterator is an Iterator implementation.
type OMIterator[T any] struct {
	index   int
	values  []T
	vlen    int
	current T
}

// HasNext returns true if the OMIterator has a next element.
func (i *OMIterator[T]) HasNext() bool {
	if i.index < i.vlen {
		return true
	}
	return false
}

// Next moves the OMIterator to the next element and returns true if there was a next element in the OMIterator.
func (i *OMIterator[T]) Next() bool {
	if i.HasNext() {
		i.index++
		i.current = i.values[i.index]
		return true
	}
	return false
}

// GetNext returns the next element in the OMIterator.
func (i *OMIterator[T]) GetNext() (T, bool) {
	if i.HasNext() {
		i.index++
		t := i.values[i.index]
		i.current = t
		return t, true
	}

	// Last element
	if i.index > -1 {
		return i.current, true
	}

	return *new(T), false
}

// GetCurrent returns the current element and true if there was a current element in the OMIterator.
func (i *OMIterator[T]) GetCurrent() (T, bool) {
	if i.index > -1 {
		return i.current, true
	}
	return *new(T), false
}

// GetCurrentV returns the current element, the index and true if there was a current element in the OMIterator.
func (i *OMIterator[T]) GetCurrentV() (T, int, bool) {
	if i.index > -1 {
		return i.current, i.index, true
	}
	return *new(T), -1, false
}
