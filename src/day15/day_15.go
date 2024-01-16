package main

import (
	"adventcode/reader"
	"fmt"
	"strings"
)

func hashStep(step string) int {
	var hash int
	for i := 0; i < len(step); i++ {
		hash += int(step[i])
		hash = hash * 17
		hash = hash % 256
	}

	return hash
}

func separateSteps(line string) []string {
	return strings.Split(line, ",")
}

func main() {
	var step_hash_sum int
	lines := reader.ReadLines("./day15/data/input_final.txt")

	for _, line := range lines {
		steps := separateSteps(line)
		for _, step := range steps {
			hash := hashStep(step)
			step_hash_sum += hash
			// fmt.Printf("step %s, hash %d\n", step, hash)
		}
	}

	fmt.Printf("Sum of all step hashes is %d\n", step_hash_sum)
}
