package queue

import (
	"reflect"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()
	val := new(Queue[int])
	if reflect.TypeOf(q) != reflect.TypeOf(val) {
		t.Error("type mismatch: func NewQueue returned wrong type")
	}
}

func TestPushBack(t *testing.T) {
	q := NewQueue[int]()
	var expected int
	var actual int

	q.PushBack(1)
	// check size of queue
	if expected != actual {
		t.Error("value mismatch: func PushBack did not update queue as expected")
	}

	// check the value in the queue
	expected = 1
	actual = q.Back()

	if expected != actual {
		t.Error("value mismatch: queue did not return the expected value")
	}
}
