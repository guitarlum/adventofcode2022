package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		var lines []int = readFileAsNumbers("input/day01.txt")
		fmt.Fprintln(w, "highest: ", calcCalories(lines))
		fmt.Fprintln(w, "top3: ", calcCaloriesTop3(lines))
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		var lines []string = readFileAsStringArray("input/day02.txt")
		fmt.Fprintln(w, "score: ", calcRPC(lines, false))
		fmt.Fprintln(w, "score2: ", calcRPC(lines, true))
	})

	http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		var lines []string = readFileAsStringArray("input/day03.txt")
		fmt.Fprintln(w, "score: ", calcRucksack(lines))
		fmt.Fprintln(w, "score of 3: ", calcRucksackOf3(lines))
	})

	http.HandleFunc("/4", func(w http.ResponseWriter, r *http.Request) {
		var lines []string = readFileAsStringArray("input/day04.txt")
		fmt.Fprintln(w, "score: ", calcCoverage(lines))
		fmt.Fprintln(w, "score singles: ", calcCoverageSingle(lines))
	})

	http.HandleFunc("/5", func(w http.ResponseWriter, r *http.Request) {
		var lines []string = readFileAsStringArray("input/day05.txt")
		fmt.Fprintln(w, "crates: ", moveCrates(lines, false))
		fmt.Fprintln(w, "crates after crane 9001: ", moveCrates(lines, true))
	})

	http.HandleFunc("/6", func(w http.ResponseWriter, r *http.Request) {
		var lines []string = readFileAsStringArray("input/day06.txt")
		fmt.Fprintln(w, "4 char package: ", findStartOfPacket(lines[0], 4))
		fmt.Fprintln(w, "14 char package: ", findStartOfPacket(lines[0], 14))
	})

	http.HandleFunc("/7", func(w http.ResponseWriter, r *http.Request) {
		var lines []string = readFileAsStringArray("input/day07.txt")
		fmt.Fprintln(w, "folders sum below 1kk: ", calcDirectorySize(lines))
		fmt.Fprintln(w, "deleteFolder size: ", findDeleteFolder(lines))
	})

	http.HandleFunc("/8", func(w http.ResponseWriter, r *http.Request) {
		treemap := readFileAsIntMap("input/day08.txt")
		trees, scenicScore := findTrees(treemap)
		fmt.Fprintln(w, "visible trees: ", trees)
		fmt.Fprintln(w, "highest scenic score: ", scenicScore)
	})

	http.HandleFunc("/9", func(w http.ResponseWriter, r *http.Request) {
		lines := readFileAsStringArray("input/day09.txt")
		fmt.Fprintln(w, "coverage small snek: ", calcMapCoverage(lines, 1))
		fmt.Fprintln(w, "coverage big snek: ", calcMapCoverage(lines, 9))
	})

	http.HandleFunc("/10", func(w http.ResponseWriter, r *http.Request) {
		lines := readFileAsStringArray("input/day10.txt")
		cycles := [6]int{20, 60, 100, 140, 180, 220}
		fmt.Fprintln(w, "x sum: ", runCycles(lines, cycles[:], w))
	})

	http.HandleFunc("/11", func(w http.ResponseWriter, r *http.Request) {
		lines := readFileAsStringArray("input/day11.txt")
		fmt.Fprintln(w, "monkees after 20: ", chaseMonkyes(lines, 20))
		fmt.Fprintln(w, "monkees after 10000: ", chaseMonkyes(lines, 10000))
	})

	http.HandleFunc("/12", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "shortestPath start from S", findShortestPath(readFileAsStringArray("input/day12.txt"), 'S'))
		fmt.Fprintln(w, "shortestPath start from a", findShortestPath(readFileAsStringArray("input/day12.txt"), 'a'))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
