package orderedmap

// Iterator is an OMIterator over a OrderedMap.
type Iterator[T any] interface {
	HasNext() bool
	Next() bool
	GetNext() (T, bool)
	GetCurrent() (T, bool)
}

// OMIterator is an Iterator implementation.
type OMIterator[T any] struct {
	index   int
	values  []T
	current T
}

// HasNext returns true if the OMIterator has a next element.
func (i *OMIterator[T]) HasNext() bool {
	if i.index < len(i.values) {
		return true
	}
	return false
}

// Next moves the OMIterator to the next element and returns true if there was a next element in the OMIterator.
func (i *OMIterator[T]) Next() bool {
	if i.HasNext() {
		i.current = i.values[i.index]
		i.index++
		return true
	}
	return false
}

// GetNext returns the next element in the OMIterator.
func (i *OMIterator[T]) GetNext() (T, bool) {
	if i.HasNext() {
		t := i.values[i.index]
		i.current = t
		i.index++
		return t, true
	}
	return *new(T), false
}

// GetCurrent returns the current element in the OMIterator.
func (i *OMIterator[T]) GetCurrent() (T, bool) {
	if i.HasNext() {
		return i.current, true
	}
	return *new(T), false
}
