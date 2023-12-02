package main

import (
	"strconv"
	"strings"
)

const (
	MAX_BLUE  = 14
	MAX_GREEN = 13
	MAX_RED   = 12
)

func ReadGameID(line string) int {
	first_space := strings.Index(line, " ")
	first_semicolon := strings.Index(line, ":")

	if first_space == -1 || first_semicolon == -1 {
		panic("error reading game ID")
	}

	id, err := strconv.Atoi(line[first_space+1 : first_semicolon])
	if err != nil {
		panic(err)
	}

	return id
}

type Set struct {
	blue  int
	green int
	red   int
}

func ReadCubeSets() []Set {
	panic("to do")
}

func main() {
	// lines := reader.ReadLines("./day02/data/input1_1.txt")
}
