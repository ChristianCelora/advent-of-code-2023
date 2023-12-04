package main

import "testing"

type Coordinates struct {
	x int
	y int
}

func TestExtractFullNumber(t *testing.T) {
	tests := []struct {
		input_line        string
		input_coordinates Coordinates
		expected_value    int
	}{
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 0},
			expected_value:    467,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 1},
			expected_value:    467,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 2},
			expected_value:    467,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 3},
			expected_value:    467,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 5},
			expected_value:    114,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 6},
			expected_value:    114,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 7},
			expected_value:    114,
		},
		{
			input_line:        "467..114..",
			input_coordinates: Coordinates{x: 0, y: 8},
			expected_value:    114,
		},
		{
			input_line:        "467....114",
			input_coordinates: Coordinates{x: 0, y: 6},
			expected_value:    114,
		},
	}

	for _, test := range tests {
		mat := CreateMatrix()
		InsertLineInMatrix(mat, test.input_line, 0)
		full_number := ExtractFullNumber(mat, test.input_coordinates.x, test.input_coordinates.y)
		if full_number != test.expected_value {
			t.Fatalf("expected %d value. actual %d for line %s", test.expected_value, full_number, test.input_line)
		}
	}
}
