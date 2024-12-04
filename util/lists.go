package util

import "sort"

func SortByLowest(list []int) []int {
	sorted := list

	sort.Ints(list)

	return sorted
}

func SumItems(list []int) int {
	sum := 0

	for _, v := range list {
		sum += v
	}

	return sum
}

func CheckItemOccurrence(list []int, item int) int {
	occurrence := 0
	for _, v := range list {
		if v == item {
			occurrence++
		}
	}
	return occurrence
}

func IsListIncreasing(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] >= list[i+1] {
			return false
		}
	}
	return true
}

func IsListDecreasing(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] <= list[i+1] {
			return false
		}
	}
	return true
}
