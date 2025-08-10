package algo

import (
	"cmp"
)

// Standard binary search algorithm.
// Works with in-built comparable types.
// Returns index of the target element.
// Returns -1 if not found.
func BinarySearch[T cmp.Ordered](array []T, target T) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2

		if target == array[mid] {
			return mid
		} else if target < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

// Standard binary search lowerbound algorithm.
// Works with in-built comparable types.
// Returns index of the target element.
// Returns -1 if not found.
func LowerBound[T cmp.Ordered](array []T, target T) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2
		if array[mid] < target {
			low = mid + 1
		} else if array[mid] > target {
			high = mid - 1
		} else if array[mid] == target {
			high = mid - 1
		}
	}
	if low >= len(array) || array[low] != target {
		return -1
	}

	return low
}

// Standard binary search upperbound algorithm.
// Works with in-built comparable types.
// Returns index of the target element.
// Returns -1 if not found.
func UpperBound[T cmp.Ordered](array []T, target T) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2

		if array[mid] < target {
			low = mid + 1
		} else if array[mid] > target {
			high = mid - 1
		} else if array[mid] == target {
			low = mid + 1
		}
	}

	if high < 0 || array[high] != target {
		return -1
	}

	return high
}
