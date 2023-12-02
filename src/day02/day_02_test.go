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

func TestReadCubeSets(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value []Set
	}{
		{
			input_line: "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected_value: []Set{
				{
					blue: 3,
					red:  4,
				},
				{
					blue:  6,
					green: 2,
					red:   1,
				},
				{
					green: 2,
				},
			},
		},
		{
			input_line: "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected_value: []Set{
				{
					blue:  1,
					red:   0,
					green: 2,
				},
				{
					blue:  4,
					red:   1,
					green: 3,
				},
				{
					blue:  1,
					red:   0,
					green: 1,
				},
			},
		},
		{
			input_line: "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected_value: []Set{
				{
					blue:  6,
					red:   20,
					green: 8,
				},
				{
					blue:  5,
					red:   4,
					green: 13,
				},
				{
					blue:  0,
					red:   1,
					green: 5,
				},
			},
		},
		{
			input_line: "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected_value: []Set{
				{
					blue:  6,
					red:   3,
					green: 1,
				},
				{
					blue:  0,
					red:   6,
					green: 3,
				},
				{
					blue:  15,
					red:   14,
					green: 3,
				},
			},
		},
		{
			input_line: "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected_value: []Set{
				{
					blue:  1,
					red:   6,
					green: 3,
				},
				{
					blue:  2,
					red:   1,
					green: 2,
				},
			},
		},
		{
			input_line: "Game 5: 6 red, 1 blue, 3 green, 1 yellow; 2 blue, 1 red, 2 green",
			expected_value: []Set{
				{
					blue:  1,
					red:   6,
					green: 3,
					other: 1,
				},
				{
					blue:  2,
					red:   1,
					green: 2,
				},
			},
		},
	}

	for _, test := range tests {
		sets := ReadCubeSets(test.input_line)
		for i, set := range sets {
			if set.blue != test.expected_value[i].blue ||
				set.green != test.expected_value[i].green ||
				set.red != test.expected_value[i].red {
				t.Fatalf("expected %v value. actual %v for line %s", test.expected_value, sets, test.input_line)
			}
		}
	}
}

func TestIsGameValid(t *testing.T) {
	tests := []struct {
		input_set      []Set
		expected_value bool
	}{
		{
			input_set: []Set{
				{
					blue: 3,
					red:  4,
				},
				{
					blue:  6,
					green: 2,
					red:   1,
				},
				{
					green: 2,
				},
			},
			expected_value: true,
		},
		{
			input_set: []Set{
				{
					blue: 20,
					red:  4,
				},
			},
			expected_value: false,
		},
		{
			input_set: []Set{
				{
					blue:  3,
					red:   4,
					other: 1,
				},
			},
			expected_value: false,
		},
	}

	for _, test := range tests {
		is_valid := IsGameValid(test.input_set)
		if is_valid != test.expected_value {
			t.Fatalf("expected %t value. actual %t for set %v", test.expected_value, is_valid, test.input_set)
		}
	}
}

func TestGetCubesForValidGame(t *testing.T) {
	tests := []struct {
		input_set      []Set
		expected_value Set
	}{
		{
			input_set: []Set{
				{
					blue: 3,
					red:  4,
				},
				{
					blue:  6,
					green: 2,
					red:   1,
				},
				{
					green: 2,
				},
			},
			expected_value: Set{
				blue:  6,
				red:   4,
				green: 2,
			},
		},
		{
			input_set: []Set{
				{
					blue: 20,
					red:  4,
				},
				{
					blue: 12,
					red:  15,
				},
			},
			expected_value: Set{
				blue:  20,
				red:   15,
				green: 0,
			},
		},
	}

	for _, test := range tests {
		cubes := GetCubesForValidGame(test.input_set)
		if cubes.blue != test.expected_value.blue ||
			cubes.green != test.expected_value.green ||
			cubes.red != test.expected_value.red {
			t.Fatalf("expected %v value. actual %v for set %v", test.expected_value, cubes, test.input_set)
		}
	}
}
