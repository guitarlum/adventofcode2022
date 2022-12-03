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

	log.Fatal(http.ListenAndServe(":8081", nil))
}
