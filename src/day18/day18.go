package main

import (
	"adventcode/matrix_helper"
	"adventcode/reader"
	"fmt"
	"strconv"
	"strings"
)

type DigInstruction struct {
	direction byte
	meters    int
}

func getDigInstructions(lines []string) []DigInstruction {
	var instructions []DigInstruction

	for _, line := range lines {
		instr := strings.Split(line, " ")
		if len(instr) != 3 {
			panic("error splitting instruction " + line)
		}
		dir := []byte(instr[0])
		val, _ := strconv.Atoi(instr[1])
		instructions = append(instructions, DigInstruction{dir[0], val})
	}

	return instructions
}

func prepareDig(instr []DigInstruction) (int, int, int, int) {
	var x, new_x, y, new_y int
	var min_x, max_x, min_y, max_y int

	for i := 0; i < len(instr); i++ {
		// move
		if instr[i].direction == 'D' {
			new_x = x + instr[i].meters
		} else if instr[i].direction == 'U' {
			new_x = x - instr[i].meters
		} else if instr[i].direction == 'L' {
			new_y = y - instr[i].meters
		} else if instr[i].direction == 'R' {
			new_y = y + instr[i].meters
		}

		if max_x < new_x {
			max_x = new_x
		} else if min_x > new_x {
			min_x = new_x
		}

		if max_y < new_y {
			max_y = new_y
		} else if min_y > new_y {
			min_y = new_y
		}

		// change x, y
		x = new_x
		y = new_y
	}

	return min_x, max_x, min_y, max_y
}

func dig(matrix [][]int, instr []DigInstruction, start_x int, start_y int) {
	var new_x, new_y int
	x := start_x
	y := start_y

	for i := 0; i < len(instr); i++ {
		// move
		if instr[i].direction == 'D' {
			new_x = x + instr[i].meters
		} else if instr[i].direction == 'U' {
			new_x = x - instr[i].meters
		} else if instr[i].direction == 'L' {
			new_y = y - instr[i].meters
		} else if instr[i].direction == 'R' {
			new_y = y + instr[i].meters
		}

		// fmt.Printf("x %d, new_x %d, inst %v\n", x, new_x, instr[i])
		for j := min(x, new_x); j <= max(x, new_x); j++ {
			matrix[j][y] = -1
		}

		// fmt.Printf("y %d, new_y %d, inst %v\n", y, new_y, instr[i])
		for j := min(y, new_y); j <= max(y, new_y); j++ {
			matrix[x][j] = -1
		}

		// change x, y
		x = new_x
		y = new_y

		// matrix_helper.PrintMatrix[int](matrix)
	}
}

func digInterior(matrix [][]int, instr []DigInstruction, start_x int, start_y int) {
	var new_x, new_y int
	x := start_x
	y := start_y

	for i := 0; i < len(instr); i++ {
		// move
		if instr[i].direction == 'D' {
			new_x = x + instr[i].meters
		} else if instr[i].direction == 'U' {
			new_x = x - instr[i].meters
		} else if instr[i].direction == 'L' {
			new_y = y - instr[i].meters
		} else if instr[i].direction == 'R' {
			new_y = y + instr[i].meters
		}

		if i > 0 {
			// prev_instr := instr[i-1]

		}

		// change x, y
		x = new_x
		y = new_y
	}
}

func main() {
	lines := reader.ReadLines("./day18/data/input1_1.txt")
	dig_instructions := getDigInstructions(lines)

	min_x, max_x, min_y, max_y := prepareDig(dig_instructions)
	dx := max_x - min_x + 1
	dy := max_y - min_y + 1
	fmt.Printf("Creating a %d x %d matrix\n", dx, dy)
	dynamic_matrix := matrix_helper.PrepareDynamicMatrix[int](dx, dy)
	matrix_helper.PrintMatrix[int](dynamic_matrix)

	start_x := 0 - min_x
	start_y := 0 - min_y
	dig(dynamic_matrix, dig_instructions, start_x, start_y)
	matrix_helper.PrintMatrix[int](dynamic_matrix)
}
