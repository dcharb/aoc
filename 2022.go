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

const (
	rock = iota
	paper
	scissor
)

const (
	win = iota
	lose
	tie
)

func do20222(lines []string) {
	// a
	total := 0
	for _, line := range lines {
		f := strings.Fields(line)
		if len(f) != 2 {
			log.Println("incorrect number of elements")
			return
		}
		them := getHandA(f[0])
		us := getHandA(f[1])
		total += rpsPoints(them, us)
	}
	fmt.Println(total)

	// b
	total = 0
	for _, line := range lines {
		f := strings.Fields(line)
		if len(f) != 2 {
			log.Println("incorrect number of elements")
			return
		}
		them := getHandA(f[0])
		result := rpsToResult(f[1])
		us := getHandB(them, result)
		total += rpsPoints(them, us)
	}
	fmt.Println(total)
}

func rpsPoints(them, us int) int {
	total := 0
	switch us {
	case rock:
		total += 1
	case paper:
		total += 2
	case scissor:
		total += 3
	}
	result := rpsResult(them, us)
	switch result {
	case win:
		total += 6
	case tie:
		total += 3
	}
	return total
}

func getHandA(a string) int {
	switch a {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissor
	}
	log.Println("invalid")
	return -1
}

func rpsToResult(a string) int {
	switch a {
	case "X":
		return lose
	case "Y":
		return tie
	case "Z":
		return win
	}
	log.Println("invalid")
	return tie
}

func getHandB(them, result int) int {
	switch result {
	case tie:
		return them
	case win:
		switch them {
		case rock:
			return paper
		case paper:
			return scissor
		case scissor:
			return rock
		}
	case lose:
		switch them {
		case rock:
			return scissor
		case paper:
			return rock
		case scissor:
			return paper
		}
	}
	log.Println("invalid")
	return rock
}

// them vs you
func rpsResult(a, b int) int {
	if a == b {
		return tie
	}
	switch a {
	case rock:
		switch b {
		case paper:
			return win
		case scissor:
			return lose
		}
	case paper:
		switch b {
		case rock:
			return lose
		case scissor:
			return win
		}
	case scissor:
		switch b {
		case rock:
			return win
		case paper:
			return lose
		}
	}
	log.Println("invalid")
	return tie
}
