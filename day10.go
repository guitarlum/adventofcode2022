package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func runCycles(lines []string, cycles []int, w http.ResponseWriter) int {
	sum := 0
	cycleCounter := 1
	frequency := 0
	instruction := make(map[int]int)
	x := 1
	readNext := true

	for i := 0; cycleCounter <= 240; cycleCounter++ {

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

		currentPosition := cycleCounter%40 - 1
		if x >= currentPosition-1 && x <= currentPosition+1 {
			fmt.Fprintf(w, "#")
		} else {
			fmt.Fprintf(w, ".")
		}

		if cycleCounter <= cycles[len(cycles)-1] && cycleCounter == cycles[frequency] {
			sum += x * cycles[frequency]
			frequency++
		}

		if instruction[cycleCounter] != 0 {
			x += instruction[cycleCounter]
			i++
			readNext = true
		}

		if cycleCounter%40 == 0 {
			fmt.Fprintln(w)
		}
	}

	return sum
}
