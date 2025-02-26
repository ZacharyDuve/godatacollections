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

func TestThatDequeueRemovesItemFromQueue(t *testing.T) {
	q := NewLQueue(0)

	value := 1

	q.Enqueue(value)

	v, _ := q.Dequeue()

	// Value should be the value that had in the queue
	if value != v {
		t.Fail()
	}

	// Head should no be cleared
	if q.head != nil {
		t.Fail()
	}

	// Since length was 1 tail should be cleared too
	if q.tail != nil {
		t.Fail()
	}
}

func TestThatDequeueWithQueueNotEmptyReturnsNoError(t *testing.T) {
	q := NewLQueue(0)

	value := 1

	q.Enqueue(value)

	_, err := q.Dequeue()

	if err != nil {
		t.Fail()
	}
}

func TestThatDequeueWithEmptyQueueReturnsError(t *testing.T) {
	q := NewLQueue(0)

	_, err := q.Dequeue()

	if err == nil {
		t.Fail()
	}
}

func TestQueueMaintainsFiFoOrderingForItems(t *testing.T) {
	q := NewLQueue(0)

	items := []int{9, 19, 3, 12}

	for _, curItem := range items {
		q.Enqueue(curItem)
	}

	i := 0
	for curItem, err := q.Dequeue(); err != nil; curItem, err = q.Dequeue() {

		if items[i] != curItem {
			t.Fail()
		}

		i++
	}
}
