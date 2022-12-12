package main

import (
	"math"
)

type coordinate struct {
	x, y int
}

type traversePoint struct {
	left, right, up, down bool
}

var traverseMapCache map[coordinate]traversePoint

func findShortestPath(mapInput []string, startFrom byte) int {
	traverseMapCache = make(map[coordinate]traversePoint)

	start := findAll('S', mapInput)[0]
	mapInput[start.y] = replaceAtIndex(mapInput[start.y], 'a', start.x)

	prepCache(mapInput)

	var startPoints []coordinate

	if startFrom == 'S' {
		startPoints = append(startPoints, start)
	} else {
		startPoints = findAll(startFrom, mapInput)
	}

	simplyTheBest := math.MaxInt32
	for _, c := range startPoints {
		if current := traverseMap(c, mapInput); current != 0 && current < simplyTheBest {
			simplyTheBest = current
		}
	}

	return simplyTheBest
}

func findAll(find byte, inputMap []string) []coordinate {
	var all []coordinate
	for y, row := range inputMap {
		for x := 0; x < len(row); x++ {
			if row[x] == find {
				all = append(all, coordinate{x, y})
			}
		}
	}
	return all
}

func traverseMap(coo coordinate, inputMap []string) int {
	stepCount := make(map[coordinate]int)
	stepCount[coo]++
	var from coordinate = coo
	pointsToTravel := whereTo(stepCount, coo)
	for len(pointsToTravel) != 0 {
		newList := make(map[coordinate]bool)
		for from = range pointsToTravel {
			if inputMap[from.y][from.x] == 'E' {
				return stepCount[from] + 1
			}

			tempList := whereTo(stepCount, from)
			for v := range tempList {
				stepCount[v] = stepCount[from] + 1
				newList[v] = true
			}
		}
		pointsToTravel = newList
	}
	return 0
}

func whereTo(past map[coordinate]int, coo coordinate) map[coordinate]bool {
	pointsToTravel := make(map[coordinate]bool)

	if left := newC(coo.x-1, coo.y); traverseMapCache[coo].left && past[left] < 1 {
		pointsToTravel[left] = true
	}

	if right := newC(coo.x+1, coo.y); traverseMapCache[coo].right && past[right] < 1 {
		pointsToTravel[right] = true
	}

	if up := newC(coo.x, coo.y-1); traverseMapCache[coo].up && past[up] < 1 {
		pointsToTravel[up] = true
	}

	if down := newC(coo.x, coo.y+1); traverseMapCache[coo].down && past[down] < 1 {
		pointsToTravel[down] = true
	}

	return pointsToTravel
}

func newC(x, y int) coordinate {
	return coordinate{x, y}
}

func prepCache(inputMap []string) {
	for y, row := range inputMap {
		for x := 0; x < len(row); x++ {
			right := coordinate{x + 1, y}
			left := coordinate{x - 1, y}
			up := coordinate{x, y - 1}
			down := coordinate{x, y + 1}
			coo := coordinate{x, y}

			traverseMapCache[coo] = traversePoint{isInDistance(inputMap, coo, left), isInDistance(inputMap, coo, right),
				isInDistance(inputMap, coo, up), isInDistance(inputMap, coo, down)}
		}
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func isInDistance(inputMap []string, from, to coordinate) bool {
	if to.x < 0 || to.y < 0 || to.y >= len(inputMap) || to.x >= len(inputMap[to.y]) {
		return false
	}
	diff := float64(inputMap[to.y][to.x]) - float64(inputMap[from.y][from.x])

	if diff <= 1 || ((inputMap[from.y][from.x] == 'z' || inputMap[from.y][from.x] == 'y') && inputMap[to.y][to.x] == 'E') {
		return true
	} else {
		return false
	}
}
