package day8

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

func Day8PartTwo(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	grid := strings.Split(strings.TrimSpace(string(f)), "\n")
	antinodes := make(map[Point]bool)
	for _, positions := range getAntenna(grid) {
		for _, pair := range getCombinations(positions) {
			x1, y1 := pair[0].x, pair[0].y
			x2, y2 := pair[1].x, pair[1].y

			directions := [][4]int{
				{x2 - x1, y2 - y1, x1, y1},
				{x1 - x2, y1 - y2, x2, y2},
			}

			for _, dir := range directions {
				dx, dy, x, y := dir[0], dir[1], dir[2], dir[3]
				for isPointValid(x, y, grid) {
					p := Point{x, y}
					antinodes[p] = true
					x += dx
					y += dy
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}

func Day8PartOne(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	grid := strings.Split(strings.TrimSpace(string(data)), "\n")
	antinodes := make(map[Point]bool)

	for _, positions := range getAntenna(grid) {
		for _, pair := range getCombinations(positions) {
			x1, y1 := pair[0].x, pair[0].y
			x2, y2 := pair[1].x, pair[1].y

			p1 := Point{2*x2 - x1, 2*y2 - y1}
			p2 := Point{2*x1 - x2, 2*y1 - y2}

			if isPointValid(p1.x, p1.y, grid) {
				antinodes[p1] = true
			}

			if isPointValid(p2.x, p2.y, grid) {
				antinodes[p2] = true
			}
		}
	}
	fmt.Println(len(antinodes))
}

func getAntenna(grid []string) map[rune][]Point {
	antennas := make(map[rune][]Point)

	for x, row := range grid {
		for y, col := range row {
			if col != '.' {
				antennas[col] = append(antennas[col], Point{x, y})
			}
		}
	}
	return antennas
}

func getCombinations(arr []Point) [][2]Point {
	var result [][2]Point

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			result = append(result, [2]Point{arr[i], arr[j]})
		}
	}

	return result
}

func isPointValid(x, y int, grid []string) bool {
	return x >= 0 && x < len(grid) && y >= 0 && y < len(grid[0])
}
