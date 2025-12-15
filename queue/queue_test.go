package queue

import (
	"errors"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int]()

	if q == nil {
		t.Fatal("NewQueue returned nil")
	}

	if q.data == nil {
		t.Fatal("Queue data is nil")
	}

	if !q.IsEmpty() {
		t.Error("New queue should be empty\n")
	}

	if q.Size() != 0 {
		t.Errorf("New queue size should be 0, got %d\n", q.Size())
	}
}

func TestPushFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushFront(1)

	if q.Size() != 1 {
		t.Errorf("Queue size should be 1, got %d", q.Size())
	}

	q.PushFront(2)
	q.PushFront(3)

	if q.Size() != 3 {
		t.Errorf("Queue size should be 3, got %d", q.Size())
	}

	val, err := q.Front()
	if err != nil {
		t.Fatalf("Front() returned error %v", err)
	}

	if val != 3 {
		t.Errorf("Expected Front() to return 3, got %d", val)
	}
}

func TestPushBack(t *testing.T) {
	q := NewQueue[int]()

	q.PushBack(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}

	q.PushBack(2)
	q.PushBack(3)

	if q.Size() != 3 {
		t.Errorf("Expected size 3, got %d", q.Size())
	}

	back, err := q.Back()
	if err != nil {
		t.Fatalf("Back() returned error: %v", err)
	}
	if back != 3 {
		t.Errorf("Expected back to be 3, got %d", back)
	}
}

func TestPopFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)

	if q.Size() != 3 {
		t.Errorf("Expected queue size 3, got %d", q.Size())
	}

	err := q.PopFront()
	if err != nil {
		t.Fatalf("PopFront() returned error %v", err)
	}

	if q.Size() != 2 {
		t.Errorf("Expected queue size 2, got %d", q.Size())
	}

	err = q.PopFront()
	if err != nil {
		t.Fatalf("PopFront() returned error %v", err)
	}

	err = q.PopFront()
	if err != nil {
		t.Fatalf("PopFront() returned error %v", err)
	}

	if !q.IsEmpty() {
		t.Fatalf("Expected queue to be empty, got %d", q.Size())
	}
}

func TestPopBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.PushBack(2)
	q.PushBack(3)

	err := q.PopBack()
	if err != nil {
		t.Fatalf("PopBack() returned error: %v", err)
	}

	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	err = q.PopBack()
	if err != nil {
		t.Fatalf("PopBack() returned error: %v", err)
	}

	err = q.PopBack()
	if err != nil {
		t.Fatalf("PopBack() returned error: %v", err)
	}

	if !q.IsEmpty() {
		t.Errorf("Queue should be empty")
	}
}

func TestPopFrontEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	err := q.PopFront()
	if err == nil {
		t.Fatal("Expected error when popping from empty queue")
	}
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}
}

func TestPopBackEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	err := q.PopBack()
	if err == nil {
		t.Fatal("Expected error when popping from empty queue")
	}
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}
}

func TestFront(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.PushBack(2)

	front, err := q.Front()
	if err != nil {
		t.Fatalf("Front() returned error: %v", err)
	}
	if front != 1 {
		t.Errorf("Expected 1, got %d", front)
	}

	// Verify Front() doesn't remove the element
	if q.Size() != 2 {
		t.Errorf("Front() should not modify size")
	}
}

func TestBack(t *testing.T) {
	q := NewQueue[int]()
	q.PushBack(1)
	q.PushBack(2)

	back, err := q.Back()
	if err != nil {
		t.Fatalf("Back() returned error: %v", err)
	}
	if back != 2 {
		t.Errorf("Expected 2, got %d", back)
	}

	// Verify Back() doesn't remove the element
	if q.Size() != 2 {
		t.Errorf("Back() should not modify size")
	}
}

func TestFrontEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	val, err := q.Front()
	if err == nil {
		t.Fatal("Expected error when calling Front() on empty queue")
	}
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

