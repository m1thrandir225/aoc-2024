package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7PartOne(filename string) {

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	totalSum := 0

	for scanner.Scan() {
		numbers, result := parseLine(scanner.Text())
		operatorCombos := generateOperatorCombinations([]string{"+", "*"}, len(numbers)-1)
		for _, operators := range operatorCombos {
			if evaluate(numbers, operators, result) == result {
				totalSum += result
				break
			}
		}
	}
	fmt.Println(totalSum)
}

func Day7PartTwo(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	totalSum := 0
	for scanner.Scan() {
		numbers, result := parseLine(scanner.Text())
		operatorCombos := generateOperatorCombinations([]string{"+", "*", "||"}, len(numbers)-1)

		for _, operators := range operatorCombos {
			if evaluate(numbers, operators, result) == result {
				totalSum += result
				break
			}
		}
	}
	fmt.Println(totalSum)
}

func parseLine(line string) ([]int, int) {
	parts := strings.Fields(line)
	result, _ := strconv.Atoi(strings.TrimSuffix(parts[0], ":"))
	numbers := make([]int, len(parts)-1)
	for i, num := range parts[1:] {
		numbers[i], _ = strconv.Atoi(num)
	}
	return numbers, result
}

func generateOperatorCombinations(operators []string, length int) [][]string {
	var combinations [][]string
	var generate func(current []string)

	generate = func(current []string) {
		if len(current) == length {
			combinations = append(combinations, append([]string(nil), current...))
			return
		}

		for _, operator := range operators {
			current = append(current, operator)
			generate(current)
			current = current[:len(current)-1]
		}
	}
	generate([]string{})
	return combinations

}

func evaluate(numbers []int, operators []string, solution int) int {
	result := numbers[0]
	for i, op := range operators {
		switch op {
		case "+":
			result += numbers[i+1]
		case "*":
			result *= numbers[i+1]
		case "||":
			result = concatenate(result, numbers[i+1])
		}

		if result > solution {
			return -1
		}
	}
	return result
}

func pow(base, exponent int) int {
	result := 1
	for exponent > 0 {
		result *= base
		exponent--
	}
	return result
}

func concatenate(a, b int) int {
	digits := 0
	temp := b
	for temp > 0 {
		temp /= 10
		digits++
	}
	return a*pow(10, digits) + b
}
