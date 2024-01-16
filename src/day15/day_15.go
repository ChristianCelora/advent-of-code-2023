package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
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

func hashLabel(step string) int {
	var hash int
	for i := 0; i < len(step); i++ {
		if step[i] == '=' || step[i] == '-' {
			break
		}
		hash += int(step[i])
		hash = hash * 17
		hash = hash % 256
	}

	return hash
}

func separateSteps(line string) []string {
	return strings.Split(line, ",")
}

type Lens struct {
	label     string
	focal_len int
}

func initBoxes() [N_BOXES][]Lens {
	var boxes [N_BOXES][]Lens
	return boxes
}

func removeStepFromBoxes(box_index int, boxes *[256][]Lens, hash_label int, index_map map[string]int, step_label string) {
	// shift index in map
	for i := box_index + 1; i <= len(boxes[hash_label])-1; i++ {
		map_key := boxes[hash_label][i].label
		index_map[map_key]--
	}
	// remove from map
	delete(index_map, step_label)
	// remove from slice
	boxes[hash_label] = append(boxes[hash_label][:box_index], boxes[hash_label][box_index+1:]...)
}

func printBoxes(boxes [N_BOXES][]Lens) {
	fmt.Println("Print Boxes")
	for i := 0; i < len(boxes); i++ {
		if len(boxes[i]) > 0 {
			fmt.Printf("Box %d: ", i)
			for j := 0; j < len(boxes[i]); j++ {
				fmt.Printf(" [%s, %d] ", boxes[i][j].label, boxes[i][j].focal_len)
			}
			fmt.Println("")
		}
	}
}

const (
	N_BOXES = 256
)

func main() {
	var step_hash_sum int
	var focal_length_sum int
	var index_map map[string]int

	lines := reader.ReadLines("./day15/data/input_final.txt")
	boxes := initBoxes()
	index_map = make(map[string]int)

	for _, line := range lines {
		steps := separateSteps(line)
		for _, step := range steps {
			// step 1
			hash := hashStep(step)
			step_hash_sum += hash

			// step 2
			hash_label := hashLabel(step)
			if strings.Index(step, "=") >= 0 {
				step_info := strings.Split(step, "=")
				fl, _ := strconv.Atoi(step_info[1])

				box_index, ok := index_map[step_info[0]]
				// fmt.Printf("step %s, hash %d, box_index %d\n", step, hash_label, box_index)
				if ok {
					// update
					boxes[hash_label][box_index].focal_len = fl
				} else {
					// insert
					boxes[hash_label] = append(
						boxes[hash_label],
						Lens{
							label:     step_info[0],
							focal_len: fl,
						},
					)
					index_map[step_info[0]] = len(boxes[hash_label]) - 1
				}
			} else {
				step_label := step[0 : len(step)-1]
				box_index, ok := index_map[step_label]
				// fmt.Printf("step %s, hash %d, box_index %d\n", step, hash_label, box_index)
				if ok {
					removeStepFromBoxes(box_index, &boxes, hash_label, index_map, step_label)
				}
			}
			// printBoxes(boxes)
			// fmt.Println()
		}
	}

	for i, box := range boxes {
		for j, lens := range box {
			focus_power := (i + 1) * (j + 1) * lens.focal_len
			focal_length_sum += focus_power
		}
	}

	fmt.Printf("Sum of all step hashes is %d\n", step_hash_sum)
	fmt.Printf("Sum of all focal length is %d\n", focal_length_sum)
}
