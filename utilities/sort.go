package utilities

import "cmp"

func merge[T cmp.Ordered](array []T, low int, mid int, high int) {
	i := low
	j := mid + 1
	k := low

	tempArray := make([]T, len(array))

	for i <= mid && j <= high {
		if array[i] < array[j] {
			tempArray[k] = array[i]
			i++
			k++
		} else {
			tempArray[k] = array[j]
			j++
			k++
		}
	}

	for i <= mid {
		tempArray[k] = array[i]
		i++
		k++
	}

	for j <= high {
		tempArray[k] = array[j]
		j++
		k++
	}

	for i := low; i <= high; i++ {
		array[i] = tempArray[i]
	}
}

func mergeWithFunc[T any](array []T, low int, mid int, high int, less func(a, b T) bool) {
	i := low
	j := mid + 1
	k := low

	tempArray := make([]T, len(array))

	for i <= mid && j <= high {
		if less(array[i], array[j]) {
			tempArray[k] = array[i]
			i++
			k++
		} else {
			tempArray[k] = array[j]
			j++
			k++
		}
	}

	for i <= mid {
		tempArray[k] = array[i]
		i++
		k++
	}

	for j <= high {
		tempArray[k] = array[j]
		j++
		k++
	}

	for i := low; i <= high; i++ {
		array[i] = tempArray[i]
	}
}

func MergeSort[T cmp.Ordered](array []T, low int, high int) {
	if low < high {
		mid := low + (high-low)/2
		MergeSort(array, low, mid)
		MergeSort(array, mid+1, high)
		merge(array, low, mid, high)
	}
}

func MergeSortWithFunc[T any](array []T, low int, high int, less func(a, b T) bool) {
	if low < high {
		mid := low + (high-low)/2
		MergeSortWithFunc(array, low, mid, less)
		MergeSortWithFunc(array, mid+1, high, less)
		mergeWithFunc(array, low, mid, high, less)
	}
}
