package main

import (
	"adventcode/reader"
	"fmt"
)

const (
	MATRIX_SIZE = 5
	// MATRIX_SIZE = 140
)

type Pipe struct {
	character    rune
	is_ground    bool
	is_start     bool
	is_connected bool
	x            int
	y            int
	top          *Pipe
	right        *Pipe
	left         *Pipe
	bottom       *Pipe
}

func GetMatrixFromLines(lines []string) [MATRIX_SIZE][MATRIX_SIZE]Pipe {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]Pipe

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		for j := 0; j < len(line); j++ {
			c := line[j]
			matrix[i][j] = Pipe{
				character: rune(c),
				is_ground: c == '.',
				is_start:  c == 'S',
				x:         j,
				y:         i,
			}
		}
	}

	// connect pipes
	var connections int
	for i := 0; i < MATRIX_SIZE; i++ {
		for j := 0; j < MATRIX_SIZE; j++ {
			pipe := &matrix[i][j]
			if pipe.is_ground {
				continue
			}

			connections = 0
			switch pipe.character {
			case '|':
				// top: F, 7, |
				if isConnected(matrix, i-1, j, []byte{'F', '7', '|'}) {
					connections++
					pipe.top = &matrix[i-1][j]
				}
				// bottom: J, L, |
				if isConnected(matrix, i+1, j, []byte{'J', 'L', '|'}) {
					connections++
					pipe.bottom = &matrix[i+1][j]
				}
			case '-':
				// left: L, F, -
				if isConnected(matrix, i, j-1, []byte{'L', 'F', '-'}) {
					connections++
					pipe.left = &matrix[i][j-1]
				}
				// right: J, 7, -
				if isConnected(matrix, i, j+1, []byte{'J', '7', '-'}) {
					connections++
					pipe.right = &matrix[i][j+1]
				}
			case 'L':
				// top: F, |, 7
				if isConnected(matrix, i-1, j, []byte{'F', '|', '7'}) {
					connections++
					pipe.top = &matrix[i-1][j]
				}
				// right: J, 7, -
				if isConnected(matrix, i, j+1, []byte{'J', '7', '-'}) {
					connections++
					pipe.right = &matrix[i][j+1]
				}
			case 'J':
				// top: F, |, 7
				if isConnected(matrix, i-1, j, []byte{'F', '|', '7'}) {
					connections++
					pipe.top = &matrix[i-1][j]
				}
				// left: L, F, -
				if isConnected(matrix, i, j-1, []byte{'L', 'F', '-'}) {
					connections++
					pipe.left = &matrix[i][j-1]
				}
			case '7':
				// left: L, F, -
				if isConnected(matrix, i, j-1, []byte{'L', 'F', '-'}) {
					connections++
					pipe.left = &matrix[i][j-1]
				}
				// bottom: J, L, |
				if isConnected(matrix, i+1, j, []byte{'J', 'L', '|'}) {
					connections++
					pipe.bottom = &matrix[i+1][j]
				}
			case 'F':
				// right: J, 7, -
				if isConnected(matrix, i, j+1, []byte{'J', '7', '-'}) {
					connections++
					pipe.right = &matrix[i][j+1]
				}
				// bottom: J, L, |
				if isConnected(matrix, i+1, j, []byte{'J', 'L', '|'}) {
					connections++
					pipe.bottom = &matrix[i+1][j]
				}
			case 'S':
				// top: F, |, 7
				if isConnected(matrix, i-1, j, []byte{'F', '|', '7'}) {
					connections++
					pipe.top = &matrix[i-1][j]
				}
				// right: J, 7, -
				if isConnected(matrix, i, j+1, []byte{'J', '7', '-'}) {
					connections++
					pipe.right = &matrix[i][j+1]
				}
				// left: L, F, -
				if isConnected(matrix, i, j-1, []byte{'L', 'F', '-'}) {
					connections++
					pipe.left = &matrix[i][j-1]
				}
				// bottom: J, L, |
				if isConnected(matrix, i+1, j, []byte{'J', 'L', '|'}) {
					connections++
					pipe.bottom = &matrix[i+1][j]
				}
			}
			if connections > 2 {
				pipe.is_connected = true
			}
		}
	}

	return matrix
}

