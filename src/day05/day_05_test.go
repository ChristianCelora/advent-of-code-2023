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

func TestReadMultipleSeeds(t *testing.T) {
	tests := []struct {
		input_line     string
		expected_value []SeedsRange
	}{
		{
			input_line: "seeds: 79 14 55 13",
			expected_value: []SeedsRange{
				{
					from: 79,
					to:   92,
				},
				{
					from: 55,
					to:   67,
				},
			},
		},
		{
			input_line: "seeds: 1 2 3 4",
			expected_value: []SeedsRange{
				{
					from: 1,
					to:   2,
				},
				{
					from: 3,
					to:   6,
				},
			},
		},
	}

	for _, test := range tests {
		seeds := ReadMultipleSeeds(test.input_line)
		// t.Logf("expected %v, actual %v\n", test.expected_value, seeds)
		for i, expected_v := range test.expected_value {
			if expected_v.from != seeds[i].from ||
				expected_v.to != seeds[i].to {
				t.Fatalf("seed not correct. Expected %v value. actual %v for line %s", test.expected_value, seeds, test.input_line)
			}
		}
	}
}

func TestAlmanacMapAddMapValues(t *testing.T) {
	tests := []struct {
		input_lines    []string
		expected_value []MapRange
	}{
		{
			input_lines: []string{
				"50 98 2",
				"52 50 48",
			},
			expected_value: []MapRange{
				{
					input_from:  98,
					input_to:    99,
					output_from: 50,
					output_to:   51,
				},
				{
					input_from:  50,
					input_to:    97,
					output_from: 52,
					output_to:   99,
				},
			},
		},
		{
			input_lines: []string{
				"0 15 37",
				"37 52 2",
				"39 0 15",
			},
			expected_value: []MapRange{
				{
					input_from:  15,
					input_to:    51,
					output_from: 0,
					output_to:   36,
				},
				{
					input_from:  52,
					input_to:    53,
					output_from: 37,
					output_to:   38,
				},
				{
					input_from:  0,
					input_to:    14,
					output_from: 39,
					output_to:   53,
				},
			},
		},
	}

	for _, test := range tests {
		almanac_map := AlmanacMap{
			name: "test",
		}
		for _, line := range test.input_lines {
			almanac_map.addMapValues(line)
		}
		for key, expected_v := range test.expected_value {
			if expected_v.input_from != almanac_map.values[key].input_from ||
				expected_v.input_to != almanac_map.values[key].input_to ||
				expected_v.output_from != almanac_map.values[key].output_from ||
				expected_v.output_to != almanac_map.values[key].output_to {
				t.Fatalf("map not correct for key %d. Expected %v map. actual %v for line %v", key, test.expected_value, almanac_map.values, test.input_lines)
			}
		}
	}
}

func TestMapRangeTransform(t *testing.T) {
	tests := []struct {
		input_map      MapRange
		input_value    int
		expected_value int
	}{
		{
			input_map: MapRange{
				input_from:  50,
				input_to:    97,
				output_from: 52,
				output_to:   99,
			},
			input_value:    14,
			expected_value: 14,
		},
		{
			input_map: MapRange{
				input_from:  50,
				input_to:    51,
				output_from: 87,
				output_to:   88,
			},
			input_value:    99,
			expected_value: 99,
		},
		{
			input_map: MapRange{
				input_from:  50,
				input_to:    51,
				output_from: 98,
				output_to:   99,
			},
			input_value:    50,
			expected_value: 98,
		},
		{
			input_map: MapRange{
				input_from:  50,
				input_to:    51,
				output_from: 98,
				output_to:   99,
			},
			input_value:    51,
			expected_value: 99,
		},
		{
			input_map: MapRange{
				input_from:  50,
				input_to:    97,
				output_from: 52,
				output_to:   99,
			},
			input_value:    60,
			expected_value: 62,
		},
		{
			input_map: MapRange{
				input_from:  53,
				input_to:    61,
				output_from: 49,
				output_to:   57,
			},
			input_value:    53,
			expected_value: 49,
		},
		{
			input_map: MapRange{
				input_from:  53,
				input_to:    61,
				output_from: 49,
				output_to:   57,
			},
			input_value:    61,
			expected_value: 57,
		},
		{
			input_map: MapRange{
				input_from:  53,
				input_to:    61,
				output_from: 49,
				output_to:   57,
			},
			input_value:    60,
			expected_value: 56,
		},
	}

	for _, test := range tests {
		converted_value := test.input_map.transform(test.input_value)
		if converted_value != test.expected_value {
			t.Fatalf("seed not correct. Expected %d value. actual %d for input %d", test.expected_value, converted_value, test.input_value)
		}
	}
}
