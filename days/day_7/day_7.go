package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day7PartOne(filename string) {
	table := make(map[int][]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		value, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		possibilityListStr := strings.Split(parts[1], " ")
		possibilityList := make([]int, 0)
		for _, item := range possibilityListStr {
			numberValueItem, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			possibilityList = append(possibilityList, numberValueItem)
		}
		table[value] = possibilityList
	}

	for target, numbersToCombine := range table {
		possibleOperators := len(numbersToCombine) - 1

	}
}

func evaluate(nums []int, operatorLenght int, operator string) int {
	result := nums[0]
	for i := 0; i < operatorLenght; i++ {
		if operator == "+" {
			result += nums[i+1]
		} else if operator == "*" {
			result *= nums[i+1]
		}
	}
	return result
}

func matchTarget(nums []int, target int, operatorsLenght int, index int, calculated map[int]bool) {
	if index == len(nums) {
		result := evaluate(nums, , operator string)
	}
}

func Day7PartTwo(filename string) {}
