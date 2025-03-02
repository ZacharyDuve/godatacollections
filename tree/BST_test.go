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

func TestBSTReturnsErrorIfMissingKeyCompFunc(t *testing.T) {
	bst, err := NewBST[int, int](nil, func(i int) int { return i }, -1)

	if err == nil {
		t.Fail()
	}

	if bst != nil {
		t.Fail()
	}
}

func TestBSTReturnsErrorIfMissingTToKFunc(t *testing.T) {
	bst, err := NewBST[int, int](func(i1, i2 int) int { return i1 - i2 }, nil, -1)

	if err == nil {
		t.Fail()
	}

	if bst != nil {
		t.Fail()
	}
}

// -------------------------------------- Adding ------------------------------------------

func TestBSTAddingOneOnlyAddsOneItem(t *testing.T) {
	bst := intBST(0)

	err := bst.Insert(5)

	if err != nil {
		t.Fatal("expected no error on insert of root node")
	}

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

func TestAddingChildNodeNonDuplicateReturnsNoError(t *testing.T) {
	bst := intBST(-1)

	err := bst.Insert(6)

	if err != nil {
		t.Fatal(err)
	}

	// Insert new different node shouldn't return error
	err = bst.Insert(4)

	if err != nil {
		t.Fatal(err)
	}
}

// -------------------------------------- Removing ------------------------------------------

// Root Specific Tests
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
	if bst.tToKFunc(bst.root.t) != 1 {
		t.Fatal("expected root to be key of 1")
	}
	bst.Insert(2)
	if bst.tToKFunc(bst.root.t) != 1 {
		t.Fatal("root should not have changed")
	}

	if bst.tToKFunc(bst.root.right.t) != 2 {
		t.Fatal("expected root's right child to have key of 2")
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
	if bst.tToKFunc(bst.root.t) != 1 {
		t.Fatal("expected root to be key of 1")
	}
	bst.Insert(0)
	if bst.tToKFunc(bst.root.t) != 1 {
		t.Fatal("root should not have changed")
	}

	if bst.tToKFunc(bst.root.left.t) != 0 {
		t.Fatal("expecter root's left child to have key of 2")
	}

	if bst.root.left == nil {
		t.Fatal("expected to have a left child but there was none")
	}

	bst.Remove(0)

	if bst.root.left != nil {
		t.Fatal("expected root's left to have been removed")
	}
}

func TestRemovingRootWhenItHasOneChildOnLeftReplacesRootWithLeftChild(t *testing.T) {
	bst := intBST(-1)

	rootKey := 5
	childKey := 2
	bst.Insert(rootKey)
	bst.Insert(childKey)

	bst.Remove(rootKey)

	if bst.tToKFunc(bst.root.t) != childKey {
		t.Fail()
	}

	if bst.root.left != nil {
		t.Fail()
	}

	if bst.root.right != nil {
		t.Fail()
	}
}

func TestRemovingRootWhenItHasOneChildOnRightReplacesRootWithRightChild(t *testing.T) {
	bst := intBST(-1)

	rootKey := 5
	childKey := 7
	bst.Insert(rootKey)
	bst.Insert(childKey)

	bst.Remove(rootKey)

	if bst.tToKFunc(bst.root.t) != childKey {
		t.Fail()
	}

	if bst.root.left != nil {
		t.Fail()
	}

	if bst.root.right != nil {
		t.Fail()
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
	bst.Remove(bst.tToKFunc(oldRoot.t))

	if bst.root != oldRoot.right {
		t.Fail()
	}
}

// Blank remove
func TestRemovingWhenNotingHasBeenAddedReturnsError(t *testing.T) {
	bst := intBST(0)

	err := bst.Remove(1)

	if err == nil {
		t.Fail()
	}
}

// Single child remove non-root node

func TestRemovingNonRootNodeWithSingleLeftChildRemovesNodeAndPromotesChild(t *testing.T) {
	bst := intBST(-1)

	nodeKeyDelete := 7
	nodeChildKey := 6

	bst.Insert(5)
	//This node is to be deleted
	bst.Insert(nodeKeyDelete)
	//This should be previouses left child
	bst.Insert(nodeChildKey)

	//if bst.root.right.left.key != nodeChildKey {
	if bst.tToKFunc(bst.root.right.left.t) != nodeChildKey {
		t.Fatal("Failed to setup correctly")
	}

	bst.Remove(nodeKeyDelete)

	if bst.tToKFunc(bst.root.right.t) != nodeChildKey {
		t.Fatal("Child Failed to be promoted")
	}
}

func TestRemovingNonRootNodeWithSingleRightChildRemovesNodeAndPromotesChild(t *testing.T) {
	bst := intBST(-1)

	nodeKeyDelete := 3
	nodeChildKey := 4

	bst.Insert(5)
	//This node is to be deleted
	bst.Insert(nodeKeyDelete)
	//This should be previouses left child
	bst.Insert(nodeChildKey)

	if bst.tToKFunc(bst.root.left.right.t) != nodeChildKey {
		t.Fatal("Failed to setup correctly")
	}

	bst.Remove(nodeKeyDelete)

	if bst.tToKFunc(bst.root.left.t) != nodeChildKey {
		t.Fatal("Child Failed to be promoted")
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

func TestContainsWorksForLeftChildNodes(t *testing.T) {
	bst := intBST(-1)

	leftKey := 25
	bst.Insert(50)

	bst.Insert(leftKey)

	if !bst.Contains(leftKey) {
		t.Fail()
	}
}

func TestContainsWorksForRightChildNodes(t *testing.T) {
	bst := intBST(-1)

	leftKey := 75
	bst.Insert(50)

	bst.Insert(leftKey)

	if !bst.Contains(leftKey) {
		t.Fail()
	}
}

// -------------------------------------- Get By Key ------------------------------------------

func TestGetByKeyWithNoNodeMatchingReturnsZeroValue(t *testing.T) {
	zeroValue := -1
	bst := intBST(zeroValue)

	bst.Insert(4)
	bst.Insert(8)
	bst.Insert(5)

	v, err := bst.GetByKey(2)

	if v != zeroValue {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

func TestGetByKeyWithNodeMatchingReturnsValue(t *testing.T) {
	zeroValue := -1
	bst := intBST(zeroValue)

	bst.Insert(4)
	bst.Insert(8)
	bst.Insert(5)

	v, err := bst.GetByKey(8)

	if v != 8 {
		t.Fail()
	}

	if err != nil {
		t.Fail()
	}
}

// -------------------------------------- Iterator ------------------------------------------

func TestIteratorReturnsIterator(t *testing.T) {
	values := make([]int, 0)

	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	bst := intBST(-1)

	for _, curVal := range values {
		bst.Insert(curVal)
	}

	iter := bst.Iterator()

	if iter == nil {
		t.Fail()
	}

	valuesFromIter := make([]int, 0)

	for iter.HasNext() {
		curVal, err := iter.Next()

		if err != nil {
			t.Fail()
		}
		valuesFromIter = append(valuesFromIter, curVal)
	}

	if len(valuesFromIter) != len(values) {
		t.Fail()
	}

	for _, curValue := range values {
		if !containsInSlice(curValue, valuesFromIter) {
			t.Fatal("Failed to find value from iterator in slice")
		}
	}
}
