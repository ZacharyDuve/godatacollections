package tree

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func intBST(zeroVal int) *BST[int, int] {
	bst, _ := NewBST(func(a, b int) int { return a - b }, func(a int) int { return a }, zeroVal)

	return bst
}

func TestBSTImplementsSet(t *testing.T) {
	var _ godatacollections.Set[int, int] = intBST(0)
}

func TestBSTStartsEmpty(t *testing.T) {
	bst := intBST(0)

	if bst.root != nil {
		t.Fail()
	}
}

func TestBSTAddingOneOnlyAddsOneItem(t *testing.T) {
	bst := intBST(0)

	bst.Insert(5)

	if bst.root == nil {
		t.Fail()
	}

	iter := bst.Iterator()
	count := 0

	for iter.HasNext() {
		iter.Next()
		count++
	}

}

func TestAddingIn3ValuesAndThenRemovingRootPromotesRightSideNode(t *testing.T) {
	bst := intBST(0)

	// First should be root
	bst.Insert(5)
	// Goes left
	bst.Insert(1)
	// Goes right
	bst.Insert(10)

	oldRoot := bst.root
	// Removes our root
	bst.Remove(oldRoot.key)

	if bst.root != oldRoot.right {
		t.Fail()
	}
}

func TestRemovingWhenNotingHasBeenAddedReturnsError(t *testing.T) {
	bst := intBST(0)

	err := bst.Remove(1)

	if err == nil {
		t.Fail()
	}
}

func TestAddingDuplicateReturnsError(t *testing.T) {
	bst := intBST(0)

	err := bst.Insert(1)

	if err != nil {
		t.Fail()
	}

	err = bst.Insert(1)

	if err == nil {
		t.Fail()
	}
}
