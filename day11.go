package main

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkee struct {
	items            []int
	operation        string
	operationValue   int
	testDivisor      int
	trueMonkeeIndex  int
	falseMonkeeIndex int
	inspectCounter   int
}

func chaseMonkyes(lines []string, rounds int) int {
	monkees, lcm := extractMonkees(lines)

	for i := 0; i < rounds; i++ {
		for mm, m := range monkees {
			for _, item := range m.items {
				switch m.operation {
				case "+":
					item += m.operationValue
				case "*":
					if m.operationValue == 0 {
						item *= item
					} else {
						item *= m.operationValue
					}
				}
				if rounds <= 20 {
					item /= 3
				} else {
					item %= lcm
				}

				if item%m.testDivisor == 0 {
					monkees[m.trueMonkeeIndex].items = append(monkees[m.trueMonkeeIndex].items, item)
				} else {
					monkees[m.falseMonkeeIndex].items = append(monkees[m.falseMonkeeIndex].items, item)
				}
				monkees[mm].inspectCounter++
			}

			monkees[mm].items = nil
		}
	}

	sort.Slice(monkees[:], func(i, j int) bool {
		return monkees[i].inspectCounter < monkees[j].inspectCounter
	})

	return monkees[6].inspectCounter * monkees[7].inspectCounter
}

func extractMonkees(lines []string) ([]Monkee, int) {
	var monkees []Monkee
	lcm := 1
	for i := 0; i < len(lines); i = i + 7 {
		var monkee Monkee

		regex := regexp.MustCompile(`\s\d+`)
		matches := regex.FindAllString(lines[i+1], -1)
		for _, match := range matches {
			item, _ := strconv.Atoi(match[1:])
			monkee.items = append(monkee.items, item)
		}

		opParts := strings.Split(lines[i+2], " ")
		opValue, _ := strconv.Atoi(opParts[len(opParts)-1])
		monkee.operationValue = opValue
		monkee.operation = opParts[len(opParts)-2]

		divisPart := strings.Split(lines[i+3], " ")
		divisor, _ := strconv.Atoi(divisPart[len(divisPart)-1])
		monkee.testDivisor = divisor
		lcm *= divisor

		trueMonkeePart := strings.Split(lines[i+4], " ")
		trueMonkee, _ := strconv.Atoi(trueMonkeePart[len(trueMonkeePart)-1])
		monkee.trueMonkeeIndex = trueMonkee

		falseMonkeePart := strings.Split(lines[i+5], " ")
		falseMonkee, _ := strconv.Atoi(falseMonkeePart[len(falseMonkeePart)-1])
		monkee.falseMonkeeIndex = falseMonkee

		monkees = append(monkees, monkee)
	}
	return monkees, lcm
}
