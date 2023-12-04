package main

import "testing"

func TestGetNumbersFromLine(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value Scratchcard
	}{
		{
			input_line: " 41 48 83 86 17 | 83 86 6 31 17 9 48 53",
			expected_value: Scratchcard{
				winning_numbers: []int{41, 48, 83, 86, 17},
				user_numbers:    []int{83, 86, 6, 31, 17, 9, 48, 53},
			},
		},
		{
			input_line: " 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			expected_value: Scratchcard{
				winning_numbers: []int{13, 32, 20, 16, 61},
				user_numbers:    []int{61, 30, 68, 82, 17, 32, 24, 19},
			},
		},
		{
			input_line: " 1 21 53 59 44 | 69 82 63 72 16 21 14 1",
			expected_value: Scratchcard{
				winning_numbers: []int{1, 21, 53, 59, 44},
				user_numbers:    []int{69, 82, 63, 72, 16, 21, 14, 1},
			},
		},
	}

	for _, test := range tests {
		scratchcard := GetScratchcardFromLine(test.input_line)
		for _, n := range test.expected_value.winning_numbers {
			if SliceIndexOf(n, scratchcard.winning_numbers) == -1 {
				t.Fatalf("Winning numbers not correct. Expected %v value. actual %v for line %s", test.expected_value, scratchcard.winning_numbers, test.input_line)
			}
		}
		for _, n := range test.expected_value.user_numbers {
			if SliceIndexOf(n, scratchcard.user_numbers) == -1 {
				t.Fatalf("User numbers not correct. Expected %v value. actual %v for line %s", test.expected_value, scratchcard.user_numbers, test.input_line)
			}
		}
	}
}

func TestCalcScratchCardPoints(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value int
	}{
		{
			input_line:     " 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expected_value: 8,
		},
		{
			input_line:     " 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			expected_value: 2,
		},
		{
			input_line:     "  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expected_value: 2,
		},
		{
			input_line:     " 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			expected_value: 1,
		},
		{
			input_line:     " 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			expected_value: 0,
		},
		{
			input_line:     " 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			expected_value: 0,
		},
	}

	for _, test := range tests {
		scratchcard := GetScratchcardFromLine(test.input_line)
		points := CalcScratchCardPoints(scratchcard)
		if points != test.expected_value {
			t.Fatalf("Expected %d value. actual %d for line %s", test.expected_value, points, test.input_line)
		}
	}
}
