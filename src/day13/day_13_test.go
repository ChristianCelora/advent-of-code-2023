package main

import "testing"

func TestFindReflectionPoint(t *testing.T) {
	tests := []struct {
		input_line      []string
		expected_result int
	}{
		{
			input_line: []string{
				"#.##..##.",
				"..#.##.#.",
				"##......#",
				"##......#",
				"..#.##.#.",
				"..##..##.",
				"#.#.##.#.",
			},
			expected_result: 5,
		},
		{
			input_line: []string{
				"#...##..#",
				"#....#..#",
				"..##..###",
				"#####.##.",
				"#####.##.",
				"..##..###",
				"#....#..#",
			},
			expected_result: 400,
		},
		{
			input_line: []string{
				"#####.###",
				"..######.",
				"####..###",
				"####..###",
				"##......#",
				"###.##.##",
				"..######.",
			},
			expected_result: 1,
		},
		{
			input_line: []string{
				".#.##..##",
				"...#..#..",
				"###..#.##",
				"##...#.##",
				"##.####..",
				".#..##.#.",
				".#..##.#.",
				"##..###..",
				"##...#.##",
				"##...#.##",
				"##..###..",
				".#..##.#.",
				".#..##.#.",
				"##.####..",
				"##...#.##",
			},
			expected_result: 900,
		},
	}

	for _, test := range tests {
		reflection_point := FindReflectionPoint(test.input_line)
		if reflection_point != test.expected_result {
			t.Fatalf("Expected %d. actual %d for input %v", test.expected_result, reflection_point, test.input_line)
		}
	}
}
