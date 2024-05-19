package main

import (
	"testing"
)

// Define a mockShape struct implementing Shape interface for testing
type mockShape struct {
	areaValue float64
}

func (m mockShape) area() float64 {
	return m.areaValue
}

func TestGetArea(t *testing.T) {
	testCases := []struct {
		shape         Shape
		expectedPrint float64
	}{
		{mockShape{10}, 10.000000},
		{mockShape{20}, 20.000000},
	}

	for _, tc := range testCases {
		output := getArea(tc.shape)
		if output != tc.expectedPrint {
			t.Errorf("printArea output incorrect. Expected: %f, Got: %f", tc.expectedPrint, output)
		}
	}
}
