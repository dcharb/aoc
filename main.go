package aoc

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func day1(filename string) (int, int) {
	p1 := 0
	for _, v := range getInput(filename) {
		p1 += calc(v)
	}

	p2 := 0
	for _, v := range getInput(filename) {
		p2 += getMine(v)
	}

	return p1, p2
}

func getMine(input int) int {
	total := 0
	new := calc(input)
	for new > 0 {
		total += new
		new = calc(new)
	}
	return total
}

func calc(input int) int {
	return int(math.Floor(float64(input)/3) - 2)
}

func getInput(filename string) []int {
	// Get file contents
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading file:%s", filename))
	}

	// Convert to int slice
	s := string(b)
	split := strings.Split(strings.TrimSuffix(s, "\n"), "\n")
	var result []int
	for _, v := range split {
		asInt, _ := strconv.Atoi(v)
		result = append(result, asInt)
	}
	return result
}
