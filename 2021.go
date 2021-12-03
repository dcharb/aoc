package main

import (
	"fmt"
	"strconv"
	"strings"
)

func do20211(lines []string) {
	measures := toInts(lines)
	var previous int
	var largerCount int
	var larger2 int
	for i, measure := range measures {
		if i == 0 {
			previous = measure
			continue
		}
		if i >= 3 {
			prev := add(measures[i-3 : i])
			this := add(measures[i-2 : i+1])
			if this > prev {
				larger2++
			}
		}
		if measure > previous {
			largerCount++
		}
		previous = measure
	}
	fmt.Println(largerCount)
	fmt.Println(larger2)
}

func add(input []int) int {
	var sum int
	for _, a := range input {
		sum += a
	}
	return sum
}

func do20212(lines []string) {
	var hor int
	var depth int
	for _, line := range lines {
		f := strings.Fields(line)
		x, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "forward":
			hor += x
		case "down":
			depth += x
		case "up":
			depth -= x
		}
	}
	fmt.Println(hor * depth)
}
