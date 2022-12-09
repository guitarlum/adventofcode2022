package main

import (
	"math"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func calcMapCoverage(lines []string) int {
	visitedMap := make(map[point]int)
	var h point = point{0, 0}
	var t point = point{0, 0}

	visitedMap[t]++

	for _, v := range lines {
		parts := strings.Split(v, " ")
		distance, _ := strconv.Atoi(parts[1])
		for i := 0; i < distance; i++ {
			switch parts[0] {
			case "U":
				h.y++
				if isTooFar(h, t) {
					t.x = h.x
					t.y = h.y - 1
				}
			case "D":
				h.y--
				if isTooFar(h, t) {
					t.x = h.x
					t.y = h.y + 1
				}
			case "L":
				h.x--
				if isTooFar(h, t) {
					t.y = h.y
					t.x = h.x + 1
				}
			case "R":
				h.x++
				if isTooFar(h, t) {
					t.y = h.y
					t.x = h.x - 1
				}
			}
			visitedMap[t]++
		}
	}
	return len(visitedMap)
}

func isTooFar(h point, t point) bool {
	if math.Abs(float64(h.x-t.x)) > 1 || math.Abs(float64(h.y-t.y)) > 1 {
		return true
	}
	return false
}
