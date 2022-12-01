package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		var lines []int = readFileAsNumbers("input/day1.txt")
		fmt.Fprintln(w, "highest: ", calcCalories(lines))
		fmt.Fprintln(w, "top3: ", calcCaloriesTop3(lines))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
