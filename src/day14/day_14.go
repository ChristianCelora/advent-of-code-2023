package main

import (
	"adventcode/reader"
	"fmt"
)

const (
	N = 10 // test
	// N = 140

	// N_CICLES = 3 // test
	// N_CICLES = 1_000_000 // test
	N_CICLES = 1_000_000_000
)

func getPlatformMatrix(lines []string) [N][N]string {
	var platform [N][N]string

	for i := 0; i < len(lines); i++ {
		// platform_line := make([]string, 0)
		for j := 0; j < len(lines[i]); j++ {
			// platform_line = append(platform_line, string(lines[i][j]))
			platform[i][j] = string(lines[i][j])
		}
		// platform = append(platform, platform_line)
	}

	return platform
}

func tiltPlatformNorth(platform *[N][N]string) {
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

func tiltPlatformEast(platform *[N][N]string) {
	min_free_space := make([]int, len(platform))
	for i := 0; i < len(min_free_space); i++ {
		min_free_space[i] = len(platform) - 1
	}

	for i := 0; i < len(platform); i++ {
		for j := len(platform[0]) - 1; j >= 0; j-- {
			if platform[i][j] == "#" {
				min_free_space[i] = j - 1
			} else if platform[i][j] == "O" {
				if j < min_free_space[i] {
					// tilt right
					platform[i][min_free_space[i]] = "O"
					platform[i][j] = "."
				}
				min_free_space[i]--
			}
		}
	}
}

func tiltPlatformSouth(platform *[N][N]string) {
	min_free_space := make([]int, len(platform))
	for i := 0; i < len(min_free_space); i++ {
		min_free_space[i] = len(platform) - 1
	}

	for i := len(platform[0]) - 1; i >= 0; i-- {
		for j := len(platform[0]) - 1; j >= 0; j-- {
			if platform[i][j] == "#" {
				min_free_space[j] = i - 1
			} else if platform[i][j] == "O" {
				if i < min_free_space[j] {
					// tilt bottom
					platform[min_free_space[j]][j] = "O"
					platform[i][j] = "."
				}
				min_free_space[j]--
			}
		}
	}
}

func tiltPlatformWest(platform *[N][N]string) {
	min_free_space := make([]int, len(platform))

	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform); j++ {
			if platform[i][j] == "#" {
				min_free_space[i] = j + 1
			} else if platform[i][j] == "O" {
				if j > min_free_space[i] {
					// tilt left
					platform[i][min_free_space[i]] = "O"
					platform[i][j] = "."
				}
				min_free_space[i]++
			}
		}
	}
}

func printPlatform(platform [N][N]string) {
	for i := 0; i < len(platform); i++ {
		for j := 0; j < len(platform[i]); j++ {
			fmt.Printf("%s", platform[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func sumPlatformLoad(platform [N][N]string) int {
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
func rotateMatrix(mat [N][N]string) [N][N]string {
	n := len(mat)
	var rotated_mat [N][N]string

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

func copyMatrix[T any](matrix [N][N]T) [N][N]T {
	var duplicate [N][N]T
	for i := range matrix {
		for j := range matrix[i] {
			duplicate[i][j] = matrix[i][j]
		}
	}

	return duplicate
}

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
		fmt.Printf("cicle %d\n", i)
		// for j := 0; j < 4; j++ {
		// 	tiltPlatformNorth(&platform)
		// 	// platform = rotateMatrix(platform)
		// }
		tiltPlatformNorth(&platform)
		tiltPlatformWest(&platform)
		tiltPlatformSouth(&platform)
		tiltPlatformEast(&platform)
	}
	// printPlatform(platform)
	fmt.Printf("sum of platform load after %d cycles is %d\n", N_CICLES, sumPlatformLoad(platform))
}
