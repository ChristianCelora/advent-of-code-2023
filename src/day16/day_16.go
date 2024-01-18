package main

import (
	"adventcode/reader"
	"fmt"
)

const (
	// N = 10 // test
	// N = 17 // test
	N = 110
)

type Cell struct {
	symbol    byte
	energized bool
	visited   [4]bool // N E S W
}

func (c *Cell) isVisited(dir_x int8, dir_y int8) bool {
	visited_index := c.getVisitedIndex(dir_x, dir_y)
	if visited_index == -1 {
		return false // second param error is better
	}

	return c.visited[visited_index]
}

func (c *Cell) setVisited(dir_x int8, dir_y int8) {
	visited_index := c.getVisitedIndex(dir_x, dir_y)
	if visited_index != -1 {
		c.visited[visited_index] = true
	}
}

func (c *Cell) getVisitedIndex(dir_x int8, dir_y int8) int {
	if dir_x == 1 && dir_y == 0 {
		return 0
	} else if dir_x == -1 && dir_y == 0 {
		return 2
	} else if dir_x == 0 && dir_y == 1 {
		return 1
	} else if dir_x == 0 && dir_y == -1 {
		return 3
	}

	return -1
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

func getStartingBeam(dir int, idx int) Beam {
	var b Beam

	if dir == 0 { // top
		b = Beam{0, idx, 1, 0}
	} else if dir == 1 { // right
		b = Beam{idx, N - 1, 0, -1}
	} else if dir == 2 { // bottom
		b = Beam{N - 1, idx, -1, 0}
	} else if dir == 3 { // left
		b = Beam{idx, 0, 0, 1}
	}

	return b
}

func energizeCells(matrix *[N][N]Cell, starting_beam Beam) {
	var beam Beam
	beams := []Beam{
		starting_beam,
	}
	// matrix[starting_beam.x][starting_beam.y].energized = true

	for len(beams) > 0 {
		// pop beam
		beam = beams[len(beams)-1]
		beams = beams[:len(beams)-1]
		// fmt.Printf("\nbeam %v, beams len %d\n", beam, len(beams))

		for beam.x >= 0 && beam.x < N && beam.y >= 0 && beam.y < N {

			// if out-of-bounds pass to next beam
			/*if beam.x < 0 || beam.x >= N || beam.y < 0 || beam.y >= N {
				// fmt.Printf("beam out of bounds\n")
				break
			}*/

			if matrix[beam.x][beam.y].isVisited(beam.dir_x, beam.dir_y) {
				// fmt.Printf("alredy visited\n")
				break
			}

			// fmt.Printf("cell %s, visited %v: (%d, %d)\n", string(matrix[beam.x][beam.y].symbol), matrix[beam.x][beam.y].visited, beam.x, beam.y)
			matrix[beam.x][beam.y].energized = true
			matrix[beam.x][beam.y].setVisited(beam.dir_x, beam.dir_y)

			if isSplitter(matrix[beam.x][beam.y]) {
				// fmt.Printf("splitter found %s (%d, %d)\n", string(matrix[beam.x][beam.y].symbol), beam.x, beam.y)
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
				// fmt.Printf("mirror found %s (%d, %d)\n", string(matrix[beam.x][beam.y].symbol), beam.x, beam.y)

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

			}

			// move beam
			beam.x += int(beam.dir_x)
			beam.y += int(beam.dir_y)
		}
	}
}

func main() {
	var sum_energized_cells int
	var max_sum_energized_cells int
	var max_starting_pos [2]int
	var beam Beam
	lines := reader.ReadLines("./day16/data/input_final.txt")
	for i := 0; i < 4; i++ {
		for j := 0; j < N; j++ {
			matrix := createMatrix(lines)
			// beam = Beam{0, 0, 0, 1}
			fmt.Printf("\nStarting beam %v\n", beam)
			beam = getStartingBeam(i, j)
			energizeCells(&matrix, beam)

			sum_energized_cells = 0
			for _, row := range matrix {
				for _, cell := range row {
					if cell.energized {
						sum_energized_cells++
					}
				}
			}
			fmt.Printf("Energized cells %d\n", sum_energized_cells)

			if sum_energized_cells > max_sum_energized_cells {
				max_starting_pos = [2]int{beam.x, beam.y}
				max_sum_energized_cells = sum_energized_cells
			}

			// part 1
			if beam.x == 0 && beam.y == 0 {
				fmt.Printf("Sum of energized cells from top-left corner is %d\n", sum_energized_cells)
			}
		}
	}
	// part 2
	fmt.Printf("Max sum of energized cells is %d, from position (%d, %d)\n", max_sum_energized_cells, max_starting_pos[0], max_starting_pos[1])
}
