package main

import (
	"adventcode/reader"
	"fmt"
)

func getMazes(lines []string) [][]string {
	var mazes [][]string
	var line string
	var maze []string

	maze = make([]string, 0)
	for i := 0; i < len(lines); i++ {
		line = lines[i]
		if line == "" {
			mazes = append(mazes, maze)
			maze = make([]string, 0)
		} else {
			maze = append(maze, line)
		}
	}
	mazes = append(mazes, maze)

	return mazes
}

func FindReflectionPoint(maze []string) int {
	// vertical
	rotated_maze := rotateMatrix(maze)
	rp := findRP(rotated_maze)
	if rp > -1 {
		return rp
	}

	// horizontal
	rp = findRP(maze)
	if rp > -1 {
		return rp * 100
	}

	return 0
}

func rotateMatrix(mat []string) []string {
	var rotated_mat []string
	for i := 0; i < len(mat[0]); i++ {
		line := ""
		for j := 0; j < len(mat); j++ {
			line += string(mat[j][i])
		}
		rotated_mat = append(rotated_mat, line)
	}

	return rotated_mat
}

func findRP(maze []string) int {
	for i := 1; i < len(maze); i++ {
		if maze[i] == maze[i-1] {
			if isReflection(maze, i) {
				return i
			}
		}
	}
	return -1
}

func isReflection(maze []string, rp int) bool {
	is_reflection := true
	for j := 0; j < min(rp, len(maze)-rp); j++ {
		if maze[rp-j-1] != maze[rp+j] {
			is_reflection = false
			break
		}
	}
	return is_reflection
}

func main() {
	var sum_reflection_points int
	lines := reader.ReadLines("./day13/data/input_final.txt")
	mazes := getMazes(lines)

	for _, maze := range mazes {
		reflection_point := FindReflectionPoint(maze)
		fmt.Printf("reflection point is %d\n", reflection_point)
		sum_reflection_points += reflection_point
	}

	fmt.Printf("the sum of all reflection points is %d\n", sum_reflection_points)

}
