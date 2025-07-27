package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	testCasesWithInts := []struct {
		input  []int
		low    int
		high   int
		output []int
	}{
		{[]int{1, 4, 2, 8, 11, 5, 2, 1, 5}, 0, 9, []int{1, 1, 2, 2, 4, 5, 5, 8, 11}},
	}

	for _, tc := range testCasesWithInts {
		MergeSort(tc.input, tc.low, tc.high-1)
		fmt.Println(tc.input, tc.output)
		if !reflect.DeepEqual(tc.output, tc.input) {
			t.Error("MergeSort: mismatch with expected output!")
		}
	}
}
