package algo

import "cmp"

// Standard binary search algorithm.
// Works with in-built comparable types.
func BinarySearch[T cmp.Ordered](array []T, target T) (int, bool) {
	low := 0
	high := len(array) - 1

	for low < high {
		mid := low + (high-low)/2

		if target == array[mid] {
			return mid, true
		} else if target < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}

// Standard binary search lowerbound algorithm.
// Works with in-built comparable types.
func LowerBound[T cmp.Ordered](array []T, target T) int {
	low := 0
	high := len(array) - 1

	for low < high {
		mid := low + (high-low)/2

		if array[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// Standard binary search upperbound algorithm.
// Works with in-built comparable types.
func UpperBound[T cmp.Ordered](array []T, target T) int {
	low := 0
	high := len(array)

	for low < high {
		mid := low + (high-low)/2

		if array[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// Standard binary search lowerbound algorithm.
// Takes a 'comparator' function which decides how two custom values are compared.
// If the two values are equal, comparator function should return 0.
// If A is less than B, comparator should return -1.
// If A is greater than B, comparator should return 1.
func LowerBoundWithFunc[T any](array []T, target T, comparator func(A, B T) int) int {
	low := 0
	high := len(array) - 1

	for low < high {
		mid := low + (high-low)/2
		if comparator(array[mid], target) == -1 {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}

// Standard binary search upperbound algorithm.
// Takes a 'comparator' function which decides how two custom values are compared.
// If the two values are equal, comparator function should return 0.
// If A is less than B, comparator should return -1.
// If A is greater than B, comparator should return 1.
func UpperBoundWithFunc[T any](array []T, target T, comparator func(A, B T) int) int {
	low, high := 0, len(array)-1

	for low < high {
		mid := low + (high-low)/2

		if comparator(array[mid], target) == 0 || comparator(array[mid], target) == -1 {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
