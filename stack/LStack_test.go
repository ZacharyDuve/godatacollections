package stack

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func TestDLStackImplementsStack(t *testing.T) {
	var _ godatacollections.Stack[int] = NewLStack[int](0)
}

func TestDLStackPoppingOnEmptyStackReturnsZeroValue(t *testing.T) {
	zeroValue := 666
	s := NewLStack(zeroValue)

	pVal, _ := s.Pop()

	if pVal != zeroValue {
		t.Fail()
	}
}

func TestDLStackPoppingOnEmptyStackReturnsError(t *testing.T) {
	zeroValue := 666
	s := NewLStack(zeroValue)

	_, err := s.Pop()

	if !godatacollections.IsEmptyError(err) {
		t.Fail()
	}
}

func TestDLStackMaintainsLIFOOrder(t *testing.T) {
	zeroValue := -1
	first := 7
	second := 2

	s := NewLStack(zeroValue)
	s.Push(first)
	s.Push(second)
	pV, _ := s.Pop()

	if pV != second {
		t.Fail()
	}

	pV, _ = s.Pop()

	if pV != first {
		t.Fail()
	}

	pV, _ = s.Pop()

	if pV != zeroValue {
		t.Fail()
	}
}
