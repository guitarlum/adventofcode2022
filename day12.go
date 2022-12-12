package main

type coordinate struct {
	x, y int8
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
	var visited []coordinate

	inputMap[start.y] = replaceAtIndex(inputMap[start.y], 'a'-1, start.x)

	prepCache()

	// for i := 0; steps == 0 && i < 1000; i++ {
	// fmt.Println("Starting for depth ", i)
	traverseMap(visited, start, bestSteps, 1000)
	// }

	return steps
}

func traverseMap(past []coordinate, coo coordinate, bestSteps *int, max int) {
	// for ;*bestSteps == 0; {


	// }
	past = append(past, coo)

	if len(past) > max {
		return
	}

	if  && *bestSteps < len(past)-1 {
		return
	}
	if inputMap[coo.y][coo.x] == 'E' {
		*bestSteps = len(past) - 1
		return
	}

	left := coordinate{coo.x - 1, coo.y}
	if traverseMapCache[coo].left && !contains(past, left) {
		traverseMap(past, left, bestSteps, max)
	}

	right := coordinate{coo.x + 1, coo.y}
	if traverseMapCache[coo].right && !contains(past, right) {
		traverseMap(past, right, bestSteps, max)
	}

	up := coordinate{coo.x, coo.y - 1}
	if traverseMapCache[coo].up && !contains(past, up) {
		traverseMap(past, up, bestSteps, max)
	}

	down := coordinate{coo.x, coo.y + 1}
	if traverseMapCache[coo].down && !contains(past, down) {
		traverseMap(past, down, bestSteps, max)
	}
}

func prepCache() {
	for y, row := range inputMap {
		for x := 0; x < len(row); x++ {

			right := coordinate{int8(x) + 1, int8(y)}
			left := coordinate{int8(x) - 1, int8(y)}
			up := coordinate{int8(x), int8(y) - 1}
			down := coordinate{int8(x), int8(y) + 1}
			coo := coordinate{int8(x), int8(y)}

			traverseMapCache[coo] = traversePoint{isInDistance(coo, left), isInDistance(coo, right), isInDistance(coo, up), isInDistance(coo, down)}
		}
	}
}

func replaceAtIndex(in string, r rune, i int8) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func findStart(mapInput []string) coordinate {
	for y, row := range mapInput {
		for x := 0; x < len(row); x++ {
			if row[x] == 'S' {
				return coordinate{int8(x), int8(y)}
			}
		}
	}
	return coordinate{0, 0}
}

func isInDistance(from, to coordinate) bool {
	if to.x < 0 || to.y < 0 || to.y >= int8(len(inputMap)) || to.x >= int8(len(inputMap[to.y])) {
		return false
	}
	diff := inputMap[to.y][to.x] - inputMap[from.y][from.x]
	if diff <= 1 || diff == 255 || (inputMap[from.y][from.x] == 'z' && inputMap[to.y][to.x] == 'E') {
		return true
	} else {
		return false
	}
}

func copyMap(from map[coordinate]bool) map[coordinate]bool {
	newMap := make(map[coordinate]bool)
	for k, v := range from {
		newMap[k] = v
	}
	return newMap
}

func contains(s []coordinate, e coordinate) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
