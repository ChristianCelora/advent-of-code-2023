package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
	"strings"
)

func GetNextSubsequence(sequence []int) []int {
	var subsequence []int
	subsequence = make([]int, len(sequence)-1)

	if len(sequence) == 1 {
		panic("error")
	}

	for i := 1; i < len(sequence); i++ {
		delta := sequence[i] - sequence[i-1]
		subsequence[i-1] = delta
	}

	return subsequence
}

func areAllElementsZero(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}

func GetSequenceFromLine(line string) []int {
	var sequence []int

	for _, v := range strings.Split(line, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		sequence = append(sequence, num)
	}

	return sequence
}

func main() {
	var sequence, next_sequence []int
	var sequence_history [][]int
	var total_elements int
	// var total_sum int
	lines := reader.ReadLines("./day09/data/input1_2.txt")

	for _, line := range lines {
		sequence = GetSequenceFromLine(line)
		fmt.Printf("sequence %v\n", sequence)
		sequence_history = make([][]int, 1)
		sequence_history[0] = sequence
		for true {
			next_sequence = GetNextSubsequence(sequence_history[len(sequence_history)-1])
			fmt.Printf("subsequence %v\n", next_sequence)
			sequence_history = append(sequence_history, next_sequence)
			if areAllElementsZero(sequence_history[len(sequence_history)-1]) {
				break
			}
		}

		fmt.Printf("found the sequence 0\n")

		final_element := 0
		for i := len(sequence_history) - 1; i >= 0; i-- {
			s := sequence_history[i]
			final_element += s[len(s)-1]
		}
		fmt.Printf("next element is %d\n", final_element)
		total_elements += final_element
	}
	fmt.Printf("The sum of the final elements is %d\n", total_elements)
}
