package main

import (
	"fmt"
	"strconv"
	"strings"
)

var yToLook int = 2000000

type SBPair struct {
	sensor, beacon Point
}

type Range struct {
	from, to   int
	isDistress bool
}

func calcDistress(lines []string) int {
	ranges := extractRanges(lines)

	// count := countRanges(ranges)
	// fmt.Println(cMap)
	count := 0
	for _, v := range ranges {
		if v == 1 {
			count++
		}
	}
	return count
}

func countRanges(ranges []Range) int {
	sum := 0

	return sum
}

func extractRanges(lines []string) map[int]int {
	completeM := make(map[Point]byte)

	// var ranges []Range
	dMap := make(map[int]int)
	var pairs []SBPair
	for _, l := range lines {
		parts := strings.Split(l, " ")

		sxPart := strings.Split(parts[2], "=")
		sx, _ := strconv.Atoi(sxPart[1][:len((sxPart[1]))-1])
		syPart := strings.Split(parts[3], "=")
		sy, _ := strconv.Atoi(syPart[1][:len((syPart[1]))-1])
		s := Point{sx, sy}

		completeM[s] = 'S'
		if s.y == yToLook {
			dMap[s.x] = '2'
			// ranges = append(ranges, Range{s.x, s.x, false})
		}

		bxPart := strings.Split(parts[8], "=")
		bx, _ := strconv.Atoi(bxPart[1][:len((bxPart[1]))-1])
		byPart := strings.Split(parts[9], "=")
		by, _ := strconv.Atoi(byPart[1])
		b := Point{bx, by}

		completeM[b] = 'B'
		if b.y == yToLook {
			dMap[b.x] = '2'
			// ranges = append(ranges, Range{b.x, b.x, false})
		}
		pairs = append(pairs, SBPair{s, b})
	}

	// pairs = nil
	// pairs = append(pairs, SBPair{Point{8, 7}, Point{2, 10}})
	for _, p := range pairs {
		distance := absInt(p.sensor.x-p.beacon.x) + absInt(p.sensor.y-p.beacon.y)

		startP := p.sensor

		// fmt.Print("from ", startP.y-distance)
		// fmt.Print("to ", startP.y+distance)
		// fmt.Print("and its in range: ", startP.y+distance >= 2000000 && startP.y-distance <= 2000000)
		if startP.y+distance >= yToLook && startP.y-distance <= yToLook {
			fmt.Println("impact in range!")
			span := distance - absInt(p.sensor.y-yToLook)
			xFrom := p.sensor.x - span
			xTo := p.sensor.x + span

			for i := xFrom; i <= xTo; i++ {
				if dMap[i] == 0 {
					dMap[i] = 1
				}
			}

		}
		// fmt.Println()
		// if that distance can go to y = 2000000, calc the new range in there
		// check overlap with existing ranges

		// for d := 0; d <= distance; d++ {
		// 	for x := 0; x <= distance-d; x++ {
		// 		fillDistress(startP.x-x, startP.y+d, &completeM, &distressM)
		// 		fillDistress(startP.x+x, startP.y+d, &completeM, &distressM)
		// 		fillDistress(startP.x-x, startP.y-d, &completeM, &distressM)
		// 		fillDistress(startP.x+x, startP.y-d, &completeM, &distressM)
		// 	}
		// }
	}

	// for y := -10; y < 24; y++ {
	// 	fmt.Print(y)
	// 	for x := -8; x < 30; x++ {
	// 		switch completeM[Point{x, y}] {
	// 		case 0:
	// 			fmt.Print(".")
	// 		default:
	// 			fmt.Print(string(completeM[Point{x, y}]))
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	return dMap
}

func fillDistress(x, y int, completeM *map[Point]byte, distressM *map[int]int) {
	if y == 2000000 {
		toPut := Point{x, y}
		if (*completeM)[toPut] == 0 {
			(*completeM)[toPut] = '#'
			(*distressM)[toPut.y]++
		}
	}
}

func absInt(x int) int {
	return absDiffInt(x, 0)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
