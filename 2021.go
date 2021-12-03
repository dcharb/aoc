package main

import (
	"fmt"
	"log"
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
	var aim int
	for _, line := range lines {
		f := strings.Fields(line)
		x, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "forward":
			hor += x
			depth += aim * x
		case "down":
			aim += x
		case "up":
			aim -= x
		}
	}
	fmt.Println(hor * depth)
}

func do20213(lines []string) {
	var gammaString string
	var epsString string
	var gamma0 []int
	var gamma1 []int
	for _, line := range lines {
		if gamma0 == nil {
			gamma0 = make([]int, len(line))
			gamma1 = make([]int, len(line))
		}
		for i, c := range line {
			if c == '0' {
				gamma0[i]++
			} else {
				gamma1[i]++
			}
		}
	}
	for i := range gamma0 {
		if gamma0[i] > gamma1[i] {
			gammaString += "0"
			epsString += "1"
		} else {
			gammaString += "1"
			epsString += "0"
		}
	}
	gamma := bin2dec(gammaString)
	eps := bin2dec(epsString)
	fmt.Println(gamma * eps)
}

func bin2dec(bin string) int {
	o, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(o)
}
