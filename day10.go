package main

import (
	"strconv"
	"strings"
)

func runCycles(lines []string, cycles []int) int {
	var sum int = 0
	var cycleCounter int = 1
	var frequency int = 0
	instruction := make(map[int]int)
	var x int = 1
	readNext := true

	for i := 0; cycleCounter <= cycles[len(cycles)-1]; cycleCounter++ {

		if readNext {
			if i == (len(lines)) {
				i = 0
			}
			parts := strings.Split(lines[i], " ")
			if parts[0] == "addx" {
				value, _ := strconv.Atoi(parts[1])
				instruction[cycleCounter+1] = value
				readNext = false
			} else {
				i++
			}
		}

		if cycleCounter == cycles[frequency] {
			sum += x * cycles[frequency]
			frequency++
		}

		if instruction[cycleCounter] != 0 {
			x += instruction[cycleCounter]
			i++
			readNext = true
		}
	}

	return sum
}
