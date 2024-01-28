package main

import "testing"

func TestCountStraightLine(t *testing.T) {
	tests := []struct {
		input_cells     []Cell
		expected_result int
	}{
		{
			input_cells: []Cell{
				{x: 0, y: 0},
				{x: 0, y: 1},
				{x: 0, y: 2},
			},
			expected_result: 3,
		},
		{
			input_cells: []Cell{
				{x: 0, y: 1},
				{x: 0, y: 0},
				{x: 1, y: 0},
				{x: 2, y: 0},
				{x: 3, y: 0},
			},
			expected_result: 4,
		},
		{
			input_cells: []Cell{
				{x: 0, y: 1},
			},
			expected_result: 1,
		},
		{
			input_cells: []Cell{
				{x: 0, y: 1},
				{x: 1, y: 1},
			},
			expected_result: 2,
		},
	}

	for _, test := range tests {
		// set prev pointer
		for i := range test.input_cells {
			if i > 0 {
				test.input_cells[i].prev = &test.input_cells[i-1]
			}
		}

		last_cell := test.input_cells[len(test.input_cells)-1]
		count := CountStraightLine(last_cell)
		if count != test.expected_result {
			t.Fatalf("Expected %d. actual %d for input %v", test.expected_result, count, test.input_cells)
		}
	}
}
