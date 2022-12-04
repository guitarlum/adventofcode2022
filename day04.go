package main

import (
	"strconv"
	"strings"
)

type Elf struct {
	From int
	To   int
}

type Pair struct {
	One Elf
	Two Elf
}

func calcCoverage(lines []string) int {
	var coverage int = 0
	elves := extractElves(lines)

	for _, e := range elves {
		if e.One.From >= e.Two.From && e.One.To <= e.Two.To {
			coverage++
		} else if e.Two.From >= e.One.From && e.Two.To <= e.One.To {
			coverage++
		}
	}
	return coverage
}

func calcCoverageSingle(lines []string) int {
	var coverage int = 0
	elves := extractElves(lines)

	for _, e := range elves {
		if (e.One.From >= e.Two.From && e.One.From <= e.Two.To) ||
			(e.One.To <= e.Two.To && e.One.To >= e.Two.From) {
			coverage++
		} else if (e.Two.From >= e.One.From && e.Two.From <= e.One.To) ||
			(e.Two.To <= e.One.To && e.Two.To >= e.One.From) {
			coverage++
		}
	}
	return coverage
}

func extractElves(lines []string) []Pair {
	var elves []Pair

	for _, v := range lines {
		var pair Pair
		both := strings.Split(v, ",")
		fromTo1 := strings.Split(both[0], "-")
		fromTo2 := strings.Split(both[1], "-")

		pair.One = createElf(fromTo1[0], fromTo1[1])
		pair.Two = createElf(fromTo2[0], fromTo2[1])
		elves = append(elves, pair)
	}

	return elves
}

func createElf(from string, to string) Elf {
	var elf Elf

	elf.From, _ = strconv.Atoi(from)
	elf.To, _ = strconv.Atoi(to)

	return elf
}
