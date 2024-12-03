package days

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day3PartOne(filename string) {
	//multipliers := make([]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)

	textString := string(b)

	matches := getMatches(textString)
	result := parseMatches(matches)

	fmt.Println(result)
}

func Day3PartTwo(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b, err := io.ReadAll(f)

	textString := string(b)

	matches := getMatchesWithDos(textString)
	result := parseMatches(matches)
	fmt.Println(result)
}

func getMatchesWithDos(s string) []string {
	pattern := `(do\(\))|(don't\(\))|(mul\(\d+,\d+\))`

	re := regexp.MustCompile(pattern)

	matches := re.FindAllString(s, -1)

	var finalResult []string
	enabledMultiplier := true
	for _, match := range matches {
		if match == "do()" {
			enabledMultiplier = true
		} else if match == "don't()" {
			enabledMultiplier = false
		} else if strings.HasPrefix(match, "mul(") {
			if enabledMultiplier {
				finalResult = append(finalResult, match)
			}
		}
	}
	return finalResult
}

func getMatches(s string) []string {
	pattern := `mul\(\d+,\d+\)`

	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(s, -1)

	return matches
}

func parseMatches(list []string) int {
	result := 0
	for i := range list {
		item := list[i]
		pattern := `[0-9]+`
		re := regexp.MustCompile(pattern)
		numbers := re.FindAllString(item, -1)
		first, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}

		result += first * second
	}
	return result
}
