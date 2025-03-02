package tree

import (
	"errors"
	"fmt"

	"github.com/ZacharyDuve/godatacollections"
	"github.com/ZacharyDuve/godatacollections/stack"
)

type AVLTree[K, T any] struct {
	zeroValue T
	// Pointer to the root node of the tree
	root *avlTreeNode[K, T]
	// Function to compare the Keys of two nodes
	kCompFunc func(K, K) int
	// Function to convert a T into a K
	tToKFunc func(T) K
}

type avlTreeNode[K, T any] struct {
	// Item that we are storing
	t T
	// How balanced the subtree is. Should be at most 1 which means that the left and right
	// differ in height by at most 1
	balance uint8
	// Pointer to the left child
	left *avlTreeNode[K, T]
	// Pointer to the right child
	right *avlTreeNode[K, T]
}

func (this *AVLTree[K, T]) Insert(t T) error {
	return errors.ErrUnsupported
}

func (this *AVLTree[K, T]) insertRec(tToBeInserted T, curNode *avlTreeNode[K, T]) error {
	if curNode == nil {
		panic("Should never have gotten to a nil child on an insert")
	}

	curComp
}

func (this *AVLTree[K, T]) Remove(key K) error {
	return errors.ErrUnsupported
}

func (this *AVLTree[K, T]) Contains(key K) bool {
	// Contains in a AVL Tree is just as a BST tree

	if this.root == nil {
		return false
	}

	curNode := this.root

	for curNode != nil {
		cmp := this.kCompFunc(key, this.tToKFunc(curNode.t))

		if cmp == 0 {
			return true
		} else if cmp < 0 {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}

	return false
}

func (this *AVLTree[K, T]) GetByKey(key K) (T, error) {
	curNode := this.root

	for curNode != nil {
		curComp := this.kCompFunc(key, this.tToKFunc(curNode.t))
		if curComp == 0 {
			return curNode.t, nil
		} else if curComp < 0 {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}

	return this.zeroValue, fmt.Errorf("unable to find T with key %v", key)
}

// Completely reusing code from the BST
func (this *AVLTree[K, T]) Iterator() godatacollections.Iterator[T] {
	iter := &avlIterator[K, T]{nodeStack: stack.NewLStack[*avlTreeNode[K, T]](nil), zeroValue: this.zeroValue}

	next := this.root

	for next != nil {
		iter.nodeStack.Push(next)
		next = next.left
	}
	// Need to ensure that the first value for next is preped
	iter.prepNext()

	return iter
}

type avlIterator[K, T any] struct {
	nodeStack *stack.LStack[*avlTreeNode[K, T]]
	next      *avlTreeNode[K, T]
	zeroValue T
}

func (this *avlIterator[K, T]) Close() error {
	return nil
}

func (this *avlIterator[K, T]) HasNext() bool {
	return this.next != nil
}

func (this *avlIterator[K, T]) prepNext() {

	// Save to ignore error as we are using a nil value for zero so we can tell when we have reached the end
	this.next, _ = this.nodeStack.Pop()
}

func (this *avlIterator[K, T]) Next() (T, error) {

	if this.next == nil {
		return this.zeroValue, errors.New("nothing left to iterate over")
	}

	if this.next.right != nil {
		nextAfter := this.next.right

		for nextAfter != nil {
			this.nodeStack.Push(nextAfter)
			nextAfter = nextAfter.left
		}
	}

	retNext := this.next

	// Need to prepare the next value
	this.prepNext()

	return retNext.t, nil
}
