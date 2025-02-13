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
	newNode := &rbNode[T]{t: newT}
	// First insert like a normal BST
	err := this.insertRBTreeLikeBST(newNode)
	if err != nil {
		return err
	}
	// At this point the new node should have been inserted if not duplicate has been found
	// Check if the parent of the node is black or not
	if newNode.parent != nil && newNode.parent.color == black {
		// If the parent is black then we are all set
		return nil
	}
	// We have a red parent so we need to do something
	// Need to recolor first

	// Need to handle rotations if recoloring didn't work out well for us
}

func (this *RedBlackTree[K, T]) insertRBTreeLikeBST(newNode *rbNode[T]) error {
	// Inserting as a set we don't want duplicates
	if this.root == nil {
		// If we do not have a root node yet
		this.root = newNode
		newNode.color = black
		// For root just insert as black and return
		return nil
	}

	// Otherwise if we have root lets start doing normal BST insertion
	newKey := this.tToKFunc(newNode.t)
	// New Node is supposed to be Red
	//newNode := &rbNode[T]{color: red, t: newT}
	// We need to know the parent trust us on this
	curNode := this.root

	// Handle insert as BST at first
	for {
		curComp := this.keyCompFunc(newKey, this.tToKFunc(curNode.t))

		if curComp == 0 {
			// We have a duplicate
			// Return error of duplicate, nothing else to do
			return errors.New("Unable to insert due to duplicate record")
		} else if curComp < 0 {
			// We need to go left since we are less
			if curNode.left == nil {
				// Found our spot to put the new node
				// Set the parents left child to our node
				curNode.left = newNode
				// Set the new node's parent
				newNode.parent = curNode
				// Once we have set the link than we stop BST insert
				break
			} else {
				curNode = curNode.left
			}
		} else {
			// We need to go right since we are greater

			if curNode.right == nil {
				// Found our spot to put the new node
				// Set the parents right child to our node
				curNode.right = newNode
				// Set the new node's parent
				newNode.parent = curNode
				// Once we have set the link than we stop BST insert
				break
			} else {
				curNode = curNode.right
			}
		}
	}

	// If we made it this far then we did an insert correctly
	return nil
}

func resolveRBTreeInsertViolations[T any](curNode *rbNode[T]) {

	for curNode.parent != nil && curNode.parent.parent != nil {
		// While we can find an uncle lets keep going

		gParent := curNode.parent.parent
		var uncle *rbNode[T]
		if gParent.left == curNode.parent {
			// Parent is the left node so uncle is on the right
			uncle = gParent.right
		} else {
			// Otherwise the uncle is on the left
			uncle = gParent.left
		}

		if uncle == nil || uncle.color == black {
			// We treat nil refs as black

		}
	}
}
