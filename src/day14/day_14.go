package main

import (
	"adventcode/reader"
	"fmt"
	"slices"
)

const (
	N = 10 // test
	// N = 140

	// N_CICLES = 3 // test
	N_CICLES = 1_000_000 // test
	// N_CICLES = 1_000_000_000
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

func printPlatformFromObjects(objects []obj) {
	var printed_mat [N][N]string
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			printed_mat[i][j] = "."
		}
	}

	for _, obj := range objects {
		printed_mat[obj.x][obj.y] = string(obj.symbol)
	}

	printPlatform(printed_mat)
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

type obj struct {
	symbol byte
	x      int
	y      int
}

func sumObjectsLoad(objects []obj, max_load int) int {
	var sum_load int

	for _, obj := range objects {
		if obj.symbol == 'O' {
			sum_load += max_load - obj.x
		}
	}

	return sum_load
}

func tiltPlatformNorthFaster(objects []obj) {
	// divide each column
	var columns [N][]*obj
	for i := 0; i < len(objects); i++ {
		columns[objects[i].y] = append(columns[objects[i].y], &objects[i])
	}

	// sort by x asc
	for _, column_objs := range columns {
		slices.SortFunc(column_objs, func(a, b *obj) int {
			return a.x - b.x
		})
	}

	// move the objects based on free_position
	for _, column_objs := range columns {
		free_position := 0
		for _, obj := range column_objs {
			if obj.symbol == '#' {
				free_position = obj.x + 1
			} else if obj.symbol == 'O' {
				obj.x = free_position
				free_position++
			}
		}
	}
}

func tiltPlatformSouthFaster(objects []obj) {
	// divide each column
	var columns [N][]*obj
	for i := 0; i < len(objects); i++ {
		columns[objects[i].y] = append(columns[objects[i].y], &objects[i])
	}

	// sort by x desc
	for _, column_objs := range columns {
		slices.SortFunc(column_objs, func(a, b *obj) int {
			return b.x - a.x
		})
	}

	// move the objects based on free_position
	for _, column_objs := range columns {
		free_position := N - 1
		for _, obj := range column_objs {
			if obj.symbol == '#' {
				free_position = obj.x - 1
			} else if obj.symbol == 'O' {
				obj.x = free_position
				free_position--
			}
		}
	}
}

func tiltPlatformEastFaster(objects []obj) {
	// divide each row
	var rows [N][]*obj
	for i := 0; i < len(objects); i++ {
		rows[objects[i].x] = append(rows[objects[i].x], &objects[i])
	}

	// sort by y desc
	for _, row_objs := range rows {
		slices.SortFunc(row_objs, func(a, b *obj) int {
			return b.y - a.y
		})
	}

	// move the objects based on free_position
	for _, column_objs := range rows {
		free_position := N - 1
		for _, obj := range column_objs {
			if obj.symbol == '#' {
				free_position = obj.y - 1
			} else if obj.symbol == 'O' {
				obj.y = free_position
				free_position--
			}
		}
	}
}

func tiltPlatformWestFaster(objects []obj) {
	// divide each row
	var rows [N][]*obj
	for i := 0; i < len(objects); i++ {
		rows[objects[i].x] = append(rows[objects[i].x], &objects[i])
	}

	// sort by y asc
	for _, row_objs := range rows {
		slices.SortFunc(row_objs, func(a, b *obj) int {
			return a.y - b.y
		})
	}

	// move the objects based on free_position
	for _, column_objs := range rows {
		free_position := 0
		for _, obj := range column_objs {
			if obj.symbol == '#' {
				free_position = obj.y + 1
			} else if obj.symbol == 'O' {
				obj.y = free_position
				free_position++
			}
		}
	}
}

func main() {
	lines := reader.ReadLines("./day14/data/input.txt")
	platform := getPlatformMatrix(lines)
	fmt.Printf("platform: \n")
	// step 1
	tiltPlatformNorth(&platform)
	printPlatform(platform)
	fmt.Printf("sum of platform load is %d\n", sumPlatformLoad(platform))

	// // step 2
	// for i := 0; i < N_CICLES; i++ {
	// 	fmt.Printf("cicle %d\n", i)
	// 	// for j := 0; j < 4; j++ {
	// 	// 	tiltPlatformNorth(&platform)
	// 	// 	platform = rotateMatrix(platform)
	// 	// }
	// 	tiltPlatformNorth(&platform)
	// 	tiltPlatformWest(&platform)
	// 	tiltPlatformSouth(&platform)
	// 	tiltPlatformEast(&platform)
	// }
	// // printPlatform(platform)
	// fmt.Printf("sum of platform load after %d cycles is %d\n", N_CICLES, sumPlatformLoad(platform))

	// solution 2 - do not consider the empty space
	var objects []obj
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == '#' || lines[i][j] == 'O' {
				objects = append(objects, obj{
					symbol: lines[i][j],
					x:      i,
					y:      j,
				})
			}
		}
	}

	// step 1
	// tiltPlatformNorthFaster(objects)
	// printPlatformFromObjects(objects)
	// fmt.Printf("sum of platform load is %d\n", sumObjectsLoad(objects, N))

	// step 2
	for i := 0; i < N_CICLES; i++ {
		fmt.Printf("cicle %d\n", i)
		tiltPlatformNorthFaster(objects)
		tiltPlatformWestFaster(objects)
		tiltPlatformSouthFaster(objects)
		tiltPlatformEastFaster(objects)
		// printPlatformFromObjects(objects)
	}
	fmt.Printf("sum of platform load is %d\n", sumObjectsLoad(objects, N))
}
