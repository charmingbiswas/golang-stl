package utilities

import (
	"testing"
)

func TestKmpPatternMatching(t *testing.T) {
	tests := []struct {
		input   string
		pattern string
		ok      bool
	}{
		{"mississippi", "ippi", true},
		{"mississipi", "", false},
	}

	for _, tt := range tests {
		_, ok := KmpPatternMatching(tt.input, tt.pattern)
		if ok != tt.ok {
			t.Error("KmpPatternMatching() returned wrong output")
		}
	}
}
