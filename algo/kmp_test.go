package algo

import (
	"reflect"
	"testing"
)

func TestKnuthMorrisPrattStringMatching(t *testing.T) {
	testCases := []struct {
		input   string
		pattern string
		output  []int
	}{
		{"mississippi", "ippi", []int{7}},
		{"", "", []int{}},
	}

	for _, tc := range testCases {
		out := KnuthMorrisPrattStringMatching(tc.input, tc.pattern)
		if !reflect.DeepEqual(out, tc.output) {
			t.Error("KnuthMorrisPrattStringMatching: output mismatch!")
		}
	}
}
