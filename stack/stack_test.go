package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	t.Run("initializing new stack with int data", func(t *testing.T) {
		t.Parallel()
		st := NewStack[int]()
		if st == nil {
			t.Fatal("NewStack returned nil")
		}

		if !st.IsEmpty() {
			t.Error("New stack should be empty")
		}

		if st.Size() != 0 {
			t.Errorf("Expected stack size to be 0, got %d", st.Size())
		}

		mockValues := []int{1, 2, 3, 4}

		for _, val := range mockValues {
			st.Push(val)
		}

		var result []int

		for !st.IsEmpty() {
			result = append(result, st.Top())
			st.Pop()
		}

		expectedOutput := []int{4, 3, 2, 1}

		for index := range result {
			if result[index] != expectedOutput[index] {
				t.Errorf("stack property violated: want %d, got %d", expectedOutput[index], result[index])
			}
		}
	})

	t.Run("initializing new stack with custom data type", func(t *testing.T) {
		t.Parallel()
		type mockStruct struct {
			key   int
			value any
		}

		st := NewStack[mockStruct]()
		if st == nil {
			t.Fatal("NewStack returned nil")
		}

		if !st.IsEmpty() {
			t.Error("New stack should be empty")
		}

		if st.Size() != 0 {
			t.Errorf("Expected stack size to be 0, got %d", st.Size())
		}

		mockValues := []mockStruct{
			{key: 1, value: "a"},
			{key: 2, value: "b"},
			{key: 3, value: "c"},
		}

		for _, val := range mockValues {
			st.Push(val)
		}

		var result []mockStruct

		for !st.IsEmpty() {
			result = append(result, st.Top())
			st.Pop()
		}

		expectedOutput := []mockStruct{
			{key: 3, value: "c"},
			{key: 2, value: "b"},
			{key: 1, value: "a"},
		}

		for index := range result {
			if result[index] != expectedOutput[index] {
				t.Errorf("stack property violated: want %#v, got %#v", expectedOutput[index], result[index])
			}
		}
	})
}
