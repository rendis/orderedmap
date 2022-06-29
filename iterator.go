package orderedmap

// Iterator is an iterator over a OrderedMap.
type Iterator[T any] interface {
	// HasNext returns true if the omIter has a next element.
	HasNext() bool

	// Next moves the omIter to the next element and returns true if there was a next element in the omIter.
	Next() bool

	// GetNext returns the next element in the omIter.
	GetNext() (T, bool)

	// GetCurrent returns the current element and true if there was a current element in the omIter.
	GetCurrent() (T, bool)

	// GetCurrentV returns the current element, the index and true if there was a current element in the omIter.
	GetCurrentV() (T, int, bool)
}

// NewIterator returns a new Iterator.
func NewIterator[T any](values []T) Iterator[T] {
	return &omIter[T]{
		values: values,
		index:  -1,
		vlen:   len(values) - 1,
	}
}

// omIter is an Iterator implementation.
type omIter[T any] struct {
	index   int
	values  []T
	vlen    int
	current T
}

func (i *omIter[T]) HasNext() bool {
	if i.index < i.vlen {
		return true
	}
	return false
}

func (i *omIter[T]) Next() bool {
	if i.HasNext() {
		i.index++
		i.current = i.values[i.index]
		return true
	}
	return false
}

func (i *omIter[T]) GetNext() (T, bool) {
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

func (i *omIter[T]) GetCurrent() (T, bool) {
	if i.index > -1 {
		return i.current, true
	}
	return *new(T), false
}

func (i *omIter[T]) GetCurrentV() (T, int, bool) {
	if i.index > -1 {
		return i.current, i.index, true
	}
	return *new(T), -1, false
}
