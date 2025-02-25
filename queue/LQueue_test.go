package queue

import (
	"testing"

	"github.com/ZacharyDuve/godatacollections"
)

func TestLQueueImplementsQueue(t *testing.T) {
	var _ godatacollections.Queue[int] = &LQueue[int]{}
}

func TestThatLQueueStartsEmpty(t *testing.T) {
	q := NewLQueue[*int](nil)

	if q.head != nil {
		t.Fail()
	}

	if q.tail != nil {
		t.Fail()
	}
}

func TestThatEnqueueAddsItemsToEmptyQueue(t *testing.T) {
	q := NewLQueue(0)

	q.Enqueue(1)

	if q.head.t != 1 {
		t.Fail()
	}

	if q.tail.t != 1 {
		t.Fail()
	}
}

func TestThatDequeueRemovesItemFromQueue(t *testing.T)
