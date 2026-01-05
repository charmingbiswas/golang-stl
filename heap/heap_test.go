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
