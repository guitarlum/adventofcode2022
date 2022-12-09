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

func calcMapCoverage(lines []string, length int) int {
	visitedMap := make(map[point]int)
	rope := make([]point, length+1)

	visitedMap[rope[length]]++

	for _, v := range lines {
		parts := strings.Split(v, " ")
		distance, _ := strconv.Atoi(parts[1])
		for i := 0; i < distance; i++ {
			switch parts[0] {
			case "U":
				rope[0].y++
			case "D":
				rope[0].y--
			case "L":
				rope[0].x--
			case "R":
				rope[0].x++
			}
			for j := 0; j < length; j++ {
				if isTooFar(rope[j], rope[j+1]) {
					if rope[j].y-rope[j+1].y > 0 {
						rope[j+1].y++
					} else if rope[j].y-rope[j+1].y < 0 {
						rope[j+1].y--
					}
					if rope[j].x-rope[j+1].x > 0 {
						rope[j+1].x++
					} else if rope[j].x-rope[j+1].x < 0 {
						rope[j+1].x--
					}
				}
			}
			visitedMap[rope[length]]++
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
