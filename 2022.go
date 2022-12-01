package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func do20221(lines []string) {
	// a
	cals := 0
	maxCals := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			if cals > maxCals {
				maxCals = cals
			}
			cals = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			cals += i
		}
	}
	if cals > maxCals {
		maxCals = cals
	}
	fmt.Println(maxCals)

	// b
	cals = 0
	var maxes []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			maxes = updateMaxes(cals, maxes)
			cals = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			cals += i
		}
	}
	maxes = updateMaxes(cals, maxes)
	fmt.Println(sum(maxes))
}

func updateMaxes(next int, maxes []int) []int {
	if len(maxes) < 3 {
		maxes = append(maxes, next)
		sort.Ints(maxes)
	} else if next > maxes[0] {
		maxes[0] = next
		sort.Ints(maxes)
	}
	return maxes
}

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
