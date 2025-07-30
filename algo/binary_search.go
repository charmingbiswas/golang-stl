package algo

import "cmp"

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
