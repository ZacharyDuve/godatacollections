package godatacollections

import "io"

type Iterator[T any] interface {
	// Including io.Closer
	// Needed for cleaning up potential locks that the iterator use
	io.Closer
	// Next returns the next item from the iterator
	// Returns zero value and EOF if there is no next item
	Next() (T, error)

	// HasNext returns if there is next item to pull
	HasNext() bool
}
