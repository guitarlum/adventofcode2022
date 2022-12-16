package main

import (
	"sort"
	"strconv"
	"strings"
)

type Pipe struct {
	name     string
	pressure int
	leadsTo  []string
}

type FromTo struct {
	from string
	to   string
}

type Action struct {
	pressureGain int
	distance     int
	targetPipe   Pipe
}

type ByPressure []Pipe

func (l ByPressure) Len() int           { return len(l) }
func (l ByPressure) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l ByPressure) Less(i, j int) bool { return l[i].pressure > l[j].pressure }

func calcPressure(lines []string) int {
	sum := 0
	pipes := extractPipes(lines)

	distances := make(map[FromTo]int)
	for _, f := range pipes {
		if f.pressure == 0 && f.name != "AA" {
			continue
		}
		from := f.name
		for _, t := range pipes {
			if t.pressure == 0 {
				continue
			}
			if f.name != t.name {
				shortestPath := 9999
				var history []string
				findShortestDistance(from, from, t.name, pipes, &shortestPath, 0, history)
				distances[FromTo{from, t.name}] = shortestPath
			}
		}
	}

	var valuePipes []Pipe
	for _, p := range pipes {
		if p.pressure != 0 {
			valuePipes = append(valuePipes, p)
		}
	}
	sort.Sort(ByPressure(valuePipes))

	currentPipe := "AA"
	// for i := 0; i < 30; i++ {
	// 	highestNextPipe := findHighest(currentPipe, distances, valuePipes, 30-i)
	// 	if highestNextPipe.pressureGain != 0 {
	// 		sum += highestNextPipe.pressureGain
	// 		i += highestNextPipe.distance + 1
	// 		currentPipe = highestNextPipe.targetPipe.name
	// 		remove(valuePipes, highestNextPipe.targetPipe)
	// 	} else {
	// 		break
	// 	}
	// }
	var pipesList []Pipe
	pipesList = append(pipesList, pipes[currentPipe])
	bruteForceIt(&sum, pipes[currentPipe], valuePipes[:], distances, 0, 0, pipesList[:])

	return sum
}

func bruteForceIt(highest *int, current Pipe, valuePipes []Pipe, distances map[FromTo]int, i int, sum int, pipesList []Pipe) {
	timeLeft := 30 - i

	if sum > *highest {
		*highest = sum
	}

	if timeLeft <= 0 {
		return
	}

	for _, p := range valuePipes {
		if !containsPipe(pipesList, p) {
			distance := distances[FromTo{current.name, p.name}]
			if (distance + 1) > timeLeft {
				continue
			}
			bruteForceIt(highest, p, valuePipes, distances, i+distance+1, sum+((timeLeft-distance-1)*p.pressure), append(pipesList, p)[:])
		}
	}
}

func remove(slice []Pipe, s int) []Pipe {
	return append(slice[:s], slice[s+1:]...)
}

func findHighest(currentPipe string, distances map[FromTo]int, valuePipes []Pipe, timeLeft int) Action {
	var bestMove Action
	for _, p := range valuePipes {
		distance := distances[FromTo{currentPipe, p.name}]
		if (distance + 1) > timeLeft {
			continue
		}
		total := (timeLeft - distance - 1) * p.pressure
		if total > bestMove.pressureGain {
			bestMove.pressureGain = total
			bestMove.targetPipe = p
			bestMove.distance = distance
		}
	}
	return bestMove
}

// func remove(s []Pipe, r Pipe) []Pipe {
// 	for i, v := range s {
// 		if v.name == r.name {
// 			return append(s[:i], s[i+1:]...)
// 		}
// 	}
// 	return s
// }

func findShortestDistance(start, from, to string, pipes map[string]Pipe, shortest *int, steps int, history []string) {
	if contains(history, from) {
		return
	}
	if contains(pipes[from].leadsTo, to) {
		if steps < *shortest {
			*shortest = steps + 1
		}
		return
	}
	for _, v := range pipes[from].leadsTo {
		history = append(history, from)
		findShortestDistance(start, v, to, pipes, shortest, steps+1, history[:])
	}
}

func containsPipe(s []Pipe, e Pipe) bool {
	for _, a := range s {
		if a.name == e.name {
			return true
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func extractPipes(lines []string) map[string]Pipe {
	pipes := make(map[string]Pipe)
	for _, v := range lines {
		parts := strings.Split(v, " ")
		name := parts[1]
		pressurePart := strings.Split(parts[4], "=")
		pressure, _ := strconv.Atoi(pressurePart[1][:len(pressurePart[1])-1])

		var leadsTo []string
		for i := 9; i < len(parts); i++ {
			if i == len(parts)-1 {
				leadsTo = append(leadsTo, parts[i])
			} else {
				leadsTo = append(leadsTo, parts[i][:2])
			}
		}
		pipes[name] = Pipe{name, pressure, leadsTo}
	}
	return pipes
}
