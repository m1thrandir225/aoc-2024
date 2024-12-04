package day4

import (
	"aoc_2024/util"
	"fmt"
)

func Day4PartOne(filename string) {
	grid := util.ParseToGrid(filename)

	word := "XMAS"
	res := searchWord(grid, word)
	fmt.Println(res)
}

func Day4PartTwo(filename string) {
	grid := util.ParseToGrid(filename)

	results := countXInGrid(grid)
	fmt.Println(results)
}

var directions = [8][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
	{1, 1},
	{1, -1},
	{-1, 1},
	{-1, -1},
}

var onlyDiagonalDirections = [4][2]int{
	{-1, -1},
	{1, 1},
	{-1, 1},
	{1, -1},
}

func isValidGrid(x int, y int, rows int, cols int) bool {
	return x >= 0 && y >= 0 && x < rows && y < cols
}

func checkGridDirection(grid [][]rune, word string, x int, y int, dx int, dy int) bool {
	rows, cols := len(grid), len(grid[0])
	for i := 0; i < len(word); i++ {
		nx, ny := x+i*dx, y+i*dy
		if !isValidGrid(nx, ny, rows, cols) || grid[nx][ny] != rune(word[i]) {
			return false
		}
	}
	return true

}

func searchWord(grid [][]rune, word string) int {
	counter := 0

	rows, cols := len(grid), len(grid[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			for _, dir := range directions {
				dx, dy := dir[0], dir[1]
				if checkGridDirection(grid, word, i, j, dx, dy) {
					counter++
				}
			}
		}
	}
	return counter
}

// brute force
func countXInGrid(grid [][]rune) int {
	count := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[0])-1; j++ {
			if grid[i][j] == 'A' {
				pattern := grid[i-1][j-1] == 'M' &&
					grid[i-1][j+1] == 'M' &&
					grid[i+1][j+1] == 'S' &&
					grid[i+1][j-1] == 'S'
				if pattern {
					count++
				}
				pattern = grid[i-1][j-1] == 'S' &&
					grid[i-1][j+1] == 'S' &&
					grid[i+1][j+1] == 'M' &&
					grid[i+1][j-1] == 'M'
				if pattern {
					count++
				}
				pattern = grid[i-1][j-1] == 'M' &&
					grid[i-1][j+1] == 'S' &&
					grid[i+1][j+1] == 'S' &&
					grid[i+1][j-1] == 'M'
				if pattern {
					count++
				}
				pattern = grid[i-1][j-1] == 'S' &&
					grid[i-1][j+1] == 'M' &&
					grid[i+1][j+1] == 'M' &&
					grid[i+1][j-1] == 'S'
				if pattern {
					count++
				}
			}
		}
	}
	return count
}