func TestBackEmptyQueue(t *testing.T) {
	q := NewQueue[int]()

	val, err := q.Back()
	if err == nil {
		t.Fatal("Expected error when calling Back() on empty queue")
	}
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("Expected ErrEmptyQueue, got %v", err)
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

func TestMixedOperations(t *testing.T) {
	q := NewQueue[int]()

	q.PushBack(1)
	q.PushFront(0)
	q.PushBack(2)

	front, _ := q.Front()
	if front != 0 {
		t.Errorf("Expected front 0, got %d", front)
	}

	back, _ := q.Back()
	if back != 2 {
		t.Errorf("Expected back 2, got %d", back)
	}

	q.PopFront()
	q.PopBack()

	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}

	front, _ = q.Front()
	back, _ = q.Back()
	if front != 1 || back != 1 {
		t.Errorf("Expected both front and back to be 1")
	}
}

func TestIsEmpty(t *testing.T) {
	q := NewQueue[int]()

	if !q.IsEmpty() {
		t.Error("New queue should be empty")
	}

	q.PushBack(1)
	if q.IsEmpty() {
		t.Error("Queue with elements should not be empty")
	}

	q.PopFront()
	if !q.IsEmpty() {
		t.Error("Queue should be empty after removing all elements")
	}
}

func TestSize(t *testing.T) {
	q := NewQueue[int]()

	sizes := []int{0, 1, 2, 3, 2, 1, 0}
	operations := []func(){
		func() { /* Initial state */ },
		func() { q.PushBack(1) },
		func() { q.PushBack(2) },
		func() { q.PushFront(0) },
		func() { q.PopFront() },
		func() { q.PopBack() },
		func() { q.PopFront() },
	}

	for i, op := range operations {
		op()
		if q.Size() != sizes[i] {
			t.Errorf("After operation %d: expected size %d, got %d", i, sizes[i], q.Size())
		}
	}
}

func TestGenericString(t *testing.T) {
	q := NewQueue[string]()

	q.PushBack("hello")
	q.PushBack("world")
	q.PushFront("hi")

	front, err := q.Front()
	if err != nil {
		t.Fatalf("Front() error: %v", err)
	}
	if front != "hi" {
		t.Errorf("Expected 'hi', got '%s'", front)
	}

	err = q.PopBack()
	if err != nil {
		t.Fatalf("PopBack() error: %v", err)
	}
}

func TestGenericStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	q := NewQueue[Person]()

	alice := Person{"Alice", 30}
	bob := Person{"Bob", 25}

	q.PushBack(alice)
	q.PushBack(bob)

	front, err := q.Front()
	if err != nil {
		t.Fatalf("Front() error: %v", err)
	}
	if front.Name != "Alice" || front.Age != 30 {
		t.Errorf("Expected Alice(30), got %v", front)
	}

	err = q.PopFront()
	if err != nil {
		t.Fatalf("PopFront() error: %v", err)
	}
}

func TestLargeQueue(t *testing.T) {
	q := NewQueue[int]()
	n := 10000

	// Push many elements
	for i := 0; i < n; i++ {
		q.PushBack(i)
	}

	if q.Size() != n {
		t.Errorf("Expected size %d, got %d", n, q.Size())
	}

	// Pop all elements
	for i := 0; i < n; i++ {
		err := q.PopFront()
		if err != nil {
			t.Fatalf("PopFront() error at %d: %v", i, err)
		}
	}

	if !q.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

func TestDequePattern(t *testing.T) {
	q := NewQueue[int]()

	// Add to both ends
	q.PushBack(2)
	q.PushFront(1)
	q.PushBack(3)
	q.PushFront(0)
	// Queue: [0, 1, 2, 3]

	// Remove from both ends
	q.PopFront()
	q.PopBack()

	// Queue should now be [1, 2]
	if q.Size() != 2 {
		t.Errorf("Expected size 2, got %d", q.Size())
	}

	front, _ := q.Front()
	back, _ := q.Back()
	if front != 1 || back != 2 {
		t.Errorf("Expected [1, 2], got [%d, %d]", front, back)
	}
}
