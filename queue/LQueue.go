package queue

import "errors"

type LQueue[T any] struct {
	tZeroValue T
	head       *lQueueNode[T]
	tail       *lQueueNode[T]
}

type lQueueNode[T any] struct {
	t    T
	next *lQueueNode[T]
}

func NewLQueue[T any](tZeroValue T) *LQueue[T] {
	return &LQueue[T]{tZeroValue: tZeroValue}
}

func (this *LQueue[T]) Enqueue(t T) {
	newNode := &lQueueNode[T]{t: t}

	if this.head == nil {
		this.head = newNode
	} else {
		this.tail.next = newNode
	}
	this.tail = newNode
}

func (this *LQueue[T]) Dequeue() (T, error) {
	if this.head == nil {
		// Queue is empty so return empty error
		return this.tZeroValue, errors.New("empty queue")
	}

	retT := this.head.t

	this.head = this.head.next

	return retT, nil
}
