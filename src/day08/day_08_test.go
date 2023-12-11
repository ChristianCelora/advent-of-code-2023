package main

import "testing"

func TestCalcStepsCombinationLCM(t *testing.T) {
	tests := []struct {
		input_steps_combination [][]int
		expected_value          []int
	}{
		{
			input_steps_combination: [][]int{
				{2},
				{3},
				{5, 6},
			},
			expected_value: []int{30, 6},
		},
		{
			input_steps_combination: [][]int{
				{2},
				{3, 5},
				{5, 6},
			},
			expected_value: []int{30, 10, 6, 30},
		},
	}

	for _, test := range tests {
		stepsCombinations := CalcStepsCombinationLCM(test.input_steps_combination)
		for i, v := range stepsCombinations {
			if v != test.expected_value[i] {
				t.Fatalf("Expected %v value. actual %v for input %v", test.expected_value, stepsCombinations, test.input_steps_combination)
			}
		}
	}
}
