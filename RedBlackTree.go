package godatacollections

import (
	"errors"
)

type RedBlackTree[K, T any] struct {
	keyCompFunc func(K, K) int
	tToKFunc    func(T) K
	root        *rbNode[T]
}

type rbColor uint8

const (
	black rbColor = 0
	red   rbColor = 1
)

type rbNode[T any] struct {
	left, right *rbNode[T]
	t           T
	color       rbColor
}

func (this *RedBlackTree[K, T]) Contains(k K) bool {
	curNode := this.root

	for curNode != nil {
		comp := this.keyCompFunc(k, this.tToKFunc(curNode.t))

		if comp == 0 {
			return true
		} else if comp < 0 {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}

	//Only way that we got here is if we ended up with a nil curNode which means we didn't find what we wanted
	return false
}

func (this *RedBlackTree[K, T]) Insert(newT T) error {
	// Inserting as a set we don't want duplicates
	if this.root == nil {
		// If we do not have a root node yet
		this.root = &rbNode[T]{color: black, t: newT}
		// For root just insert as black and return
		return nil
	}

	// Otherwise if we have root lets start doing normal BST insertion
	newKey := this.tToKFunc(newT)
	// New Node is supposed to be Red
	newNode := &rbNode[T]{color: red, t: newT}
	var curGrandParentNode *rbNode[T]
	curParentNode := this.root

	// Handle insert as BST at first
	for {
		curComp := this.keyCompFunc(newKey, this.tToKFunc(curParentNode.t))

		if curComp == 0 {
			// We have a duplicate
			// Return error of duplicate, nothing else to do
			return errors.New("Unable to insert due to duplicate record")
		} else if curComp < 0 {
			// We need to go left since we are less
			if curParentNode.left == nil {
				curParentNode.left = newNode
				// Once we have set the link than we stop BST insert
				break
			} else {
				curGrandParentNode = curParentNode
				curParentNode = curParentNode.left
			}
		} else {
			// We need to go right since we are greater

			if curParentNode.right == nil {
				curParentNode.right = newNode
				// Once we have set the link than we stop BST insert
				break
			} else {
				curGrandParentNode = curParentNode
				curParentNode = curParentNode.right
			}
		}
	}
}
