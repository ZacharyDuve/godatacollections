package tree

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func TestBSTImplementsSet(t *testing.T) {
	var _ godatacollections.Set[string, string] = &BST[string, string]{}
}

func TestBSTStartsEmpty(t *testing.T) {
	bst, _ := NewBST(func(a, b int) int { return a - b }, func(a int) int { return a }, 0)

	if bst.root != nil {
		t.Fail()
	}
}

func TestBSTAddingOneOnlyAddsOneItem(t *testing.T) {
	bst, _ := NewBST(func(a, b int) int { return a - b }, func(a int) int { return a }, 0)

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
