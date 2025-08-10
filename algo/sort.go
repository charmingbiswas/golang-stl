package algo

import (
	"cmp"
)

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

// Sorts an array of elements in ascending order.
func MergeSort[T cmp.Ordered](array []T, low int, high int) {
	if low < high {
		mid := low + (high-low)/2
		MergeSort(array, low, mid)
		MergeSort(array, mid+1, high)
		merge(array, low, mid, high)
	}
}
