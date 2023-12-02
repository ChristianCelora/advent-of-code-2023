package main

import "testing"

func TestGetCalibration(t *testing.T) {
	tests := []struct {
		input_line string
		expected_value int
	} {
		{
			input_line: "1abc2",
			expected_value: 12,
		},
		{
			input_line: "pqr3stu8vwx",
			expected_value: 38,
		},
		{
			input_line: "a1b2c3d4e5f",
			expected_value: 15,
		},
		{
			input_line: "treb7uchet",
			expected_value: 77,
		},
		{
			input_line: "d0n7",
			expected_value: 7,
		},
	}

	for _, test := range tests {
		calibration := GetCalibration(test.input_line)
		if calibration != test.expected_value {
			t.Fatalf("expected %d value. actual %d for line %s", test.expected_value, calibration, test.input_line)
		}
	}
}

func TestExtractDigits(t *testing.T) {
	tests := []struct {
		input_line string
		expected_line string
	} {
		{
			input_line: "two1nine",
			expected_line: "219",
		},
		{
			input_line: "eightwothree",
			expected_line: "823",
		},
		{
			input_line: "abcone2threexyz",
			expected_line: "123",
		},
		{
			input_line: "xtwone3four",
			expected_line: "2134",
		},
		{
			input_line: "4nineeightseven2",
			expected_line: "49872",
		},
		{
			input_line: "zoneight234",
			expected_line: "18234",
		},
		{
			input_line: "7pqrstsixteen",
			expected_line: "76",
		},
		{
			input_line: "eighthree",
			expected_line: "83",
		},
		{
			input_line: "sevenine",
			expected_line: "79",
		},
		{
			input_line: "5fivezgfgcxbf3five",
			expected_line: "5535",
		},
	}

	for _, test := range tests {
		line := ExtractDigits(test.input_line)
		if line != test.expected_line {
			t.Fatalf("expected %s value. actual %s for line %s", test.expected_line, line, test.input_line)
		}
	}
}