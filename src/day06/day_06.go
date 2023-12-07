package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
	"strings"
)

type Race struct {
	time_btn_press int
	distance       int
}

func GetRacePossibilities(time int) []Race {
	var races []Race

	for i := 0; i <= time; i++ { // can be halfed - it is a normal distribution
		race := Race{
			time_btn_press: i,
			distance:       CalcDistance(time, i),
		}
		races = append(races, race)
	}

	return races
}

func CalcDistance(total_time int, time_btn_pressed int) int {
	moving_time := total_time - time_btn_pressed
	return moving_time * time_btn_pressed
}

func ReadRaceTimes(line string) []int {
	split_line := strings.Split(line, ":")
	if len(split_line) != 2 {
		panic("wrong read line")
	}
	return SplitSpace(split_line[1])
}

func ReadRaceDistanceRecord(line string) []int {
	split_line := strings.Split(line, ":")
	if len(split_line) != 2 {
		panic("wrong read line")
	}
	return SplitSpace(split_line[1])
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
	lines := reader.ReadLines("./day06/data/input2_2.txt")
	race_times := ReadRaceTimes(lines[0])
	race_distance_record := ReadRaceDistanceRecord(lines[1])
	n_possibilities := 1

	for i := 0; i < len(race_times); i++ {
		races := GetRacePossibilities(race_times[i])
		n_races_record_beaten := 0
		for _, race := range races {
			if race.distance > race_distance_record[i] {
				n_races_record_beaten++
			}
		}
		n_possibilities = n_possibilities * n_races_record_beaten
	}

	fmt.Printf("There are %d number of ways to beat a race record", n_possibilities)
}
