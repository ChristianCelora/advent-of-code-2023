package reader

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error in buffer: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
