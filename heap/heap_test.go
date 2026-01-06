package heap

import (
	"testing"
)

func TestNewMaxHeap(t *testing.T) {
	t.Run("initialize a new max heap of integers and test max heap property", func(t *testing.T) {
		t.Parallel()
		h := NewMaxHeap[int]()
		if h == nil {
			t.Fatal("NewMaxHeap returned nil")
		}

		if !h.IsEmpty() {
			t.Error("New heap should be empty")
		}

		if h.Size() != 0 {
			t.Errorf("Expected heap size to be 0, got %d", h.Size())
		}

		mockValues := []int{5, 3, 7, 1, 9, 2, 8, 4, 6}

		for _, val := range mockValues {
			h.Push(val)
		}

		result := make([]int, 0)
		for !h.IsEmpty() {
			result = append(result, h.Top())
			h.Pop()
		}

		expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}

		for index := range result {
			if result[index] != expected[index] {
				t.Errorf("Max Heap property violated at index %d: want %d, got %d", index, expected[index], result[index])
			}
		}
	})

	t.Run("initialize a new max heap of strings and test max heap property", func(t *testing.T) {
		t.Parallel()
		h := NewMaxHeap[string]()
		if h == nil {
			t.Fatal("NewMaxHeap returned nil")
		}

		if !h.IsEmpty() {
			t.Error("New heap should be empty")
		}

		if h.Size() != 0 {
			t.Errorf("Expected heap size to be 0, got %d", h.Size())
		}

		mockValues := []string{"list", "apple", "car", "dog"}

		for _, val := range mockValues {
			h.Push(val)
		}

		result := make([]string, 0)
		for !h.IsEmpty() {
			result = append(result, h.Top())
			h.Pop()
		}

		expected := []string{"list", "dog", "car", "apple"}

		for index := range result {
			if result[index] != expected[index] {
				t.Errorf("Max Heap property violated at index %d: want %s, got %s", index, expected[index], result[index])
			}
		}
	})
}

func TestNewMinHeap(t *testing.T) {
	t.Run("initialize a new min heap of integers and test min heap property", func(t *testing.T) {
		t.Parallel()
		h := NewMinHeap[int]()
		if h == nil {
			t.Fatal("NewMinHeap returned nil")
		}

		if !h.IsEmpty() {
			t.Error("New heap should be empty")
		}

		if h.Size() != 0 {
			t.Errorf("Expected heap size to be 0, got %d", h.Size())
		}

		mockValues := []int{5, 3, 7, 1, 9, 2, 8, 4, 6}

		for _, val := range mockValues {
			h.Push(val)
		}

		result := make([]int, 0)
		for !h.IsEmpty() {
			result = append(result, h.Top())
			h.Pop()
		}

		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

		for index := range result {
			if result[index] != expected[index] {
				t.Errorf("Max Heap property violated at index %d: want %d, got %d", index, expected[index], result[index])
			}
		}
	})

	t.Run("initialize a new min heap of strings and test min heap property", func(t *testing.T) {
		t.Parallel()
		h := NewMinHeap[string]()
		if h == nil {
			t.Fatal("NewMinHeap returned nil")
		}

		if !h.IsEmpty() {
			t.Error("New heap should be empty")
		}

		if h.Size() != 0 {
			t.Errorf("Expected heap size to be 0, got %d", h.Size())
		}

		mockValues := []string{"list", "apple", "car", "dog"}

		for _, val := range mockValues {
			h.Push(val)
		}

		result := make([]string, 0)
		for !h.IsEmpty() {
			result = append(result, h.Top())
			h.Pop()
		}

		expected := []string{"apple", "car", "dog", "list"}

		for index := range result {
			if result[index] != expected[index] {
				t.Errorf("Max Heap property violated at index %d: want %s, got %s", index, expected[index], result[index])
			}
		}
	})
}

func TestNewHeapWithFunc(t *testing.T) {
	t.Run("testing NewHeapWithFunc with custom data type one", func(t *testing.T) {
		t.Parallel()
		type mockPair struct {
			key   int
			value any
		}

		comparator := func(a, b mockPair) bool {
			return a.key < b.key // min heap
		}

		h := NewHeapWithFunc(comparator)

		mockValues := []mockPair{
			{key: 1, value: 1},
			{key: 5, value: 5},
			{key: 2, value: 2},
			{key: 9, value: 9},
			{key: 3, value: 3},
		}

		for _, val := range mockValues {
			h.Push(val)
		}

		result := make([]mockPair, 0)

		for !h.IsEmpty() {
			result = append(result, h.Top())
			h.Pop()
		}

		expectedOutput := []mockPair{
			{key: 1, value: 1},
			{key: 2, value: 2},
			{key: 3, value: 3},
			{key: 5, value: 5},
			{key: 9, value: 9},
		}

		for index := range result {
			if result[index] != expectedOutput[index] {
				t.Errorf("heap property failed: want %#v, got %#v", expectedOutput[index], result[index])
			}
		}
	})
}
