package stack

type DLStack[T any] struct {
	head      *dlStackNode[T]
	tail      *dlStackNode[T]
	zeroValue T
}

func NewDLStack[T any](zeroValue T) *DLStack[T] {
	return &DLStack[T]{zeroValue: zeroValue}
}

type dlStackNode[T any] struct {
	t    T
	prev *dlStackNode[T]
	next *dlStackNode[T]
}

func (this *DLStack[T]) Push(newT T) {
	newNode := &dlStackNode[T]{t: newT}
	if this.head == nil {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.prev = this.tail
		this.tail.next = newNode
		this.tail = newNode
	}
}

func (this *DLStack[T]) Pop() T {
	if this.tail == nil {
		// Nothing is in the stack so return zero value
		return this.zeroValue
	}

	oldTail := this.tail

	if this.head == this.tail {
		// We have only one value left so clear out head and tail
		this.head = nil
		this.tail = nil
	} else {
		// We are deleting an interior node so just remove it and move tail
		oldTail.prev.next = nil
		this.tail = oldTail.prev
	}

	return oldTail.t
}
