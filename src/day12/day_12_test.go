package main

import "testing"

func TestCreateCombinations(t *testing.T) {
	tests := []struct {
		input_line                string
		expected_combinations_num int
	}{
		{
			input_line:                "?.###",
			expected_combinations_num: 2,
		},
		{
			input_line:                "??.###",
			expected_combinations_num: 4,
		},
		{
			input_line:                "???.###",
			expected_combinations_num: 8,
		},
		{
			input_line:                "?###????????",
			expected_combinations_num: 512,
		},
	}

	for _, test := range tests {
		combinations := CreateCombinations(test.input_line)
		if len(combinations) != test.expected_combinations_num {
			t.Fatalf("Expected %d combinations. actual %d for input %s", test.expected_combinations_num, len(combinations), test.input_line)
		}
	}
}

func TestIsCombinationValids(t *testing.T) {
	tests := []struct {
		input_line     string
		spring_map     []int
		expected_value bool
	}{
		{
			input_line:     "#.#.###",
			spring_map:     []int{1, 1, 3},
			expected_value: true,
		},
		{
			input_line:     "#.#.",
			spring_map:     []int{1, 1, 3},
			expected_value: false,
		},
		{
			input_line:     "#",
			spring_map:     []int{1, 1, 3},
			expected_value: false,
		},
		{
			input_line:     "...#.##...",
			spring_map:     []int{1, 1, 3},
			expected_value: false,
		},
		{
			input_line:     "...#.##..###",
			spring_map:     []int{1, 1, 3},
			expected_value: false,
		},
		{
			input_line:     ".###..##.#.#",
			spring_map:     []int{3, 2, 1},
			expected_value: false,
		},
		{
			input_line:     ".###..##.#.",
			spring_map:     []int{3, 2, 1},
			expected_value: true,
		},
	}

	for _, test := range tests {
		is_valid := IsCombinationValid(test.input_line, test.spring_map)
		if is_valid != test.expected_value {
			t.Fatalf("Expected %t. actual %t for input %s", test.expected_value, is_valid, test.input_line)
		}
	}
}

func TestUnfoldSpringsDamagedMap(t *testing.T) {
	tests := []struct {
		input_line      string
		unfolded_times  int
		expected_result string
	}{
		{
			input_line:      "?.###",
			unfolded_times:  5,
			expected_result: "?.###??.###??.###??.###??.###",
		},
	}

	for _, test := range tests {
		unfolded_line := UnfoldSpringsDamagedMap(test.input_line, test.unfolded_times)
		if unfolded_line != test.expected_result {
			t.Fatalf("Expected %s. actual %s for input %s", test.expected_result, unfolded_line, test.input_line)
		}
	}
}
