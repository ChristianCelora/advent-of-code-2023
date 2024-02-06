package matrix_helper

import "fmt"

func PrintMatrix[T any](matrix [][]T) {
	fmt.Println("print matrix:")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j])
		}
		fmt.Println()
	}
}

func PrepareDynamicMatrix[T any](n_rows int, n_cols int) [][]T {
	matrix := make([][]T, n_rows)
	for i := range matrix {
		matrix[i] = make([]T, n_cols)
	}

	return matrix
}
