package day6

import (
	"bufio"
	"fmt"
	"os"

	"github.com/lucianoq/container/set"
)

type State struct {
	Pos P
	Dir uint8
}

var Size = 130

const (
	N uint8 = iota
	E
	S
	W
)

type P struct{ x, y int }

func parseMap(filename string) (map[P]struct{}, P, uint8) {
	var pos P
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	m := map[P]struct{}{}
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			switch c {
			case '#':
				m[P{i, j}] = struct{}{}
			case '^':
				pos = P{i, j}
			}
		}
	}
	return m, pos, N
}

func Day6PartOne(filename string) {
	m, p, dir := parseMap(filename)
	visited := set.Set[P]{}
	visited.Add(p)

	for {
		var next P

		switch dir {
		case N:
			next = P{p.x - 1, p.y}
		case E:
			next = P{p.x, p.y + 1}
		case S:
			next = P{p.x + 1, p.y}
		case W:
			next = P{p.x, p.y - 1}
		}

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			break
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			p = next
			visited.Add(p)
		}
	}

	fmt.Println(visited.Len())
}

func Day6PartTwo(filename string) {
	m, p, dir := parseMap(filename)

	var count int
	for i := 0; i < Size; i++ {
		for j := 0; j < Size; j++ {
			if testObstacle(m, P{i, j}, p, dir) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func testObstacle(m map[P]struct{}, obstacle, pos P, dir uint8) bool {

	// impossible if the guard is there
	if obstacle == pos {
		return false
	}

	// impossible if there is already an obstacle there
	if _, ok := m[obstacle]; ok {
		return false
	}

	m[obstacle] = struct{}{}
	defer func() {
		delete(m, obstacle)
	}()

	visited := set.Set[State]{}
	visited.Add(State{pos, dir})

	for {
		var next P

		switch dir {
		case N:
			next = P{pos.x - 1, pos.y}
		case E:
			next = P{pos.x, pos.y + 1}
		case S:
			next = P{pos.x + 1, pos.y}
		case W:
			next = P{pos.x, pos.y - 1}
		}

		if next.x < 0 || next.x >= Size || next.y < 0 || next.y >= Size {
			return false
		}

		if _, ok := m[next]; ok {
			dir = (dir + 1) % 4
		} else {
			pos = next
		}

		newState := State{pos, dir}
		if visited.Contains(newState) {
			return true
		}
		visited.Add(newState)
	}
}
