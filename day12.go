package main

type coordinate struct {
	x, y int
}

type traversePoint struct {
	left, right, up, down bool
}

var inputMap []string
var traverseMapCache map[coordinate]traversePoint

func findShortestPath(mapInput []string) int {
	inputMap = mapInput
	traverseMapCache = make(map[coordinate]traversePoint)
	steps := 0
	bestSteps := &steps
	start := findStart(mapInput)

	inputMap[start.y] = replaceAtIndex(inputMap[start.y], 'a', start.x)

	prepCache()

	traverseMap(start, bestSteps, 1000)

	return steps
}

func traverseMap(coo coordinate, bestSteps *int, max int) {
	stepCount := make(map[coordinate]int)
	stepCount[coo]++
	var from coordinate = coo
	pointsToTravel := whereTo(stepCount, coo)
	for *bestSteps == 0 {
		newList := make(map[coordinate]bool)
		for c := range pointsToTravel {
			from = c
			if inputMap[c.y][c.x] == 'E' {
				*bestSteps = stepCount[c] + 1
				return
			}

			tempList := whereTo(stepCount, c)
			for v := range tempList {
				stepCount[v] = stepCount[from] + 1
				newList[v] = true
			}
		}

		pointsToTravel = newList
	}
}

func whereTo(past map[coordinate]int, coo coordinate) map[coordinate]bool {
	pointsToTravel := make(map[coordinate]bool)

	left := coordinate{coo.x - 1, coo.y}
	if traverseMapCache[coo].left && past[left] < 1 {
		pointsToTravel[left] = true
	}

	right := coordinate{coo.x + 1, coo.y}
	if traverseMapCache[coo].right && past[right] < 1 {
		pointsToTravel[right] = true
	}

	up := coordinate{coo.x, coo.y - 1}
	if traverseMapCache[coo].up && past[up] < 1 {
		pointsToTravel[up] = true
	}

	down := coordinate{coo.x, coo.y + 1}
	if traverseMapCache[coo].down && past[down] < 1 {
		pointsToTravel[down] = true
	}

	return pointsToTravel
}

func prepCache() {
	for y, row := range inputMap {
		for x := 0; x < len(row); x++ {
			right := coordinate{x + 1, y}
			left := coordinate{x - 1, y}
			up := coordinate{x, y - 1}
			down := coordinate{x, y + 1}
			coo := coordinate{x, y}

			traverseMapCache[coo] = traversePoint{isInDistance(coo, left), isInDistance(coo, right), isInDistance(coo, up), isInDistance(coo, down)}
		}
	}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func findStart(mapInput []string) coordinate {
	for y, row := range mapInput {
		for x := 0; x < len(row); x++ {
			if row[x] == 'S' {
				return coordinate{x, y}
			}
		}
	}
	return coordinate{0, 0}
}

func isInDistance(from, to coordinate) bool {
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
