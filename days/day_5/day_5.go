package day5

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Day5PartOne(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	string := string(b)
	parts := strings.Split(string, "\n\n")
	pageRules := getPageRules(parts[0])
	// for k, v := range pageRules {
	// 	fmt.Println(k, "value is", v)
	// }
	pages := getPages(parts[1])
	coutner := 0
	for _, page := range pages {
		if pagesOrderedCorrectly(page, pageRules) {
			middle := getMiddlePart(page)
			coutner += middle
		}
	}
	fmt.Println(coutner)

}

func Day5PartTwo(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	string := string(b)
	parts := strings.Split(string, "\n\n")
	pageRules := getPageRules(parts[0])
	// for k, v := range pageRules {
	// 	fmt.Println(k, "value is", v)
	// }
	pages := getPages(parts[1])
	coutner := 0
	for _, page := range pages {
		if !pagesOrderedCorrectly(page, pageRules) {
			updatedPage := fixPageOrdering(page, pageRules)
			middle := getMiddlePart(updatedPage)
			coutner += middle
		}
	}
	fmt.Println(coutner)
}

func getPages(text string) [][]int {
	lines := strings.Split(text, "\n")
	pages := make([][]int, len(lines))

	for i, line := range lines {
		strNums := strings.Split(line, ",")
		nums := make([]int, 0)
		for _, strNum := range strNums {
			num, _ := strconv.Atoi(strNum)
			nums = append(nums, num)
		}
		pages[i] = nums
	}
	return pages
}

func fixPageOrdering(page []int, rules map[int][]int) []int {
	for pageIndex, pageNum := range page {
		afterList, ok := rules[pageNum]
		if !ok {
			continue
		}

		for i := 0; i < pageIndex; i++ {
			itemBeforeCurrent := page[i]
			if !(isItemPlacedCorrectly(afterList, itemBeforeCurrent)) {
				swapItems(page, i, pageIndex)
			}
		}
	}
	return page
}

func pagesOrderedCorrectly(pages []int, rules map[int][]int) bool {
	for pageIndex, pageNum := range pages {
		afterList, ok := rules[pageNum]
		if !ok {
			continue
		}
		//Check if items before the current one are in the list, that should be after
		for i := 0; i < pageIndex; i++ {
			itemBeforeCurrent := pages[i]
			if !(isItemPlacedCorrectly(afterList, itemBeforeCurrent)) {
				return false
			}
		}
	}
	return true
}

func isItemPlacedCorrectly(list []int, item int) bool {
	for _, v := range list {
		if v == item {
			return false
		}
	}
	return true
}

func getMiddlePart(page []int) int {
	return page[len(page)/2]
}

func swapItems(list []int, indexA, indexB int) {
	temp := list[indexA]

	list[indexA] = list[indexB]
	list[indexB] = temp
}

func getPageRules(text string) map[int][]int {
	lookupTable := make(map[int][]int, 0)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		afterList, ok := lookupTable[before]
		if !ok {
			afterList := make([]int, 0)
			afterList = append(afterList, after)
			lookupTable[before] = afterList
		}
		afterList = append(afterList, after)
		lookupTable[before] = afterList
	}
	return lookupTable
}
