package main

import "sort"

func calcCalories(lines []int) int {
	var highest int = 0
	var current int = 0

	for _, v := range lines {
		current += v
		if v == 0 {
			if current > highest {
				highest = current
			}
			current = 0
		}
	}
	return highest
}

func calcCaloriesTop3(lines []int) int {
	var ladder []int
	var current int = 0

	for _, v := range lines {
		current += v
		if v == 0 {
			ladder = append(ladder, current)
			current = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(ladder)))

	var sum = 0
	for i := 0; i < 3; i++ {
		sum += ladder[i]
	}
	return sum
}
