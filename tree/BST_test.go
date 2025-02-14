package tree

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func TestBSTImplementsSet(t *testing.T) {
	var _ godatacollections.Set[string, string] = &BST[string, string]{}
}
