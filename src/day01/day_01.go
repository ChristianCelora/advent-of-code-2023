package main

import (
	"adventcode/reader"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type Occurrence struct {
	index int
	value int
	label string
}

func ExtractDigits(line string) string {
	var occurrences []Occurrence
	spelledDigits := [9]string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for i, c := range line {
		if unicode.IsDigit(c) {
			occurrences = append(
				occurrences,
				Occurrence{value: int(c - '0'), label: "", index: i},
			)
		}
	}

	for i, digit := range spelledDigits {
		index := strings.Index(line, digit)
		if index >= 0 {
			occurrences = append(
				occurrences,
				Occurrence{value: i + 1, label: digit, index: index},
			)
		}
		lastIndex := strings.LastIndex(line, digit)
		if lastIndex >= 0 && lastIndex != index {
			occurrences = append(
				occurrences,
				Occurrence{value: i + 1, label: digit, index: lastIndex},
			)
		}
	}

	sort.Slice(occurrences, func(i, j int) bool {
		return occurrences[i].index < occurrences[j].index
	})

	var final_line string
	for _, oc := range occurrences {
		final_line += fmt.Sprint(oc.value)
	}

	return final_line
}

func GetCalibration(line string) int {
	var first rune
	var last rune
	for _, c := range line {
		if unicode.IsDigit(c) {
			if first == 0 {
				first = c
			} else {
				last = c
			}
		}
	}
	if last == 0 {
		last = first
	}
	res := string(first) + string(last)
	calibration, _ := strconv.Atoi(res)
	return calibration
}

func main() {
	var sum int
	lines := reader.ReadLines("./day01/data/input2_2.txt")
	for _, line := range lines {
		convertedLine := ExtractDigits(string(line))
		sum += GetCalibration(convertedLine) // just take first and last char of string
	}

	fmt.Printf("Calibration sum is %d\n", sum)
}
