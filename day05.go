package main

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Amount int
	From   int
	To     int
}

type Stack []byte

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(b byte) {
	*s = append(*s, b)
}

func (s *Stack) Pushs(b []byte) {
	*s = append(*s, b...)
}

func (s *Stack) Pop() (byte, bool) {
	if s.IsEmpty() {
		return ' ', false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Pops(a int) ([]byte, bool) {
	if s.IsEmpty() || len(*s) < a {
		return make([]byte, 0), false
	} else {
		index := len(*s)
		element := (*s)[index-a : index]
		*s = (*s)[:index-a]
		return element, true
	}
}

func moveCrates(lines []string, is9001 bool) string {
	crateStrings := lines[:8]
	crates := extractCrates(crateStrings)

	moveStrings := lines[10:]
	instructions := extractInstructions(moveStrings)

	useCrane(instructions, crates, is9001)

	var endCrateOrder string
	for c := 0; c < len(crates); c++ {
		endCrateOrder += string(crates[c][len(crates[c])-1])
	}

	return endCrateOrder

}

func useCrane(instructions []Instruction, crates []Stack, is9001 bool) {
	for _, i := range instructions {
		if is9001 {
			move, _ := crates[i.From-1].Pops(i.Amount)
			crates[i.To-1].Pushs(move)
		} else {
			for a := 0; a < i.Amount; a++ {
				move, _ := crates[i.From-1].Pop()
				crates[i.To-1].Push(move)
			}
		}
	}
}

func extractInstructions(moveStrings []string) []Instruction {
	var instructions []Instruction
	for _, v := range moveStrings {
		instr := strings.Split(v, " ")
		amount, _ := strconv.Atoi(instr[1])
		from, _ := strconv.Atoi(instr[3])
		to, _ := strconv.Atoi(instr[5])
		instructions = append(instructions, Instruction{amount, from, to})
	}
	return instructions
}

func extractCrates(crateStrings []string) []Stack {
	var crates []Stack = make([]Stack, 9)

	for i := len(crateStrings) - 1; i >= 0; i-- {
		for c, j := 0, 1; j < len(crateStrings[i]); c, j = c+1, j+4 {
			if crateStrings[i][j] != ' ' {
				crates[c].Push(crateStrings[i][j])
			}
		}
	}
	return crates
}
