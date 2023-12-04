package main

import (
	"adventcode/reader"
	"fmt"
	"strconv"
	"unicode"
)

// static matrix for now
const (
	// MATRIX_SIZE = 10 // for test
	MATRIX_SIZE = 140
)

func CreateMatrix() *[MATRIX_SIZE][MATRIX_SIZE]rune {
	var matrix [MATRIX_SIZE][MATRIX_SIZE]rune
	return &matrix
}

func InsertLineInMatrix(mat *[MATRIX_SIZE][MATRIX_SIZE]rune, line string, row int) {
	for i := range mat[row] {
		mat[row][i] = rune(line[i])
	}
}

type Symbol struct {
	symbol rune
	x      int
	y      int
}

func FindSymbols(mat *[MATRIX_SIZE][MATRIX_SIZE]rune) []Symbol {
	var res []Symbol
	for i, row := range mat {
		for j, cell := range row {
			s := string(cell)
			if s == "*" ||
				s == "&" ||
				s == "-" ||
				s == "/" ||
				s == "@" ||
				s == "#" ||
				s == "$" ||
				s == "=" ||
				s == "%" ||
				s == "+" {
				s := Symbol{
					symbol: cell,
					x:      i,
					y:      j,
				}
				res = append(res, s)
			}
		}
	}
	return res
}

type GearParts struct {
	parts       []int
	gear_ratios []int
}

func GetEngineParts(mat *[MATRIX_SIZE][MATRIX_SIZE]rune, symbols []Symbol) GearParts {
	var res GearParts
	var cell_x, cell_y int
	directions := [3]int{-1, 0, 1}
	for k := 0; k <= len(symbols)-1; k++ {
		s := symbols[k]
		var numbers []int
		for i := 0; i <= len(directions)-1; i++ {
			for j := 0; j <= len(directions)-1; j++ {
				delta_x := directions[i]
				delta_y := directions[j]
				cell_x = s.x + delta_x
				cell_y = s.y + delta_y

				if delta_y == 0 && delta_x == 0 {
					continue
				}
				if cell_x < 0 || cell_x >= MATRIX_SIZE {
					continue // out of bounds
				}
				if cell_y < 0 || cell_y >= MATRIX_SIZE {
					continue // out of bounds
				}

				if unicode.IsDigit(mat[cell_x][cell_y]) {
					// find number
					number := ExtractFullNumber(mat, cell_x, cell_y)
					numbers = append(numbers, number)
					res.parts = append(res.parts, number)
				}
			}
		}
		if len(numbers) == 2 {
			gear_ratio := 1
			for _, gr := range numbers {
				gear_ratio = gear_ratio * gr
			}
			res.gear_ratios = append(res.gear_ratios, gear_ratio)
		}
	}

	return res
}

func ExtractFullNumber(mat *[MATRIX_SIZE][MATRIX_SIZE]rune, x int, y int) int {
	// walk matrix from x,y to the left and right until we found a '.' or the edge
	var s string
	var i int
	if unicode.IsDigit(mat[x][y]) {
		s = string(mat[x][y])
	}
	i = y - 1
	for i >= 0 && unicode.IsDigit(mat[x][i]) {
		s = string(mat[x][i]) + s
		mat[x][i] = '.' // do not count the number again
		i = i - 1
	}

	i = y + 1
	for i < MATRIX_SIZE && unicode.IsDigit(mat[x][i]) {
		s = s + string(mat[x][i])
		mat[x][i] = '.' // do not count the number again
		i = i + 1
	}

	number, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return number
}

func main() {
	var sum, sum_gear_ratio int
	lines := reader.ReadLines("./day03/data/input1_2.txt")
	matrix := CreateMatrix()
	for i, line := range lines {
		InsertLineInMatrix(matrix, line, i)
	}
	symbols := FindSymbols(matrix)
	fmt.Printf("symbols found: %d\n", len(symbols))

	// part 1
	gear_data := GetEngineParts(matrix, symbols)
	for _, part := range gear_data.parts {
		sum = sum + part
	}
	fmt.Printf("The sum of the parts is %d\n", sum)

	// part 2
	for _, gr := range gear_data.gear_ratios {
		sum_gear_ratio = sum_gear_ratio + gr
	}
	fmt.Printf("The sum of the gear ratios is %d\n", sum_gear_ratio)
}
