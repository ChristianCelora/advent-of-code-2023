package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
	"strings"
)

func CountCombinaitons(spring_line string, spring_map []int) int {
	var num_combinations int
	var combinations []string
	var valid_combinations map[string]bool

	valid_combinations = make(map[string]bool)

	combinations = CreateCombinations(spring_line)
	fmt.Printf("total combinations for line %s are %d\n", spring_line, len(combinations))
	for _, combination := range combinations {
		is_valid, ok := valid_combinations[combination]
		if !ok {
			is_valid = IsCombinationValid(combination, spring_map)
			valid_combinations[combination] = is_valid
		}
		if is_valid {
			num_combinations++
		}
	}

	return num_combinations
}

func CreateCombinations(spring_line string) []string {
	var combinations []string

	idx := strings.Index(spring_line, "?")
	if idx == -1 {
		return []string{spring_line}
	}

	if spring_line[idx] == '?' {
		new_line1 := strings.Clone(spring_line)
		new_line1 = replaceAtIndex(new_line1, '.', idx)
		combinations = append(combinations, CreateCombinations(new_line1)...)

		new_line2 := strings.Clone(spring_line)
		new_line2 = replaceAtIndex(new_line2, '#', idx)
		combinations = append(combinations, CreateCombinations(new_line2)...)
	}

	return combinations
}

func IsCombinationValid(combination string, spring_map []int) bool {
	var is_valid bool
	var idx int
	var groups_final []string

	groups := strings.Split(combination, ".")
	for i := 0; i < len(groups); i++ {
		if len(groups[i]) != 0 {
			groups_final = append(groups_final, groups[i])
		}
	}

	// fmt.Printf("groups_final %v, spring_map %v\n", groups_final, spring_map)
	// to be valid the both maps should have the same group length
	if len(groups_final) != len(spring_map) {
		return false
	}

	is_valid = true
	for i := 0; i < len(groups_final); i++ {
		if len(groups_final[i]) != spring_map[idx] {
			//fmt.Printf("idx %d, spring_map %d, len_groups %d\n", idx, spring_map[idx], len(groups_final))
			is_valid = false
			break
		}

		idx++
	}

	if idx != len(spring_map) {
		is_valid = false
	}

	return is_valid
}

func replaceAtIndex(in string, b byte, i int) string {
	out := []byte(in)
	out[i] = b
	return string(out)
}

func getSpringsMapNumbers(line string, times_folded int) []int {
	var num_map []int
	var num_map_unfolded []int

	for _, num := range strings.Split(line, ",") {
		n, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		num_map = append(num_map, n)
	}

	// unfold map
	for i := 0; i < times_folded; i++ {
		num_map_unfolded = append(num_map_unfolded, num_map...)
	}

	return num_map_unfolded
}

func UnfoldSpringsDamagedMap(line string, times_folded int) string {
	var final_line string
	for i := 0; i < times_folded; i++ {
		final_line += line + "?"
	}

	return final_line[:len(final_line)-1]
}

func main() {
	var sum_combinations int
	var spring_num_map []int

	lines := reader.ReadLines("./day12/data/input1_1.txt")
	for _, line := range lines {
		info := strings.Split(line, " ")
		if len(info) != 2 {
			panic("err split line " + line)
		}
		spring_damaged_map := info[0]
		// step 1
		// spring_num_map = getSpringsMapNumbers(info[1])

		// step 2
		spring_num_map = getSpringsMapNumbers(info[1], 5)
		spring_damaged_map = UnfoldSpringsDamagedMap(spring_damaged_map, 5)
		fmt.Printf("unfolded dmg map %s, unfolded num map %v\n", spring_damaged_map, spring_num_map)

		n_combinations := CountCombinaitons(spring_damaged_map, spring_num_map)
		fmt.Printf("Valid combinations for %s are %d\n", spring_damaged_map, n_combinations)
		sum_combinations += n_combinations
	}

	fmt.Printf("The sum of possible combination is %d", sum_combinations)
}
