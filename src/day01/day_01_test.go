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

func TestConvertSpelledDigits(t *testing.T) {
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
			expected_line: "8wo3",
		},
		{
			input_line: "abcone2threexyz",
			expected_line: "abc123xyz",
		},
		{
			input_line: "xtwone3four",
			expected_line: "x2ne34",
		},
	}

	for _, test := range tests {
		line := ConvertSpelledDigits(test.input_line)
		if line != test.expected_line {
			t.Fatalf("expected %s value. actual %s for line %s", test.expected_line, line, test.input_line)
		}
	}
}