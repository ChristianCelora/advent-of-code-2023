package main

import "testing"

func TestGetHandType(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value int
	}{
		{
			input_line:     "23459",
			expected_value: 0,
		},
		{
			input_line:     "23452",
			expected_value: 1,
		},
		{
			input_line:     "T3T32",
			expected_value: 2,
		},
		{
			input_line:     "T5515",
			expected_value: 3,
		},
		{
			input_line:     "Q2Q2Q",
			expected_value: 4,
		},
		{
			input_line:     "1AAAA",
			expected_value: 5,
		},
		{
			input_line:     "AAAAA",
			expected_value: 6,
		},
		// For Step 2: Jokers 'J' can assume any card values
		{
			input_line:     "QJQJQ",
			expected_value: 6,
		},
		{
			input_line:     "JAAAA",
			expected_value: 6,
		},
		{
			input_line:     "T3T3J",
			expected_value: 4,
		},
		{
			input_line:     "2345J",
			expected_value: 1,
		},
		{
			input_line:     "JJJJJ",
			expected_value: 6,
		},
	}

	for _, test := range tests {
		hand_type := GetHandType(test.input_line)
		if hand_type != test.expected_value {
			t.Fatalf("Expected %d value. actual %d for line %s", test.expected_value, hand_type, test.input_line)
		}
	}
}
