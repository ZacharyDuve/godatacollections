package tree

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func TestAVLTreeImplementsSet(t *testing.T) {
	// Values for K and T are arbitrary
	var _ godatacollections.Set[int, int] = &AVLTree[int, int]{}
}

func intAVL(zeroVal int) *AVLTree[int, int] {
	return &AVLTree[int, int]{zeroValue: zeroVal}
}

// -------------------------------------- Iterator ------------------------------------------

func TestAVLIteratorReturnsIterator(t *testing.T) {
	values := make([]int, 0)

	for i := 0; i < 1000; i++ {
		values = append(values, i)
	}

	avl := intAVL(-1)

	for _, curVal := range values {
		avl.Insert(curVal)
	}

	iter := avl.Iterator()

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
