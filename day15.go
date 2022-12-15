package main

import (
	"sort"
	"strconv"
	"strings"
)

type SBPair struct {
	sensor, beacon Point
}

type Range struct {
	from, to   int
	isDistress bool
}

type ByRange []Range

func (l ByRange) Len() int           { return len(l) }
func (l ByRange) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByRange) Less(i, j int) bool { return l[i].from < l[j].from }

func calcDistress(lines []string) int {
	yToLook := 2000000
	dMap := make(map[int]int)
	var pairs []SBPair

	for _, l := range lines {
		s, b := extractSB(l)
		if s.y == yToLook {
			dMap[s.x] = '2'
		}
		if b.y == yToLook {
			dMap[b.x] = '2'
		}
		pairs = append(pairs, SBPair{s, b})
	}

	for _, p := range pairs {
		distance := calcManhattenD(p)

		if p.sensor.y+distance >= yToLook && p.sensor.y-distance <= yToLook {
			span := distance - absInt(p.sensor.y-yToLook)
			xFrom := p.sensor.x - span
			xTo := p.sensor.x + span

			for i := xFrom; i <= xTo; i++ {
				if dMap[i] == 0 {
					dMap[i] = 1
				}
			}
		}
	}

	return countDistress(dMap)
}

func calcManhattenD(p SBPair) int {
	distance := absInt(p.sensor.x-p.beacon.x) + absInt(p.sensor.y-p.beacon.y)
	return distance
}

func countDistress(dMap map[int]int) int {
	count := 0
	for _, v := range dMap {
		if v == 1 {
			count++
		}
	}
	return count
}

func findMissingBeacon(lines []string) int {
	var pairs []SBPair
	for _, l := range lines {
		s, b := extractSB(l)
		pairs = append(pairs, SBPair{s, b})
	}

	sum := 0
	for y := 0; y <= 4000000; y++ {
		go findX(pairs, y, &sum)
	}
	return sum
}

func findX(pairs []SBPair, y int, sum *int) {
	var distress []Range
	for _, p := range pairs {
		distance := calcManhattenD(p)

		if p.sensor.y+distance >= y && p.sensor.y-distance <= y {
			span := distance - absInt(p.sensor.y-y)
			xFrom := p.sensor.x - span
			xTo := p.sensor.x + span
			distress = append(distress, Range{xFrom, xTo, true})
		}
	}
	sort.Sort(ByRange(distress))
	for i, upper := 0, distress[0].to; i < len(distress)-1; i++ {
		if upper < distress[i].to {
			upper = distress[i].to
		}
		if upper < distress[i+1].from {
			*sum = ((upper + 1) * 4000000) + y
		}
	}
}

func extractSB(l string) (Point, Point) {
	parts := strings.Split(l, " ")

	sxPart := strings.Split(parts[2], "=")
	sx, _ := strconv.Atoi(sxPart[1][:len((sxPart[1]))-1])
	syPart := strings.Split(parts[3], "=")
	sy, _ := strconv.Atoi(syPart[1][:len((syPart[1]))-1])
	s := Point{sx, sy}

	bxPart := strings.Split(parts[8], "=")
	bx, _ := strconv.Atoi(bxPart[1][:len((bxPart[1]))-1])
	byPart := strings.Split(parts[9], "=")
	by, _ := strconv.Atoi(byPart[1])
	b := Point{bx, by}
	return s, b
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
