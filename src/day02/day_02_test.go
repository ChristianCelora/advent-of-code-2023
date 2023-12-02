package main

import "testing"

func TestReadGameID(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value int
	}{
		{
			input_line:     "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected_value: 1,
		},
		{
			input_line:     "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected_value: 2,
		},
		{
			input_line:     "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected_value: 3,
		},
	}

	for _, test := range tests {
		calibration := ReadGameID(test.input_line)
		if calibration != test.expected_value {
			t.Fatalf("expected %d value. actual %d for line %s", test.expected_value, calibration, test.input_line)
		}
	}
}
