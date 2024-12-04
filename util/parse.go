package util

import (
	"bufio"
	"io"
	"os"
)

func ParseToGrid(filename string) [][]rune {
	grid := make([][]rune, 0)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
	}

	return grid
}

func ParseToString(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	b, err := io.ReadAll(f)
	return string(b)
}
