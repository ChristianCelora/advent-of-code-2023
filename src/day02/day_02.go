package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_BLUE  = 14
	MAX_GREEN = 13
	MAX_RED   = 12
)

func ReadGameID(line string) int {
	first_space := strings.Index(line, " ")
	first_colon := strings.Index(line, ":")

	if first_space == -1 || first_colon == -1 {
		panic("error reading game ID")
	}

	id, err := strconv.Atoi(line[first_space+1 : first_colon])
	if err != nil {
		panic(err)
	}

	return id
}

type Set struct {
	blue  int
	green int
	red   int
	other int
}

func ReadCubeSets(line string) []Set {
	var res []Set
	first_colon := strings.Index(line, ":")
	rounds := strings.Split(line[first_colon+1:], ";")

	for _, r := range rounds {
		set := Set{}
		r := strings.TrimSpace(r)
		cubes := strings.Split(r, ",")
		for _, cube := range cubes {
			cube = strings.TrimSpace(cube)
			c := strings.Split(cube, " ")
			n, err := strconv.Atoi(c[0])
			if err != nil {
				panic(err)
			}
			color := c[1]
			switch color {
			case "blue":
				set.blue = n
			case "green":
				set.green = n
			case "red":
				set.red = n
			default:
				set.other = n
			}
		}
		res = append(res, set)
	}

	return res
}

func IsGameValid(game_set []Set) bool {
	for _, set := range game_set {
		if set.blue > MAX_BLUE ||
			set.green > MAX_GREEN ||
			set.red > MAX_RED ||
			set.other > 0 {
			return false
		}
	}

	return true
}

func GetCubesForValidGame(game_set []Set) Set {
	var max_cubes Set
	for _, set := range game_set {
		if set.blue > max_cubes.blue {
			max_cubes.blue = set.blue
		}
		if set.green > max_cubes.green {
			max_cubes.green = set.green
		}
		if set.red > max_cubes.red {
			max_cubes.red = set.red
		}
	}

	return max_cubes
}

func main() {
	var sum int
	var sum_power int
	// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	lines := reader.ReadLines("./day02/data/input1_2.txt")
	for _, line := range lines {
		game_id := ReadGameID(line)
		cube_sets := ReadCubeSets(line)
		// star 1
		if IsGameValid(cube_sets) {
			sum += game_id
		}
		// star 2
		cubes_valid_game := GetCubesForValidGame(cube_sets)
		power_set := cubes_valid_game.blue * cubes_valid_game.green * cubes_valid_game.red
		sum_power += power_set
	}

	fmt.Printf("Game valid ID sum is %d\n", sum)
	fmt.Printf("Game power set sum is %d\n", sum_power)
}
