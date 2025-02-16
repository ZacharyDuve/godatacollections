package tree

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func intBST(zeroVal int) *BST[int, int] {
	bst, _ := NewBST(func(a, b int) int { return a - b }, func(a int) int { return a }, zeroVal)

	return bst
}

// -------------------------------------- Implements Set ------------------------------------------

func TestBSTImplementsSet(t *testing.T) {
	var _ godatacollections.Set[int, int] = intBST(0)
}

// -------------------------------------- Initialization ------------------------------------------

func TestBSTStartsEmpty(t *testing.T) {
	bst := intBST(0)

	if bst.root != nil {
		t.Fail()
	}
}

// -------------------------------------- Adding ------------------------------------------

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

// -------------------------------------- Removing ------------------------------------------
func TestRemovingRootDeletesRoot(t *testing.T) {
	bst := intBST(0)

	bst.Insert(1)
	bst.Remove(1)

	if bst.root != nil {
		t.Fail()
	}
}

func TestRemovingRightChildOfRootRemovesChild(t *testing.T) {
	bst := intBST(0)

	bst.Insert(1)
	if bst.root.right != nil {
		t.Fatal("expect to have pre add of right to not have a right child")
	}
	if bst.root.key != 1 {
		t.Fatal("expected root to be key of 1")
	}
	bst.Insert(2)
	if bst.root.key != 1 {
		t.Fatal("root should not have changed")
	}

	if bst.root.right.key != 2 {
		t.Fatal("expecter root's right child to have key of 2")
	}

	if bst.root.right == nil {
		t.Fatal("expected to have a right child but there was none")
	}

	bst.Remove(2)

	if bst.root.right != nil {
		t.Fatal("expected root's right to have been removed")
	}
}

func TestRemovingRootsLeftChildRemovesCorrectly(t *testing.T) {
	bst := intBST(0)

	bst.Insert(1)
	if bst.root.right != nil {
		t.Fatal("expect to have pre add of right to not have a right child")
	}
	if bst.root.key != 1 {
		t.Fatal("expected root to be key of 1")
	}
	bst.Insert(0)
	if bst.root.key != 1 {
		t.Fatal("root should not have changed")
	}

	if bst.root.left.key != 0 {
		t.Fatal("expecter root's right child to have key of 2")
	}

	if bst.root.left == nil {
		t.Fatal("expected to have a right child but there was none")
	}

	bst.Remove(0)

	if bst.root.left != nil {
		t.Fatal("expected root's right to have been removed")
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

// -------------------------------------- Contains ------------------------------------------

func TestContainsReturnsFalseWhenItemNotInBST(t *testing.T) {
	bst := intBST(0)

	// If bst actually contains 0 since it never was added then fail
	if bst.Contains(0) {
		t.Fail()
	}
}

func TestContainsReturnsTrueWhenItemHasBeenAddedToBST(t *testing.T) {
	bst := intBST(0)

	bst.Insert(0)

	if !bst.Contains(0) {
		t.Fail()
	}
}

func TestContainsReturnsFalseAfterItemHasBeenRemovedFromBST(t *testing.T) {
	bst := intBST(-1)

	bst.Insert(100)
	bst.Remove(100)

	if bst.Contains(100) {
		t.Fail()
	}
}

// -------------------------------------- Iterator ------------------------------------------

func TestIteratorReturnsIterator(t *testing.T) {
	t.Fatal("TEST NOT IMPLEMENTED")
}
