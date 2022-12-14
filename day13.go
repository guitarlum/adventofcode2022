package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Leaf struct {
	index  int
	value  int
	parent *Leaf
	sub    []*Leaf
}

type ByOrder []Leaf

func (l ByOrder) Len() int      { return len(l) }
func (l ByOrder) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l ByOrder) Less(i, j int) bool {
	order := Continue
	compareLeafes(&l[i], &l[j], &order)
	return order == RightOrder
}

type Order int

const (
	Continue Order = iota
	WrongOrder
	RightOrder
)

func calcSumOfIndicese(lines []string) (int, int) {
	sum := 0
	var leafs []Leaf

	for i, c := 0, 1; i < len(lines); i, c = i+3, c+1 {
		leftLeaf := extractLeaf(lines[i], false)
		rightLeaf := extractLeaf(lines[i+1], false)
		leafs = append(leafs, *leftLeaf)
		leafs = append(leafs, *rightLeaf)
		order := Continue
		compareLeafes(leftLeaf, rightLeaf, &order)
		if order == RightOrder {
			sum += c
		}
	}

	add := []string{"[2]", "[6]"}
	leafs = append(leafs, *extractLeaf(add[0], true))
	leafs = append(leafs, *extractLeaf(add[1], true))

	sort.Sort(ByOrder(leafs))

	decoderKey := 1
	for i, l := range leafs {
		for _, s := range l.sub {
			if s.index == -666 {
				decoderKey *= (i + 1)
			}
		}
	}

	return sum, decoderKey
}

func compareLeafes(l, r *Leaf, order *Order) {
	if l.value != -1 && r.value == -1 {
		makeList(l)
		compareLeafes(l, r, order)
	} else if l.value == -1 && r.value != -1 {
		makeList(r)
		compareLeafes(l, r, order)
	}

	if l.value == -1 && r.value == -1 {
		for i := 0; i < min(len(l.sub), len(r.sub)); i++ {
			compareLeafes(l.sub[i], r.sub[i], order)
			if *order != Continue {
				return
			}
		}
		if len(l.sub) > len(r.sub) {
			*order = WrongOrder
			return
		} else if len(l.sub) == len(r.sub) {
			*order = Continue
		} else {
			*order = RightOrder
			return
		}
	} else {
		if l.value > r.value {
			*order = WrongOrder
		} else if l.value == r.value {
			*order = Continue
		} else {
			*order = RightOrder
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func makeList(l *Leaf) {
	n := new(Leaf)
	n.value = l.value
	n.index = l.index
	n.parent = l
	l.value = -1
	l.sub = append(l.sub, n)
}

func printLeaf(l *Leaf) {
	if l.value == -1 {
		fmt.Print("[")
		for i, v := range l.sub {
			if v.value == -1 {
				printLeaf(v)
			} else {
				fmt.Print(v.value)
				if i < len(l.sub)-1 {
					fmt.Print(",")
				}
			}
		}
		fmt.Print("]")
	}
	if l.parent == nil {
		fmt.Println()
	}
}

func extractLeaf(left string, find bool) *Leaf {
	var currentLeaf *Leaf

	for x := 0; x < len(left); x++ {
		switch left[x] {
		case ',':
			continue
		case '[':
			newL := new(Leaf)
			newL.index = x
			newL.value = -1
			if currentLeaf != nil {
				newL.parent = currentLeaf
				currentLeaf.sub = append(currentLeaf.sub, newL)
			}
			currentLeaf = newL
		case ']':
			if currentLeaf.parent != nil {
				currentLeaf = currentLeaf.parent
			}
		default:
			var number int
			if left[x+1] == '[' || left[x+1] == ']' || left[x+1] == ',' {
				number, _ = strconv.Atoi(string(left[x]))
			} else {
				number, _ = strconv.Atoi(string(left[x:(x + 2)]))
				x++
			}
			newL := new(Leaf)
			newL.index = x
			newL.parent = currentLeaf
			newL.value = number
			if find {
				newL.index = -666
			}
			newL.parent.sub = append(newL.parent.sub, newL)
		}
	}

	for currentLeaf.parent != nil {
		currentLeaf = currentLeaf.parent
	}

	return currentLeaf
}
