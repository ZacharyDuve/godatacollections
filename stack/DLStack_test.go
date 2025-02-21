package stack

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func TestDLStackImplementsStack(t *testing.T) {
	var _ godatacollections.Stack[int] = NewDLStack[int](0)
}

func TestDLStackPoppingOnEmptyStackReturnsZeroValue(t *testing.T) {
	zeroValue := 666
	s := NewDLStack(zeroValue)

	pVal := s.Pop()

	if pVal != zeroValue {
		t.Fail()
	}
}

func TestDLStackMaintainsLIFOOrder(t *testing.T) {
	zeroValue := -1
	first := 7
	second := 2

	s := NewDLStack(zeroValue)
	s.Push(first)
	s.Push(second)

	if s.Pop() != second {
		t.Fail()
	}

	if s.Pop() != first {
		t.Fail()
	}

	if s.Pop() != zeroValue {
		t.Fail()
	}
}
