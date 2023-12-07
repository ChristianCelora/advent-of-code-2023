package main

import "testing"

func TestCalcDistance(t *testing.T) {
	tests := []struct {
		input_total_time       int
		input_btn_pressed_time int
		expected_value         int
	}{
		{
			input_total_time:       7,
			input_btn_pressed_time: 0,
			expected_value:         0,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 1,
			expected_value:         6,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 2,
			expected_value:         10,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 3,
			expected_value:         12,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 4,
			expected_value:         12,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 5,
			expected_value:         10,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 6,
			expected_value:         6,
		},
		{
			input_total_time:       7,
			input_btn_pressed_time: 7,
			expected_value:         0,
		},
	}

	for _, test := range tests {
		distance := CalcDistance(test.input_total_time, test.input_btn_pressed_time)
		if distance != test.expected_value {
			t.Fatalf("Expected %d value. actual %d for time (total: %d, pressed: %d)", test.expected_value, distance, test.input_total_time, test.input_btn_pressed_time)
		}
	}
}
