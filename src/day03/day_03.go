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
			// if regexp.MatchString("[\*\&\-\/\@\#\$\=\%\+]", cell) {
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
				// fmt.Printf("sym found: %s, %d, %d\n", string(s.symbol), s.x, s.y)
				res = append(res, s)
			}
		}
	}
	return res
}

func GetEngineParts(mat *[MATRIX_SIZE][MATRIX_SIZE]rune, symbols []Symbol) []int {
	var res []int
	var cell_x, cell_y int
	directions := [3]int{-1, 0, 1}
	for k := 0; k <= len(symbols)-1; k++ {
		s := symbols[k]
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
					res = append(res, number)
				}
			}
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
	var sum int
	lines := reader.ReadLines("./day03/data/input1_2.txt")
	matrix := CreateMatrix()
	for i, line := range lines {
		InsertLineInMatrix(matrix, line, i)
	}
	symbols := FindSymbols(matrix)
	fmt.Printf("symbols found: %d\n", len(symbols))
	parts := GetEngineParts(matrix, symbols)
	for _, part := range parts {
		sum = sum + part
	}

	fmt.Printf("The sum of the parts is %d\n", sum)
}
