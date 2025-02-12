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
	// Using parent as that saves us having to store entire ancesty during insert/deletes
	// Also allows for easy checks for which path we came down
	parent *rbNode[T]
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
	// We need to know the parent trust us on this
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
				// Found our spot to put the new node
				// Set the parents left child to our node
				curParentNode.left = newNode
				// Set the new node's parent
				newNode.parent = curParentNode
				// Once we have set the link than we stop BST insert
				break
			} else {
				curParentNode = curParentNode.left
			}
		} else {
			// We need to go right since we are greater

			if curParentNode.right == nil {
				// Found our spot to put the new node
				// Set the parents right child to our node
				curParentNode.right = newNode
				// Set the new node's parent
				newNode.parent = curParentNode
				// Once we have set the link than we stop BST insert
				break
			} else {
				curParentNode = curParentNode.right
			}
		}
	}

	// At this point the new node should have been inserted if not duplicate has been found
	// Need to recolor first

	// Need to handle rotations if recoloring didn't work out well for us
}
