package main

import (
	"adventcode/reader"
	"fmt"
)

func getPlatformMatrix(lines []string) [][]string {
	var platform [][]string

	for i := 0; i < len(lines); i++ {
		platform_line := make([]string, 0)
		for j := 0; j < len(lines[i]); j++ {
			platform_line = append(platform_line, string(lines[i][j]))
		}
		platform = append(platform, platform_line)
	}

	return platform
}

func tiltPlatformNorth(platform [][]string) {
	// save top min positions
	min_free_space := make([]int, len(platform[0]))

	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == "#" {
				min_free_space[j] = i + 1
			} else if platform[i][j] == "O" {
				if i > min_free_space[j] {
					// tilt up
					platform[min_free_space[j]][j] = "O"
					platform[i][j] = "."
				}
				min_free_space[j]++
			}
		}
	}
}

func printPlatform(platform [][]string) {
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			fmt.Printf("%s", platform[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func sumPlatformLoad(platform [][]string) int {
	var sum_load int

	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			if platform[i][j] == "O" {
				sum_load += len(platform) - i
			}
		}
	}
	return sum_load
}

// rotate 90 degrees clockwise
func rotateMatrix(mat [][]string) [][]string {
	n := len(mat)
	rotated_mat := make([][]string, len(mat))
	for i := 0; i < n; i++ {
		rotated_mat[i] = make([]string, len(mat))
	}

	// transpose matrix
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated_mat[j][i] = mat[i][j]
		}
	}

	// swap first and last column
	var swap string
	for j := 0; j < n/2; j++ {
		for i := 0; i < n; i++ {
			swap = rotated_mat[i][j]
			rotated_mat[i][j] = rotated_mat[i][n-j-1]
			rotated_mat[i][n-j-1] = swap
		}
	}

	return rotated_mat
}

const (
	// N_CICLES = 3 // test
	N_CICLES = 1_000_000 // test
	// N_CICLES = 1_000_000_000
)

func main() {
	lines := reader.ReadLines("./day14/data/input.txt")
	platform := getPlatformMatrix(lines)
	fmt.Printf("platform: \n")
	printPlatform(platform)
	// step 1
	// tiltPlatformNorth(platform)
	// fmt.Printf("sum of platform load is %d", sumPlatformLoad(platform))

	// step 2
	for i := 0; i < N_CICLES; i++ {
		// fmt.Printf("cicle %d\n", i)
		for j := 0; j < 4; j++ {
			tiltPlatformNorth(platform)
			platform = rotateMatrix(platform)
		}
		// fmt.Printf("platform after cicle %d: \n", i+1)
		// printPlatform(platform)
	}
	printPlatform(platform)
	fmt.Printf("sum of platform load is %d\n", sumPlatformLoad(platform))
}
