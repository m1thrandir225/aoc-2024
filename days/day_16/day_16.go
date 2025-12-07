package day_16

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

//Parse the input into a grid of strings
//Find the starting node -> S
//Find a path to node E
//Find all paths to node E
//Calculate the points for each move

const (
	N uint8 = iota
	E
	S
	W
)

type Point struct {
	x int
	y int
}

type Reindeer struct {
	position  Point
	direction uint8
}

type Grid [][]string

func Day16PartOne(filename string) {
	grid := make(Grid, 0)
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		grid = append(grid, strings.Fields(scanner.Text()))
	}
	startNode, err := grid.getStartNode()
	if err != nil {
		panic(err)
	}
	endNode, err := grid.getEndNode()
	if err != nil {
		panic(err)
	}

	fmt.Println(grid)
}

func (grid *Grid) getStartNode() (Reindeer, error) {
	for x, row := range *grid {
		for y, cell := range row {
			if cell == "S" {
				return Reindeer{
					Point{x, y},
					E,
				}, nil
			}
		}
	}
	return Reindeer{}, errors.New("could not find starting point")
}

func (grid *Grid) getEndNode() (Point, error) {
	for x, row := range *grid {
		for y, cell := range row {
			if cell == "E" {
				return Point{x, y}, nil
			}
		}
	}
	return Point{}, errors.New("could not find ending point")
}

func (point *Reindeer) Rotate(direction uint8) {
	switch direction {
	case 0: // North
		point.direction = (point.direction + 1) % 4
	case 1: // East
	case 2: // South
	case 3: // West
	}
}

func Day16PartTwo(filename string) {}
