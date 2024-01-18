package main

import (
	"adventcode/reader"
	"fmt"
)

const (
	N = 10 // test
	// N = 140
)

type Cell struct {
	symbol    byte
	energized bool
}

func createMatrix(lines []string) [N][N]Cell {
	var matrix [N][N]Cell

	for i, line := range lines {
		for j, c := range line {
			matrix[i][j] = Cell{
				symbol: byte(c),
			}
		}
	}

	return matrix
}

type Beam struct {
	x     int
	y     int
	dir_x int8 // 1, 0, -1
	dir_y int8 // 1, 0, -1
}

func isSplitter(c Cell) bool {
	return c.symbol == '-' || c.symbol == '|'
}

func isMirror(c Cell) bool {
	return c.symbol == '/' || c.symbol == '\\'
}

func main() {
	var sum_energized_cells int
	var beam Beam
	lines := reader.ReadLines("./day16/data/input1_1.txt")
	matrix := createMatrix(lines)
	beams := []Beam{
		{0, 0, 0, 1},
	}

	for len(beams) > 0 {
		// pop beam
		beam = beams[len(beams)-1]
		beams = beams[:len(beams)-1]
		fmt.Printf("beam %v, beams len %d\n", beam, len(beams))

		// move beam
		for beam.x >= 0 && beam.x < N && beam.y >= 0 && beam.y < N {
			beam.x += int(beam.dir_x)
			beam.y += int(beam.dir_y)

			// if out-of-bounds pass to next beam
			if beam.x < 0 || beam.x >= N || beam.y < 0 || beam.y >= N {
				fmt.Printf("beam out of bounds\n")
				break
			}

			fmt.Printf("cell %s: (%d, %d)\n", string(matrix[beam.x][beam.y].symbol), beam.x, beam.y)
			matrix[beam.x][beam.y].energized = true

			if isSplitter(matrix[beam.x][beam.y]) {
				fmt.Printf("splitter found %s (%d, %d)\n", string(matrix[beam.x][beam.y].symbol), beam.x, beam.y)
				if matrix[beam.x][beam.y].symbol == '|' {
					// check if opposite direction
					if beam.dir_y == 1 || beam.dir_y == -1 {
						beams = append(beams, Beam{
							x:     beam.x,
							y:     beam.y,
							dir_x: 1,
							dir_y: 0,
						})
						beams = append(beams, Beam{
							x:     beam.x,
							y:     beam.y,
							dir_x: -1,
							dir_y: 0,
						})
						break
					}
				} else if matrix[beam.x][beam.y].symbol == '-' {
					// check if opposite direction
					if beam.dir_x == 1 || beam.dir_x == -1 {
						beams = append(beams, Beam{
							x:     beam.x,
							y:     beam.y,
							dir_x: 0,
							dir_y: 1,
						})
						beams = append(beams, Beam{
							x:     beam.x,
							y:     beam.y,
							dir_x: 0,
							dir_y: -1,
						})
						break
					}
				}
			}

			if isMirror(matrix[beam.x][beam.y]) {
				fmt.Printf("mirror found %s (%d, %d)\n", string(matrix[beam.x][beam.y].symbol), beam.x, beam.y)
				fmt.Printf("old direction (%d, %d)", beam.dir_x, beam.dir_y)
				if matrix[beam.x][beam.y].symbol == '\\' {
					if beam.dir_x == 1 || beam.dir_x == -1 {
						beam.dir_y = beam.dir_x
						beam.dir_x = 0
					} else {
						beam.dir_x = beam.dir_y
						beam.dir_y = 0
					}
				} else if matrix[beam.x][beam.y].symbol == '/' {
					if beam.dir_x == 1 || beam.dir_x == -1 {
						beam.dir_y = beam.dir_x * -1
						beam.dir_x = 0
					} else {
						beam.dir_x = beam.dir_y * -1
						beam.dir_y = 0
					}
				}
				fmt.Printf("new direction (%d, %d)", beam.dir_x, beam.dir_y)
			}
		}
	}

	for _, row := range matrix {
		for _, cell := range row {
			if cell.energized {
				sum_energized_cells++
			}
		}
	}

	fmt.Printf("Sum of energized cells is %d\n", sum_energized_cells)
}
