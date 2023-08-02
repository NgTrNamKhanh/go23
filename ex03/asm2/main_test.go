package main

import (
	"go23/ex03/asm2/ulti"
	"testing"
)

func TestNumDifferentInteger(t *testing.T) {
	testCases := []struct {
		input          string
		expectedlength int
		expectedarray  []int
	}{
		{
			input:          "a123bc34d8ef34",
			expectedlength: 3,
			expectedarray:  []int{123, 34, 8},
		},
		{
			input:          "asdnf23asd33ss1",
			expectedlength: 3,
			expectedarray:  []int{23, 33, 1},
		},
		{
			input:          "abc02asdc30",
			expectedlength: 2,
			expectedarray:  []int{02, 30},
		},
	}
	for _, tc := range testCases {
		resultlength, resultarray := ulti.NumDifferentInteger(tc.input)
		if resultlength != tc.expectedlength && !equal(resultarray, tc.expectedarray) {
			t.Errorf("Expected %d, but got %d", tc.expectedarray, resultarray)
		}
	}
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
