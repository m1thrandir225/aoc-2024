package day6

import (
	"aoc_2024/util"
	"fmt"
)

type Position struct {
	x, y int
}

func Day6PartOne(filename string) {
	grid := util.ParseToGrid(filename)
	row_start, col_start := getGuardStartingPosition(grid)
	guard := Guard{
		x:            row_start,
		y:            col_start,
		max_x:        len(grid),
		max_y:        len(grid[0]),
		passedRoutes: make([]Position, 0),
		direction:    "up",
	}
	guard.calculateRoute(grid)
}

func Day6PartTwo(filename string) {}

func getGuardStartingPosition(grid [][]rune) (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 94 {
				return i, j
			}
		}
	}
	return 0, 0
}

type Guard struct {
	x            int
	y            int
	max_y        int
	max_x        int
	passedRoutes []Position
	direction    string
}

func (guard *Guard) calculateRoute(grid [][]rune) {
	width := grid
	height := grid[0]
	for guard.x < len(width) {
		for guard.y < len(height) {
			toTurn, direction := guard.NeedsToTurn(grid)
			if toTurn {
				guard.turnDirection(grid, direction)
			}
			guard.move(grid)
			fmt.Println(len(guard.passedRoutes))
			fmt.Println("--------------------")
		}

	}
}

func (guard *Guard) move(grid [][]rune) {
	switch guard.direction {
	case "up":
		grid[guard.x][guard.y] = 46
		grid[guard.x-1][guard.y] = 96
		guard.x--
		guard.addToPositions(guard.x, guard.y)
	case "down":
		grid[guard.x][guard.y] = 46
		grid[guard.x+1][guard.y] = 96
		guard.x++
		guard.addToPositions(guard.x, guard.y)

	case "left":
		grid[guard.x][guard.y] = 46
		grid[guard.x][guard.y-1] = 96
		guard.y--
		guard.addToPositions(guard.x, guard.y)
	case "right":
		grid[guard.x][guard.y] = 46
		grid[guard.x][guard.y+1] = 96

		guard.y++
		guard.addToPositions(guard.x, guard.y)
	}
}

// Returns if the guard needs to turn and the next direction
func (guard *Guard) NeedsToTurn(grid [][]rune) (bool, string) {
	switch guard.direction {
	case "up":
		if grid[guard.x-1][guard.y] == 35 {
			return true, "right"
		}
	case "down":
		if grid[guard.x+1][guard.y] == 35 {
			return true, "left"
		}
	case "right":
		if grid[guard.x][guard.y+1] == 35 {
			return true, "down"
		}
	case "left":
		if grid[guard.x][guard.y-1] == 35 {
			return true, "up"
		}
	}
	return false, ""
}

func (guard *Guard) turnDirection(grid [][]rune, direction string) {
	switch direction {
	case "up":
		grid[guard.x][guard.y] = 96
		guard.direction = "up"
	case "down":
		grid[guard.x][guard.y] = 118
		guard.direction = "down"
	case "left":
		grid[guard.x][guard.y] = 60
		guard.direction = "left"
	case "right":
		grid[guard.x][guard.y] = 62
		guard.direction = "right"
	}
}

func (guard *Guard) addToPositions(row, col int) {
	// First check if position already exists
	for _, current := range guard.passedRoutes {
		if current.x == row && current.y == col {
			return // Exit if position already exists
		}
	}

	// If we've gotten here, the position is unique
	newPosition := Position{
		x: row,
		y: col,
	}
	guard.passedRoutes = append(guard.passedRoutes, newPosition)
}

func printGrid(grid [][]rune) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
