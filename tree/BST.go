package tree

import (
	"errors"
	"fmt"
)

type BST[K, T any] struct {
	kCompFunc func(K, K) int
	tToKFunc  func(T) K
	zeroValue T
	root      *bstNode[K, T]
}

// NewBST creates a new Binary Search Tree
// K is the key of the items T
//
//	Each T should be able to be turned into a K
//
// T is the items actually being stored
// Example would be T is an "Employee" and K is "EmployeeIDNumber"
// kCompFunc is a function that can compare two K values.
//
//	If first K is less than second K then returned value < 0
//	If first and second K are equal then return 0
//	If first K is greater than second K than return > 0
//
// tToKFunc returns a value of type K from type T
// tZeroValue is a zero or nil state of T
//
//	This is needed due to tree allowing for non-pointer types and those need a zero value
func NewBST[K, T any](kCompFunc func(K, K) int, tToKFunc func(T) K, tZeroValue T) (*BST[K, T], error) {
	if kCompFunc == nil {
		return nil, errors.New("unable to create BST without a function to compare Keys")
	}

	if tToKFunc == nil {
		return nil, errors.New("unable to create BST without a function to convert T to a Key")
	}

	return &BST[K, T]{kCompFunc: kCompFunc, tToKFunc: tToKFunc, zeroValue: tZeroValue}, nil
}

type bstNode[K, T any] struct {
	key   K
	t     T
	left  *bstNode[K, T]
	right *bstNode[K, T]
}

func (this *BST[K, T]) Insert(newT T) error {
	newKey := this.tToKFunc(newT)

	if this.root == nil {
		// If we have no root then it is super easy as we just insert
		this.root = &bstNode[K, T]{key: newKey, t: newT}
		return nil
	}

	// So not as simple as we are going to need to compare
	// Lets calcualte the the key once as we might have to search more than once

	curNode := this.root

	// Going to delay creation of the node until we need it in case we have duplicate

	for curNode != nil {
		curComp := this.kCompFunc(newKey, curNode.key)
		if curComp == 0 {
			return fmt.Errorf("unable to insert duplicate T for key %v", newKey)
		} else if curComp < 0 {
			if curNode.left == nil {
				curNode.left = &bstNode[K, T]{key: newKey, t: newT}
			} else {
				curNode = curNode.left
			}
		} else {
			if curNode.right == nil {
				curNode.right = &bstNode[K, T]{key: newKey, t: newT}
			} else {
				curNode = curNode.right
			}
		}
	}
	return nil
}

func (this *BST[K, T]) Contains(key K) bool {
	curNode := this.root

	for curNode != nil {
		curComp := this.kCompFunc(key, curNode.key)
		if curComp == 0 {
			return true
		} else if curComp < 0 {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}

	return false
}

func (this *BST[K, T]) GetByKey(key K) T {
	curNode := this.root

	for curNode != nil {
		curComp := this.kCompFunc(key, curNode.key)
		if curComp == 0 {
			return curNode.t
		} else if curComp < 0 {
			curNode = curNode.left
		} else {
			curNode = curNode.right
		}
	}

	return this.zeroValue
}

func (this *BST[K, T]) Remove(key K) error {
	// Looking at preforming delete of the root first as that is an edge case
	// First find the node to be deleted
	var curNodeParent *bstNode[K, T] = nil
	curNode := this.root
	for curNode != nil {
		curComp := this.kCompFunc(key, curNode.key)

		if curComp < 0 {
			// Need to go left
			curNode = curNode.left
		} else if curComp > 0 {
			// Need to go right
			curNode = curNode.right
		} else {
			// curComp == 0  so we have a match
			this.deleteNode(curNode, curNodeParent)
			// Need to make sure that we return to break the loop
			return nil
		}
		curNodeParent = curNode
	}

	return fmt.Errorf("unable to delete node with key %v due to it not existing in tree", key)
}

func (this *BST[K, T]) deleteNode(nodeToBeDeleted, nodeToBeDeletedParent *bstNode[K, T]) *bstNode[K, T] {
	if nodeToBeDeleted.left == nil && nodeToBeDeleted.right == nil {
		// Our node to be deleted has no children there for we just make it disapear
		if nodeToBeDeleted == this.root {
			// Special case of if this is our root node then we just clear root
			this.root = nil
		} else {
			// Otherwise just delete the child from the parent
			if nodeToBeDeletedParent.left == nodeToBeDeleted {
				// Node is on parent's left
				nodeToBeDeletedParent.left = nil
			} else {
				// Node is on parent's right
				nodeToBeDeletedParent.right = nil
			}
		}
	} else if nodeToBeDeleted.left != nil && nodeToBeDeleted.right == nil {
		// There is a left hand child so we just replace deleted node with it
		if nodeToBeDeleted == this.root {
			// Special case for root is we set root to the old root's left child
			this.root = nodeToBeDeleted.left
		} else {
			// Otherwise set the direction for the parent that this was on to this left child
			if nodeToBeDeletedParent.left == nodeToBeDeleted {
				nodeToBeDeletedParent.left = nodeToBeDeleted.left
			} else {
				nodeToBeDeletedParent.right = nodeToBeDeleted.left
			}
		}
	} else if nodeToBeDeleted.left == nil && nodeToBeDeleted.right != nil {
		// There is a right hand child so we just replace deleted node with it
		if nodeToBeDeleted == this.root {
			// Special case for root is we set root to the old root's right child
			this.root = nodeToBeDeleted.right
		} else {
			// Otherwise set the direction for the parent that this was on to this right child
			if nodeToBeDeletedParent.left == nodeToBeDeleted {
				nodeToBeDeletedParent.left = nodeToBeDeleted.right
			} else {
				nodeToBeDeletedParent.right = nodeToBeDeleted.right
			}
		}
	} else {
		// The node that we are trying to delete has two children
		successor, successorParent := this.findSuccessor(nodeToBeDeleted, nodeToBeDeletedParent)

	}

}

func (this *BST[K, T]) findSuccessor(nodeToBeDeleted, nodeToBeDeletedParent *bstNode[K, T]) (successor, successorParent *bstNode[K, T]) {
	successorParent = nodeToBeDeletedParent
	successor = nodeToBeDeleted.right

	// Because we chose the right side we need to keep going left after the first right jump
	// Continue until we can't go left anymore. That last node is our
	for successor != nil && successor.left != nil {
		successorParent = successor
		successor = successor.left
	}

	return successor, successorParent
}
