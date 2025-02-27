package stack

import "github.com/ZacharyDuve/godatacollections"

// Implementation of Stack via a linked list
type LStack[T any] struct {
	// Since we are a stack we only need to have a head
	head      *lStackNode[T]
	zeroValue T
}

func NewLStack[T any](zeroValue T) *LStack[T] {
	return &LStack[T]{zeroValue: zeroValue}
}

type lStackNode[T any] struct {
	t    T
	next *lStackNode[T]
}

func (this *LStack[T]) Push(newT T) {
	newNode := &lStackNode[T]{t: newT, next: this.head}

	this.head = newNode
}

func (this *LStack[T]) Pop() (T, error) {
	if this.head == nil {
		// Nothing is in the stack so return zero value
		return this.zeroValue, godatacollections.EmptyError()
	}

	retVal := this.head.t

	this.head = this.head.next

	return retVal, nil
}
