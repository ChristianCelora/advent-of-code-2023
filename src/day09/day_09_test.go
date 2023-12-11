package main

import "testing"

func TestGetNextSubsequence(t *testing.T) {
	tests := []struct {
		input_sequence []int
		expected_value []int
	}{
		{
			input_sequence: []int{0, 3, 6, 9, 12, 15},
			expected_value: []int{3, 3, 3, 3, 3},
		},
		{
			input_sequence: []int{3, 3, 3, 3, 3},
			expected_value: []int{0, 0, 0, 0},
		},
		{
			input_sequence: []int{-3, 3, -1, 3, 3},
			expected_value: []int{6, 4, 4, 0},
		},
	}

	for _, test := range tests {
		subsequence := GetNextSubsequence(test.input_sequence)
		for i, v := range subsequence {
			if v != test.expected_value[i] {
				t.Fatalf("Expected %v value. actual %v for input %v", test.expected_value, subsequence, test.input_sequence)
			}
		}
	}
}
