package godatacollections

type Queue[T any] interface {
	// Enqueue adds value T to the end of the queue
	Enqueue(T)
	// Enqueue removes value that is at front of queue and returns it
	// returns error if there is nothing in the queue
	Dequeue() (T, error)
}
