package main

import (
	"testing"

	"go23/ex03/asm1/ulti" // Make sure to import the correct package path
)

func TestCountRectangle(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected int
	}{
		{
			input: [][]int{
				{0, 0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1, 0},
				{0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			expected: 3,
		},
		{
			input: [][]int{
				{0, 0, 1, 1, 0, 1},
				{0, 0, 1, 1, 1, 0},
				{0, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0},
			},
			expected: 4,
		},
		// Add more test cases here if needed
	}

	for _, tc := range testCases {
		result := ulti.CountRectangle(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %d, but got %d", tc.expected, result)
		}
	}
}