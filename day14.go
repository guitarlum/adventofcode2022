package main

import (
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func calcSandAtRest(lines []string, floor bool) int {
	cave, maxDepth := createCave(lines, floor)
	landed := putSand(&cave, maxDepth)
	return landed
}

func putSand(cave *[200][700]int, maxDepth int) int {
	landed := 0

nextSand:
	for {
		sand := Point{500, 0}
	keepFalling:
		for {
			if sand.y >= maxDepth {
				return landed
			}
			switch cave[sand.y+1][sand.x] {
			case 0:
				sand.y++
			default:
				if cave[sand.y+1][sand.x-1] != 0 {
					if cave[sand.y+1][sand.x+1] != 0 {
						cave[sand.y][sand.x] = 2
						landed++
						if sand.y == 0 && sand.x == 500 {
							return landed
						}
						goto nextSand
					} else {
						sand.x++
						goto keepFalling
					}
				} else {
					sand.x--
					goto keepFalling
				}
			}
		}
	}
}

func createCave(lines []string, floor bool) ([200][700]int, int) {
	cave := [200][700]int{}
	maxDepth := 0

	for _, v := range lines {
		parts := strings.Split(v, " -> ")
		points := extractPoints(parts)
		for i := 0; i < len(points); i++ {
			if points[i].y > maxDepth {
				maxDepth = points[i].y
			}
			if i == len(points)-1 {
				break
			}
			fillStones(&cave, points[i], points[i+1])
		}
	}

	if floor {
		for x := 0; x < len(cave[maxDepth+2]); x++ {
			cave[maxDepth+2][x] = 1
		}
		maxDepth += 2
	}

	return cave, maxDepth
}

func fillStones(cave *[200][700]int, from, to Point) {
	xDiff := from.x - to.x
	yDiff := from.y - to.y

	if xDiff != 0 {
		for x := from.x; x != to.x; {
			cave[from.y][x] = 1
			if xDiff > 0 {
				x--
			} else {
				x++
			}
		}
	} else if yDiff != 0 {
		for y := from.y; y != to.y; {
			cave[y][from.x] = 1
			if yDiff > 0 {
				y--
			} else {
				y++
			}
		}
	}
	cave[to.y][to.x] = 1
}

func extractPoints(parts []string) []Point {
	var points []Point
	for _, v := range parts {
		xy := strings.Split(v, ",")
		x, _ := strconv.Atoi(xy[0])
		y, _ := strconv.Atoi(xy[1])
		points = append(points, Point{x, y})
	}
	return points
}
