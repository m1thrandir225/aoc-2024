package day1

import (
	"aoc_2024/util"
	"bufio"
	"fmt"
	"os"
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
	firstList = util.SortByLowest(firstList)
	secondList = util.SortByLowest(secondList)
	for i := range firstList {
		var distance int
		if firstList[i] > secondList[i] {
			distance = firstList[i] - secondList[i]
		} else {
			distance = secondList[i] - firstList[i]
		}
		differences = append(differences, distance)
	}

	sum := util.SumItems(differences)

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
	firstList = util.SortByLowest(firstList)
	secondList = util.SortByLowest(secondList)

	for i := range firstList {
		current := firstList[i]
		occurrence := util.CheckItemOccurrence(secondList, current)

		simillarity := current * occurrence
		listSimillarity = append(listSimillarity, simillarity)
	}

	similarityScore := util.SumItems(listSimillarity)
	fmt.Println(similarityScore)
}
