package tree

import (
	"errors"
	"fmt"
	"log"

	"github.com/ZacharyDuve/godatacollections"
	"github.com/ZacharyDuve/godatacollections/stack"
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
	t     T
	left  *bstNode[K, T]
	right *bstNode[K, T]
}

func (this *BST[K, T]) Insert(newT T) error {
	newKey := this.tToKFunc(newT)

	if this.root == nil {
		// If we have no root then it is super easy as we just insert
		this.root = &bstNode[K, T]{t: newT}
		return nil
	}

	// So not as simple as we are going to need to compare
	// Lets calcualte the the key once as we might have to search more than once

	curNode := this.root

	// Going to delay creation of the node until we need it in case we have duplicate

	for curNode != nil {
		curComp := this.kCompFunc(newKey, this.tToKFunc(curNode.t))
		if curComp == 0 {
			return fmt.Errorf("unable to insert duplicate T for key %v", newKey)
		} else if curComp < 0 {
			if curNode.left == nil {
				curNode.left = &bstNode[K, T]{t: newT}
				// We have completed insert so lets break
				break
			} else {
				curNode = curNode.left
			}
		} else {
			if curNode.right == nil {
				curNode.right = &bstNode[K, T]{t: newT}
				// We have completed insert so lets break
				break
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
		curComp := this.kCompFunc(key, this.tToKFunc(curNode.t))
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

func (this *BST[K, T]) GetByKey(key K) (T, error) {
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

func (this *BST[K, T]) Remove(key K) error {
	// Looking at preforming delete of the root first as that is an edge case
	// First find the node to be deleted
	var curNodeParent *bstNode[K, T] = nil
	curNode := this.root
	for curNode != nil {
		curComp := this.kCompFunc(key, this.tToKFunc(curNode.t))

		if curComp < 0 {
			// Need to go left
			curNodeParent = curNode
			curNode = curNode.left
		} else if curComp > 0 {
			// Need to go right
			curNodeParent = curNode
			curNode = curNode.right

		} else {
			// curComp == 0  so we have a match
			this.deleteNode(curNode, curNodeParent)
			// Need to make sure that we return to break the loop
			return nil
		}

	}

	return fmt.Errorf("unable to delete node with key %v due to it not existing in tree", key)
}

func (this *BST[K, T]) deleteNode(nodeToBeDeleted, nodeToBeDeletedParent *bstNode[K, T]) {
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
				log.Printf("Deleting Parent's (key: %v) right", this.tToKFunc(nodeToBeDeleted.t))
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
		successor, successorParent := this.findSuccessor(nodeToBeDeleted)
		this.succeedDeletedNode2Children(nodeToBeDeleted, nodeToBeDeletedParent, successor, successorParent)

	}
}

// Only can get here if our node to be deleted has only two children
func (this *BST[K, T]) succeedDeletedNode2Children(nodeToDelete, nodeToBeDeletedParent, successor, successorParent *bstNode[K, T]) {
	// Things that we know at this point about the successor node
	// - It has only a right child if any children exist
	// - It is either a direct decendant of the deleted node at which point it right side decendant
	// - OR it if it isn't a direct decendant then it is the parent's left node

	// First need to handle the special case of if it is root
	if this.root == nodeToDelete {
		this.root = successor
	} else {
		if nodeToBeDeletedParent.left == nodeToDelete {
			nodeToBeDeletedParent.left = successor
		} else {
			nodeToBeDeletedParent.right = successor
		}
	}
	// Copy away the successors decendants so we don't loose them
	// Due to knowing that the successor only has a right child from succession rules we don't need to copy the left as that will always be nil
	oldSuccessorRight := successor.right
	// Now move the old deleted node's decendants over to the successor
	successor.left = nodeToDelete.left
	successor.right = nodeToDelete.right

	if successorParent == nodeToDelete {
		// If we were direct decentant of the deleted node then successors reattach successor right
		successor.right = oldSuccessorRight
	} else {
		// Our successor was not direct decendant of deleted node
		// Set successor's old parent's left to oldSuccessor's right
		successorParent.left = oldSuccessorRight
	}
}

func (this *BST[K, T]) findSuccessor(nodeToBeDeleted *bstNode[K, T]) (successor, successorParent *bstNode[K, T]) {
	successor = nodeToBeDeleted.right
	successorParent = nodeToBeDeleted

	// Because we chose the right side we need to keep going left after the first right jump
	// Continue until we can't go left anymore. That last node is our
	for successor != nil && successor.left != nil {
		successorParent = successor
		successor = successor.left
	}

	return successor, successorParent
}

func (this *BST[K, T]) Iterator() godatacollections.Iterator[T] {
	iter := &bstIterator[K, T]{nodeStack: stack.NewLStack[*bstNode[K, T]](nil), zeroValue: this.zeroValue}

	next := this.root

	for next != nil {
		iter.nodeStack.Push(next)
		next = next.left
	}
	// Need to ensure that the first value for next is preped
	iter.prepNext()

	return iter
}

type bstIterator[K, T any] struct {
	nodeStack *stack.LStack[*bstNode[K, T]]
	next      *bstNode[K, T]
	zeroValue T
}

func (this *bstIterator[K, T]) Close() error {
	return nil
}

func (this *bstIterator[K, T]) HasNext() bool {
	return this.next != nil
}

func (this *bstIterator[K, T]) prepNext() {

	// Save to ignore error as we are using a nil value for zero so we can tell when we have reached the end
	this.next, _ = this.nodeStack.Pop()
}

func (this *bstIterator[K, T]) Next() (T, error) {

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
