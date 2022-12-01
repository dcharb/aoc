package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Arg validation
	if len(os.Args) != 3 {
		log.Fatal("invalid args")
	}
	year := os.Args[1]
	test := os.Args[2]

	// Setup test functions
	funcMap := map[string]func([]string){
		"20201": do20201,
		"20202": do20202,
		"20203": do20203,
		"20204": do20204,
		"20205": do20205,
		"20206": do20206,
		"20211": do20211,
		"20212": do20212,
		"20213": do20213,
		"20221": do20221,
	}

	// Get input and run
	lines := getLines(fmt.Sprintf("data/%s/%s.txt", year, test))
	f, ok := funcMap[fmt.Sprintf("%s%s", year, test)]
	if !ok {
		log.Fatal("invalid test")
	}
	f(lines)
}

func getLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := strings.Split(string(b), "\n")
	return s[:len(s)-1]
}

func toInts(nums []string) []int {
	var ints []int
	for _, num := range nums {
		i, err := strconv.Atoi(num)
		if err != nil {
			log.Panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
