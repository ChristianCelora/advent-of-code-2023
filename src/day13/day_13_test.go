package main

import "testing"

func TestFindReflectionPoint(t *testing.T) {
	tests := []struct {
		input_line       []string
		expected_result  int
		expected_result2 int
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
			expected_result:  5,
			expected_result2: 300,
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
			expected_result:  400,
			expected_result2: 100,
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
			expected_result:  1,
			expected_result2: 5,
		},
		// {
		// 	input_line: []string{
		// 		".#.##..##",
		// 		"...#..#..",
		// 		"###..#.##",
		// 		"##...#.##",
		// 		"##.####..",
		// 		".#..##.#.",
		// 		".#..##.#.",
		// 		"##..###..",
		// 		"##...#.##",
		// 		"##...#.##",
		// 		"##..###..",
		// 		".#..##.#.",
		// 		".#..##.#.",
		// 		"##.####..",
		// 		"##...#.##",
		// 	},
		// 	expected_result: 900,
		// },
	}

	for _, test := range tests {
		reflection_point, reflection_point2 := FindReflectionPoint(test.input_line)
		if reflection_point != test.expected_result {
			t.Fatalf("Expected %d. actual %d for input %v", test.expected_result, reflection_point, test.input_line)
		}
		if reflection_point2 != test.expected_result2 {
			t.Fatalf("Expected %d. actual %d for input %v", test.expected_result2, reflection_point2, test.input_line)
		}
	}
}
