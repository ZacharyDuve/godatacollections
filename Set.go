package godatacollections

// Set is a datastructure that doesn't allow for duplicates
// I am trying to allow for one to specify what Keys an item differently from the item itself
// K is the type of the Key that each item has.
//
//	One should be able to turn a T into a K and all K in the set should be unique
//
// T is the type that is being stored
type Set[K, T any] interface {
	// Insert allows for inserting a unique item into the set
	// An error will be returned for duplicate items
	Insert(T) error

	// Contains returns if the Set contains an item T that has an equivalent Key
	Contains(K) bool

	// GetByKey allows for retreival of the item by its key
	// Returns zero value if there is no item with matching t
	GetByKey(K) T

	// Remove will remove item T from the Set that resolves to Key passed in.
	// An error will be returned if there was nothing deleted
	Remove(K) error

	Iterator() Iterator[T]
}
