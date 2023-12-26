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

func FindReflectionPoint(maze []string) (int, int) {
	// vertical
	rotated_maze := rotateMatrix(maze)
	rp := findRP(rotated_maze, 0)
	rp2 := findRP(rotated_maze, 1)

	// horizontal
	if rp == -1 {
		rp = findRP(maze, 0) * 100
	}
	if rp2 == -1 {
		rp2 = findRP(maze, 1) * 100
	}

	return rp, rp2
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

func findRP(maze []string, accepted_diffs int) int {
	for i := 1; i < len(maze); i++ {
		if maze[i] == maze[i-1] {
			if isReflection(maze, i, accepted_diffs) {
				return i
			}
		} else if countStringDiff(maze[i], maze[i-1]) == accepted_diffs {
			if isReflection(maze, i, accepted_diffs) {
				return i
			}
		}
	}
	return -1
}

func isReflection(maze []string, rp int, accepted_diffs int) bool {
	var total_diffs int
	for j := 0; j < min(rp, len(maze)-rp); j++ {
		total_diffs += countStringDiff(maze[rp-j-1], maze[rp+j])
		if total_diffs > accepted_diffs {
			break
		}
	}
	return total_diffs == accepted_diffs
}

func countStringDiff(s1 string, s2 string) int {
	var n int
	for i := 0; i < min(len(s1), len(s2)); i++ {
		if s1[i] != s2[i] {
			n++
		}
	}
	return n
}

func main() {
	var sum_reflection_points int
	var sum_reflection_points2 int
	lines := reader.ReadLines("./day13/data/input_final.txt")
	mazes := getMazes(lines)

	for _, maze := range mazes {
		reflection_point, reflection_point2 := FindReflectionPoint(maze)
		// fmt.Printf("reflection point %d, reflection point2 %d \n", reflection_point, reflection_point2)
		sum_reflection_points += reflection_point
		sum_reflection_points2 += reflection_point2
	}

	// step 1
	fmt.Printf("the sum of all reflection points is %d\n", sum_reflection_points)
	// step 2
	fmt.Printf("the sum of all reflection points with smudge is %d\n", sum_reflection_points2)
}
