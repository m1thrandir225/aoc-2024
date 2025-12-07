package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day9PartOne(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	array := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		for _, str := range parts {
			num, _ := strconv.Atoi(str)
			array = append(array, num)
		}
	}
	filesArray := make([]int, 0)
	spaceArray := make([]int, 0)
	for i := 0; i < len(array); i++ {
		if i%2 == 0 {
			filesArray = append(filesArray, array[i])
		} else {
			spaceArray = append(spaceArray, array[i])
		}
	}

	fmt.Println(filesArray)
	fmt.Println(spaceArray)
}

func makeRow(files, space []int) string {
	if len(files) != len(space) {
		return ""
	}
	for i := 0; i < len(files); i++ {
	}

	return ""
}
func Day9PartTwo(filename string) {}