func getStartingPoint(matrix [MATRIX_SIZE][MATRIX_SIZE]Pipe) Pipe {
	for _, row := range matrix {
		for _, cell := range row {
			if cell.is_start {
				return cell
			}
		}
	}

	return Pipe{}
}

type PipePath struct {
	pipe  *Pipe
	steps int
}

func findFurthestPipeFromStart(matrix [MATRIX_SIZE][MATRIX_SIZE]Pipe, start_x int, start_y int) int {
	var max_steps int
	var pipes_stack []PipePath
	var current_pipe PipePath
	pipes_stack = append(pipes_stack, PipePath{
		pipe:  &matrix[start_y][start_x],
		steps: 0,
	})

	for len(pipes_stack) > 0 {
		current_pipe = pipes_stack[len(pipes_stack)-1]
		fmt.Printf("current pipe %v\n", current_pipe.pipe)
		fmt.Printf("current pipe steps %d\n", current_pipe.steps)
		pipes_stack = removeElementFromSlice(pipes_stack)
		if current_pipe.pipe.is_ground {
			continue
		}
		if current_pipe.steps > max_steps {
			max_steps = current_pipe.steps
		}
		if current_pipe.pipe.top != nil {
			next_pipe := matrix[current_pipe.pipe.x-1][current_pipe.pipe.y]
			pipes_stack = append(pipes_stack, PipePath{
				pipe:  &next_pipe,
				steps: current_pipe.steps + 1,
			})
		}
		if current_pipe.pipe.left != nil {
			next_pipe := matrix[current_pipe.pipe.x][current_pipe.pipe.y-1]
			pipes_stack = append(pipes_stack, PipePath{
				pipe:  &next_pipe,
				steps: current_pipe.steps + 1,
			})
		}
		if current_pipe.pipe.right != nil {
			next_pipe := matrix[current_pipe.pipe.x][current_pipe.pipe.y+1]
			pipes_stack = append(pipes_stack, PipePath{
				pipe:  &next_pipe,
				steps: current_pipe.steps + 1,
			})
		}
		if current_pipe.pipe.bottom != nil {
			next_pipe := matrix[current_pipe.pipe.x+1][current_pipe.pipe.y]
			pipes_stack = append(pipes_stack, PipePath{
				pipe:  &next_pipe,
				steps: current_pipe.steps + 1,
			})
		}
		matrix[current_pipe.pipe.x][current_pipe.pipe.y].is_ground = true // visited
		matrix[current_pipe.pipe.x][current_pipe.pipe.y].character = rune(current_pipe.steps)
	}

	return max_steps
}

func removeElementFromSlice[T any](slice []T) []T {
	return slice[:len(slice)-1]
}

func isConnected(matrix [MATRIX_SIZE][MATRIX_SIZE]Pipe, i int, j int, available_conncetions []byte) bool {
	if i < 0 || i >= MATRIX_SIZE {
		return false
	} else if j < 0 || j >= MATRIX_SIZE {
		return false
	}

	for _, ac := range available_conncetions {
		if matrix[i][j].character == rune(ac) {
			return true
		}
	}
	return false
}

// func GetPipeFromCell(cell byte) Pipe {

// }

func main() {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]Pipe

	lines := reader.ReadLines("./day10/data/input1_1.txt")
	matrix = GetMatrixFromLines(lines)

	starting_cell := getStartingPoint(matrix)
	fmt.Printf("starting_cell %v\n", starting_cell)
	steps := findFurthestPipeFromStart(matrix, starting_cell.x, starting_cell.y)
	fmt.Printf("Further steps in loop are %d\n", steps)
}
