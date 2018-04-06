package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var results chan int

func main() {
	results = make(chan int)
	http.HandleFunc("/calc", handleCalc)
	fmt.Println("Listening on localhost:7777/calc")
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func convAndCalc(x string) {
	y, err := strconv.Atoi(x)
	if err != nil {
		panic(err)
	}

	results <- y * 5
}

func handleCalc(w http.ResponseWriter, req *http.Request) {
	int1 := req.FormValue("int1")
	int2 := req.FormValue("int2")

	go convAndCalc(int1)
	go convAndCalc(int2)
	x, y := <-results, <-results
	fmt.Printf("Result1=%d, Result2=%d\n", x, y)
}
