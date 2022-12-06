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
		fmt.Fprintln(w, "4 char package: ", FindStartOfPacket(lines[0], 4))
		fmt.Fprintln(w, "14 char package: ", FindStartOfPacket(lines[0], 14))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
