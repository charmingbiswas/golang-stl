package algo

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{
			name:   "find element in middle",
			array:  []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "find first element",
			array:  []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "find last element",
			array:  []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
		{
			name:   "element not found - too small",
			array:  []int{1, 2, 3, 4, 5},
			target: 0,
			want:   -1,
		},
		{
			name:   "element not found - too large",
			array:  []int{1, 2, 3, 4, 5},
			target: 6,
			want:   -1,
		},
		{
			name:   "element not found - in between",
			array:  []int{1, 3, 5, 7, 9},
			target: 4,
			want:   -1,
		},
		{
			name:   "empty array",
			array:  []int{},
			target: 1,
			want:   -1,
		},
		{
			name:   "single element - found",
			array:  []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single element - not found",
			array:  []int{5},
			target: 3,
			want:   -1,
		},
		{
			name:   "two elements - find first",
			array:  []int{1, 2},
			target: 1,
			want:   0,
		},
		{
			name:   "two elements - find second",
			array:  []int{1, 2},
			target: 2,
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.array, tt.target)
			if got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchStrings(t *testing.T) {
	array := []string{"apple", "banana", "cherry", "date", "elderberry"}

	if got := BinarySearch(array, "cherry"); got != 2 {
		t.Errorf("BinarySearch(strings) = %v, want 2", got)
	}

	if got := BinarySearch(array, "fig"); got != -1 {
		t.Errorf("BinarySearch(strings) = %v, want -1", got)
	}
}

func TestLowerBound(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{
			name:   "find exact match",
			array:  []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "find first occurrence of duplicates",
			array:  []int{1, 2, 2, 2, 3, 4, 5},
			target: 2,
			want:   1,
		},
		{
			name:   "target smaller than all elements",
			array:  []int{2, 4, 6, 8, 10},
			target: 1,
			want:   0,
		},
		{
			name:   "target larger than all elements",
			array:  []int{1, 2, 3, 4, 5},
			target: 10,
			want:   -1,
		},
		{
			name:   "target between elements",
			array:  []int{1, 3, 5, 7, 9},
			target: 4,
			want:   2,
		},
		{
			name:   "empty array",
			array:  []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "single element - exact match",
			array:  []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single element - target smaller",
			array:  []int{5},
			target: 3,
			want:   0,
		},
		{
			name:   "single element - target larger",
			array:  []int{5},
			target: 7,
			want:   -1,
		},
		{
			name:   "all elements same - find first",
			array:  []int{3, 3, 3, 3, 3},
			target: 3,
			want:   0,
		},
		{
			name:   "find first element",
			array:  []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "find last element",
			array:  []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LowerBound(tt.array, tt.target)
			if got != tt.want {
				t.Errorf("LowerBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperBound(t *testing.T) {
	tests := []struct {
		name   string
		array  []int
		target int
		want   int
	}{
		{
			name:   "find element after target",
			array:  []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "find last occurrence of duplicates",
			array:  []int{1, 2, 2, 2, 3, 4, 5},
			target: 2,
			want:   3,
		},
		{
			name:   "target smaller than all elements",
			array:  []int{2, 4, 6, 8, 10},
			target: 1,
			want:   -1,
		},
		{
			name:   "target larger than all elements",
			array:  []int{1, 2, 3, 4, 5},
			target: 10,
			want:   4,
		},
		{
			name:   "target between elements",
			array:  []int{1, 3, 5, 7, 9},
			target: 4,
			want:   1,
		},
		{
			name:   "empty array",
			array:  []int{},
			target: 5,
			want:   -1,
		},
		{
			name:   "single element - exact match",
			array:  []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "single element - target smaller",
			array:  []int{5},
			target: 3,
			want:   -1,
		},
		{
			name:   "single element - target larger",
			array:  []int{5},
			target: 7,
			want:   0,
		},
		{
			name:   "all elements same - find last",
			array:  []int{3, 3, 3, 3, 3},
			target: 3,
			want:   4,
		},
		{
			name:   "target is first element",
			array:  []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "target is last element",
			array:  []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := UpperBound(tt.array, tt.target)
			if got != tt.want {
				t.Errorf("UpperBound() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearchFloats(t *testing.T) {
	array := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	if got := BinarySearch(array, 3.3); got != 2 {
		t.Errorf("BinarySearch(floats) = %v, want 2", got)
	}

	if got := BinarySearch(array, 3.5); got != -1 {
		t.Errorf("BinarySearch(floats) = %v, want -1", got)
	}
}

func TestLowerBoundFloats(t *testing.T) {
	array := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	if got := LowerBound(array, 3.3); got != 2 {
		t.Errorf("LowerBound(floats) = %v, want 2", got)
	}

	if got := LowerBound(array, 3.5); got != 3 {
		t.Errorf("LowerBound(floats) = %v, want 3", got)
	}
}

func TestUpperBoundFloats(t *testing.T) {
	array := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

	if got := UpperBound(array, 3.3); got != 2 {
		t.Errorf("UpperBound(floats) = %v, want 2", got)
	}

	if got := UpperBound(array, 3.5); got != 2 {
		t.Errorf("UpperBound(floats) = %v, want 2", got)
	}
}
