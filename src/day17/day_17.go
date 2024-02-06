package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
)

const (
	N = 13 // test
)

type Cell struct {
	x        int
	y        int
	loss     int
	visited  bool
	distance int
	prev     *Cell
}

func getMapMatrix(lines []string) [N][N]Cell {
	var matrix [N][N]Cell

	for i, line := range lines {
		for j, c := range line {
			loss, _ := strconv.Atoi(string(c))
			matrix[i][j] = Cell{
				x:        i,
				y:        j,
				loss:     loss,
				distance: 1_000_000,
			}
		}
	}

	return matrix
}

/*
func distanceBetweenPoints(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}
*/

// can be optimized
func CountStraightLine(cell Cell) int {
	var count_straight_cells int
	count_straight_cells = 1

	if cell.prev == nil {
		return count_straight_cells
	}

	var direction [2]int
	direction[0] = cell.x - cell.prev.x
	direction[1] = cell.y - cell.prev.y
	count_straight_cells++

	prev_cell := cell.prev
	for prev_cell != nil &&
		prev_cell.prev != nil &&
		prev_cell.x-prev_cell.prev.x == direction[0] &&
		prev_cell.y-prev_cell.prev.y == direction[1] {
		prev_cell = prev_cell.prev
		count_straight_cells++
	}

	return count_straight_cells
}

func main() {
	lines := reader.ReadLines("./day17/data/input1_1.txt")
	matrix := getMapMatrix(lines)
	matrix[0][0].distance = 0
	nodes := []Cell{
		matrix[0][0],
	}
	directions := [4][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	for len(nodes) > 0 {
		node := nodes[len(nodes)-1]
		nodes = nodes[:len(nodes)-1]
		fmt.Printf("\n\nnode (%d, %d)\n", node.x, node.y)

		for i := 0; i < len(directions); i++ {
			// out of bounds
			if node.x+directions[i][0] < 0 || node.x+directions[i][0] >= len(matrix) {
				fmt.Printf("x out of bounds\n")
				continue
			}
			if node.y+directions[i][1] < 0 || node.y+directions[i][1] >= len(matrix) {
				fmt.Printf("y out of bounds\n")
				continue
			}

			neighbour := &matrix[node.x+directions[i][0]][node.y+directions[i][1]]
			fmt.Printf("neighbour (%d, %d)\n", neighbour.x, neighbour.y)
			if neighbour.visited {
				fmt.Printf("neighbour visited\n")
				continue
			}
			sum_loss := node.distance + neighbour.loss
			fmt.Printf("new sum_loss %d, neighbour distance: %d, straight %d\n", sum_loss, neighbour.distance, CountStraightLine(*neighbour))

			if sum_loss < neighbour.distance {
				// if CountStraightLine(*neighbour) >= 4 {
				// 	fmt.Printf("straight longer than 4 tiles\n")
				// 	continue
				// }
				fmt.Printf("new sum_loss %d, n distance: %d\n", sum_loss, neighbour.distance)
				old_prev := neighbour.prev
				old_distance := neighbour.distance
				neighbour.distance = sum_loss
				neighbour.prev = &matrix[node.x][node.y]
				if CountStraightLine(*neighbour) > 4 {
					neighbour.distance = old_distance
					neighbour.prev = old_prev
					fmt.Printf("straight longer than 4 tiles neighbour (%d, %d)\n", neighbour.x, neighbour.y)
					continue
				}
				fmt.Printf("connecting node (%d, %d) to neighbour (%d, %d)\n", node.x, node.y, neighbour.x, neighbour.y)

				// insert sort DESC
				if len(nodes) == 0 {
					nodes = append(nodes, *neighbour)
				} else {
					var appended bool
					fmt.Printf("old nodes %v\n", nodes)
					for i := 0; i < len(nodes); i++ {
						if neighbour.distance > nodes[i].distance {
							nodes = append(nodes[0:i+1], nodes[i:]...)
							nodes[i] = *neighbour
							appended = true
							break
						}
					}
					if !appended {
						nodes = append(nodes, *neighbour)
					}
					fmt.Printf("new nodes %v\n", nodes)
				}
			}
		}
		matrix[node.x][node.y].visited = true
		fmt.Printf("node (%d, %d) visited\n", node.x, node.y)
	}

	fmt.Printf("final node %v\n", matrix[N-1][N-1])
	prev_node := matrix[N-1][N-1].prev
	for prev_node != nil {
		fmt.Printf("node (%d,%d) straight %d\n", prev_node.x, prev_node.y, CountStraightLine(*prev_node))
		prev_node = prev_node.prev
	}
}
