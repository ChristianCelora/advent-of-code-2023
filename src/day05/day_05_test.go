package main

import "testing"

func TestReadSeeds(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value []int
	}{
		{
			input_line:     "seeds: 79 14 55 13",
			expected_value: []int{79, 14, 55, 13},
		},
		{
			input_line:     "seeds: 1 2 3 4",
			expected_value: []int{1, 2, 3, 4},
		},
		{
			input_line:     "seeds: 1132132257 323430997 2043754183 4501055 2539071613 1059028389 1695770806 60470169 2220296232 251415938 1673679740 6063698 962820135 133182317 262615889 327780505 3602765034 194858721 2147281339 37466509",
			expected_value: []int{1132132257, 323430997, 2043754183, 4501055, 2539071613, 1059028389, 1695770806, 60470169, 2220296232, 251415938, 1673679740, 6063698, 962820135, 133182317, 262615889, 327780505, 3602765034, 194858721, 2147281339, 37466509},
		},
	}

	for _, test := range tests {
		seeds := ReadSeeds(test.input_line)
		for i, expected_v := range test.expected_value {
			if expected_v != seeds[i] {
				t.Fatalf("seed not correct. Expected %v value. actual %v for line %s", test.expected_value, seeds, test.input_line)
			}
		}
	}
}

func TestAlmanacMapAddMapValues(t *testing.T) {
	tests := []struct {
		input_lines    []string
		expected_value map[int]int
	}{
		{
			input_lines: []string{
				"50 98 2",
				"52 50 10",
			},
			expected_value: map[int]int{
				98: 50,
				99: 51,
				50: 52,
				51: 53,
				52: 54,
				53: 55,
				54: 56,
				55: 57,
				56: 58,
				57: 59,
				58: 60,
				59: 61,
			},
		},
		{
			input_lines: []string{
				"0 15 37",
				"37 52 2",
				"39 0 15",
			},
			expected_value: map[int]int{
				0:  39,
				1:  40,
				2:  41,
				3:  42,
				4:  43,
				5:  44,
				6:  45,
				7:  46,
				8:  47,
				9:  48,
				10: 49,
				11: 50,
				12: 51,
				13: 52,
				14: 53,
				15: 0,
				16: 1,
				17: 2,
				18: 3,
				19: 4,
				20: 5,
				21: 6,
				22: 7,
				23: 8,
				24: 9,
				25: 10,
				26: 11,
				27: 12,
				28: 13,
				29: 14,
				30: 15,
				31: 16,
				32: 17,
				33: 18,
				34: 19,
				35: 20,
				36: 21,
				37: 22,
				38: 23,
				39: 24,
				40: 25,
				41: 26,
				42: 27,
				43: 28,
				44: 29,
				45: 30,
				46: 31,
				47: 32,
				48: 33,
				49: 34,
				50: 35,
				51: 36,
				52: 37,
				53: 38,
			},
		},
	}

	for _, test := range tests {
		almanac_map := AlmanacMap{
			name:   "test",
			values: make(map[int]int),
		}
		for _, line := range test.input_lines {
			almanac_map.addMapValues(line)
		}
		for key, expected_v := range test.expected_value {
			if expected_v != almanac_map.values[key] {
				t.Fatalf("map not correct for key %d. Expected %v map. actual %v for line %v", key, test.expected_value, almanac_map.values, test.input_lines)
			}
		}
	}
}
