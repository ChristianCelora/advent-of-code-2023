package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ReadSeeds(line string) []int {
	split_line := strings.Split(line, ":")
	if len(split_line) != 2 {
		panic("wrong read line")
	}
	return SplitSpace(split_line[1])
}

type MapRange struct {
	input_from  int
	input_to    int
	output_from int
	output_to   int
}

func (m *MapRange) transform(input int) int {
	var output int

	if input < m.input_from || input > m.input_to {
		return input
	}

	delta := int(math.Abs(float64(m.input_from - m.output_from)))
	if m.input_from > m.output_from {
		output = input - delta
	} else {
		output = input + delta
	}

	return output
}

type AlmanacMap struct {
	name   string
	values []MapRange
}

func (almanac_map *AlmanacMap) addMapValues(line string) {
	map_instructions := SplitSpace(line)
	if len(map_instructions) != 3 {
		panic("wrong read map_instructions line")
	}
	/**
	var input, output int
	for i := 0; i < map_instructions[2]; i++ {
		input = map_instructions[1] + i
		output = map_instructions[0] + i
		almanac_map.values[input] = output
	}
	*/
	map_range := MapRange{
		input_from:  map_instructions[1],
		input_to:    map_instructions[1] + map_instructions[2] - 1,
		output_from: map_instructions[0],
		output_to:   map_instructions[0] + map_instructions[2] - 1,
	}
	almanac_map.values = append(almanac_map.values, map_range)
}

func SplitSpace(line string) []int {
	var res []int
	seeds := strings.Split(line, " ")
	for _, seed := range seeds {
		if seed == "" {
			continue
		}
		s, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		res = append(res, s)
	}
	return res
}

func main() {
	var almanac_maps []AlmanacMap
	var almanac_map_name string
	var min_location int
	lines := reader.ReadLines("./day05/data/input1_2.txt")
	seeds := ReadSeeds(lines[0])

	for _, line := range lines[1:] {
		if strings.TrimSpace(line) == "" {
			continue
		}

		if strings.Index(line, "map") != -1 {
			map_split := strings.Split(line, " ")
			if len(map_split) != 2 {
				panic("wrong read line")
			}
			almanac_map_name = map_split[0]
			almanac := AlmanacMap{
				name: almanac_map_name,
			}
			almanac_maps = append(almanac_maps, almanac)
		} else {
			almanac_maps[len(almanac_maps)-1].addMapValues(line)
		}
	}

	min_location = math.MaxInt32
	for _, seed := range seeds {
		fmt.Printf("\n\n seed %d\n", seed)

		seed_transformed := seed
		for i := 0; i < len(almanac_maps); i++ {
			almanac_map := almanac_maps[i]
			for j := 0; j < len(almanac_map.values); j++ {
				map_range := almanac_map.values[j]
				if seed_transformed >= map_range.input_from && seed_transformed <= map_range.input_to {
					output := map_range.transform(seed_transformed)
					fmt.Printf("map %s, input: %d, output: %d\n", almanac_map.name, seed_transformed, output)
					seed_transformed = output
					break
				}
			}
		}

		if seed_transformed < min_location {
			min_location = seed_transformed
		}
	}

	fmt.Printf("the minimum location is %d\n", min_location)
}
