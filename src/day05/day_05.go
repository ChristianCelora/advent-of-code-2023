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

type AlmanacMap struct {
	name   string
	values map[int]int
}

func (almanac_map *AlmanacMap) addMapValues(line string) {
	map_instructions := SplitSpace(line)
	if len(map_instructions) != 3 {
		panic("wrong read map_instructions line")
	}
	var input, output int
	for i := 0; i < map_instructions[2]; i++ {
		input = map_instructions[1] + i
		output = map_instructions[0] + i
		almanac_map.values[input] = output
	}
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
				name:   almanac_map_name,
				values: make(map[int]int),
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
			output, ok := almanac_map.values[seed_transformed]
			fmt.Printf("map %s, input: %d, output: %d\n", almanac_map.name, seed_transformed, output)
			if ok {
				seed_transformed = output
			}
		}
		if seed_transformed < min_location {
			min_location = seed_transformed
		}
	}

	fmt.Printf("the minimum location is %d\n", min_location)
}
