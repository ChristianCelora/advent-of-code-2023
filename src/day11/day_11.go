package main

import (
	"adventcode/reader"
	"fmt"
	"math"
	"strconv"
)

const (
	// SPACE_EXPAND_FACTOR = 10 // test
	SPACE_EXPAND_FACTOR = 1_000_000
)

func ExpandRows(lines []string) []string {
	var expand bool
	var lines_expanded []string
	for i := 0; i < len(lines); i++ {
		expand = true
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] != '.' {
				expand = false
				break
			}
		}
		lines_expanded = append(lines_expanded, lines[i])
		if expand {
			lines_expanded = append(lines_expanded, lines[i])
		}
	}

	return lines_expanded
}

func ExpandColumns(lines []string) []string {
	var expand bool
	var lines_expanded []string

	for i := 0; i < len(lines); i++ {
		lines_expanded = append(lines_expanded, "")
	}

	for i := 0; i < len(lines[0]); i++ {
		expand = true
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != '.' {
				expand = false
				break
			}
		}

		for j := 0; j < len(lines); j++ {
			lines_expanded[j] = lines_expanded[j] + string(lines[j][i])
			if expand {
				lines_expanded[j] = lines_expanded[j] + string(lines[j][i])
			}
		}
	}
	return lines_expanded
}

func findEmptyRows(lines []string) map[int]bool {
	var empty_rows map[int]bool
	var is_empty bool

	empty_rows = make(map[int]bool)
	for i, row := range lines {
		is_empty = true
		for _, c := range row {
			if c != '.' {
				is_empty = false
			}
		}

		if is_empty {
			empty_rows[i] = true
		}
	}

	return empty_rows
}

func findEmptyColumns(lines []string) map[int]bool {
	var empty_columns map[int]bool
	var is_empty bool

	empty_columns = make(map[int]bool)
	for i := 0; i < len(lines[0]); i++ {
		is_empty = true
		for j := 0; j < len(lines); j++ {
			if lines[j][i] != '.' {
				is_empty = false
				break
			}
		}

		if is_empty {
			empty_columns[i] = true
		}
	}

	return empty_columns
}

func initSpace(lines []string) [][]int {
	var space [][]int
	var row []int
	var value int
	var galaxy_id int

	for i := 0; i < len(lines); i++ {
		row = make([]int, 0)
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '.' {
				value = 0
			} else if lines[i][j] == '#' {
				galaxy_id++
				value = galaxy_id
			}
			row = append(row, value)
		}
		space = append(space, row)
	}

	return space
}

type Cell struct {
	x     int
	y     int
	steps int
}

func walkAllTheMatrix(space [][]int, start_x int, start_y int) map[int]int {
	var distances map[int]int
	var visited map[string]bool
	var queue []Cell
	queue = append(queue, Cell{
		x:     start_x,
		y:     start_y,
		steps: 0,
	})
	directions := [4]struct {
		x int
		y int
	}{
		{
			x: -1,
			y: 0,
		},
		{
			x: 1,
			y: 0,
		},
		{
			x: 0,
			y: -1,
		},
		{
			x: 0,
			y: 1,
		},
	}

	distances = make(map[int]int)
	visited = make(map[string]bool)

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:] // remove first element
		x_str := strconv.Itoa(cell.x)
		y_str := strconv.Itoa(cell.y)
		visited[x_str+"_"+y_str] = true

		if space[cell.x][cell.y] > 0 {
			value, found := distances[space[cell.x][cell.y]]
			if found {
				if cell.steps < value {
					distances[space[cell.x][cell.y]] = cell.steps
				}
			} else {
				distances[space[cell.x][cell.y]] = cell.steps
			}
		}

		for i := 0; i < len(directions); i++ {
			new_x := cell.x + directions[i].x
			new_y := cell.y + directions[i].y
			if new_x < 0 || new_x > len(space)-1 {
				continue
			}
			if new_y < 0 || new_y > len(space[0])-1 {
				continue
			}

			new_x_str := strconv.Itoa(new_x)
			new_y_str := strconv.Itoa(new_y)
			is_visited := visited[new_x_str+"_"+new_y_str]
			if is_visited {
				continue
			}

			queue = append(queue, Cell{new_x, new_y, cell.steps + 1})
		}
	}

	return distances
}

func findDistancesFromGalaxy(space [][]int, start_x int, start_y int, empty_rows map[int]bool, empty_columns map[int]bool) map[int]int {
	var distances map[int]int
	distances = make(map[int]int)

	for i := 0; i < len(space); i++ {
		for j := 0; j < len(space[i]); j++ {
			if start_x == i && start_y == j {
				continue
			}
			if space[i][j] != 0 {
				distances[space[i][j]] = int(math.Abs(float64(start_x-i))) + int(math.Abs(float64(start_y-j)))

				// increase distance if empty row is found
				for k := min(start_x, i); k < max(start_x, i); k++ {
					if empty_rows[k] {
						// fmt.Printf("found empty row (%d) from %d (%d, %d) to %d (%d, %d)\n", k, space[start_x][start_y], start_x, start_y, space[i][j], i, j)
						distances[space[i][j]] += SPACE_EXPAND_FACTOR - 1 // -1 cause we alredy counted the empty row in the distance
					}
				}
				// increase distance if empty column is found
				for k := min(start_y, j); k < max(start_y, j); k++ {
					if empty_columns[k] {
						// fmt.Printf("found empty column (%d) from %d (%d, %d) to %d (%d, %d)\n", k, space[start_x][start_y], start_x, start_y, space[i][j], i, j)
						distances[space[i][j]] += SPACE_EXPAND_FACTOR - 1 // -1 cause we alredy counted the empty column in the distance
					}
				}
			}
		}
	}

	return distances
}

func main() {
	var sum_distances int
	var space [][]int
	lines := reader.ReadLines("./day11/data/input_final.txt")
	/**
	fmt.Println("-----------")
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			fmt.Printf("%s", string(lines[i][j]))
		}
		fmt.Println()
	}
	fmt.Println("-----------")
	*/
	/**
	lines = ExpandRows(lines)
	lines = ExpandColumns(lines)
	*/

	empty_rows := findEmptyRows(lines)
	empty_columns := findEmptyColumns(lines)
	// fmt.Printf("empty rows are %v\n", empty_rows)
	// fmt.Printf("empty columns are %v\n", empty_columns)

	space = initSpace(lines)

	for i := 0; i < len(space); i++ {
		for j := 0; j < len(space[i]); j++ {
			if space[i][j] != 0 {
				// distances := walkAllTheMatrix(space, i, j) // too slow! Remind: there are no obstacles
				distances := findDistancesFromGalaxy(space, i, j, empty_rows, empty_columns)
				// fmt.Printf("distances from galaxy %d: %v\n", space[i][j], distances)
				for _, d := range distances {
					sum_distances += d
				}
			}
		}
	}
	fmt.Printf("The sum of distances from galaxies %d\n", sum_distances/2)
}
