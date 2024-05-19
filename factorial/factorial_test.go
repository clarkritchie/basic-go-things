package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, test := range tests {
		result := factorial(test.input)
		if result != test.expected {
			t.Errorf("Factorial of %d was incorrect, got: %d, want: %d.", test.input, result, test.expected)
		}
	}
}
