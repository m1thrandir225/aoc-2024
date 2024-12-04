package day1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day1PartOne(filename string) {
	firstList := make([]int, 0)
	secondList := make([]int, 0)
	differences := make([]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")
		first, _ := strconv.Atoi(items[0])
		second, _ := strconv.Atoi(items[1])
		firstList = append(firstList, first)
		secondList = append(secondList, second)
	}
	sortByLowest(firstList)
	sortByLowest(secondList)
	for i := range firstList {
		var distance int
		if firstList[i] > secondList[i] {
			distance = firstList[i] - secondList[i]
		} else {
			distance = secondList[i] - firstList[i]
		}
		differences = append(differences, distance)
	}

	sum := sumList(differences)

	fmt.Println(sum)

}

func Day1PartTwo(filename string) {
	firstList := make([]int, 0)
	secondList := make([]int, 0)
	listSimillarity := make([]int, 0)

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "   ")
		first, _ := strconv.Atoi(items[0])
		second, _ := strconv.Atoi(items[1])
		firstList = append(firstList, first)
		secondList = append(secondList, second)
	}
	sortByLowest(firstList)
	sortByLowest(secondList)

	for i := range firstList {
		current := firstList[i]
		occurrence := checkOccurrences(secondList, current)

		simillarity := current * occurrence
		listSimillarity = append(listSimillarity, simillarity)
	}

	similarityScore := sumList(listSimillarity)
	fmt.Println(similarityScore)
}

func checkOccurrences(list []int, number int) int {
	occurrence := 0
	for i := range list {
		if list[i] == number {
			occurrence++
		}
	}
	return occurrence
}

func sortByLowest(list []int) []int {
	sortedList := list

	sort.Ints(sortedList)

	return sortedList
}

func sumList(list []int) int {
	sum := 0
	for _, v := range list {
		sum += v
	}
	return sum
}
