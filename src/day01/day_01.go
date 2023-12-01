package main

import (
	"os"
	"fmt"
	"bufio"
	"unicode"
	"strconv"
	"strings"
	"sort"
)

type Occurrence struct {
	index int
	value int
	label string
}

func ConvertSpelledDigits(line string) string {
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

	for i, digit := range spelledDigits {
		index := strings.Index(line, digit)
		if index >= 0 {
			occurrences = append(
				occurrences,
				Occurrence{value: i+1, label: digit, index: index},
			)
		}
	}

	sort.Slice(occurrences, func(i, j int) bool {
		return occurrences[i].index < occurrences[j].index
	})

	for _, oc := range occurrences {
		line = strings.Replace(line, oc.label, fmt.Sprint(oc.value), -1)
	}

	return line
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

func ReadLines(path string) []string {
	var lines []string
	file, err := os.Open(path)
    if err != nil {
        fmt.Printf("error in buffer: %s", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
    }

	return lines
}

func main() {
	var sum int
	lines := ReadLines("./day01/data/input2_2.txt")
	for _, line := range lines {
		convertedLine := ConvertSpelledDigits(string(line))
		sum += GetCalibration(convertedLine)
	}

	fmt.Printf("Calibration sum is %d\n", sum)
}