package day2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Day2PartOne(filename string) {
	lists := make([][]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		itemsInt := make([]int, len(items))
		for i := range items {
			itemInt, _ := strconv.Atoi(items[i])
			itemsInt[i] = itemInt
		}
		lists = append(lists, itemsInt)
	}
	fmt.Println(lists)
	validCounter := 0
	for i := range lists {
		currentList := lists[i]

		diff := isValid(currentList)
		if diff {
			validCounter++
		}
	}
	fmt.Println(validCounter)
}

func Day2PartTwo(filename string) {
	lists := make([][]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		itemsInt := make([]int, len(items))
		for i := range items {
			itemInt, _ := strconv.Atoi(items[i])
			itemsInt[i] = itemInt
		}
		lists = append(lists, itemsInt)
	}

	validCounter := 0
	for i := range lists {
		currentList := lists[i]

		diff := isValid(currentList)
		canDiff := canBecomeValid(currentList)
		if diff || canDiff {
			validCounter++
		}
	}
	fmt.Println(validCounter)
}

func isIncreasing(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			return false
		}
	}
	return true
}

func canBecomeValid(row []int) bool {
	for i := range row {
		//remove current number from the array and check if its a valid row
		newList := make([]int, 0, len(row)-1)
		newList = append(newList, row[:i]...)
		newList = append(newList, row[i+1:]...)

		if isValid(newList) {
			return true
		}
	}
	return false
}

func isDecreasing(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] <= numbers[i+1] {
			return false
		}
	}
	return true
}

func isValid(list []int) bool {
	difference := 0
	listIncreasing := isIncreasing(list)
	listDecreasing := isDecreasing(list)

	if !listIncreasing && !listDecreasing {
		return false
	}

	for i := 0; i < len(list)-1; i++ {
		currentItem := list[i]
		nextItem := list[i+1]

		itemsDiff := int(math.Abs(float64(currentItem - nextItem)))
		if itemsDiff > difference {
			difference = itemsDiff
		}
	}
	return difference >= 1 && difference <= 3
}
