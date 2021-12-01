package aoc2020

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
)

func read(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	return lines
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

func Test1a(t *testing.T) {
	lines := read("testdata/1.txt")
	result, err := do1a(toInts(lines))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
}

func Test1b(t *testing.T) {
	lines := read("testdata/1.txt")
	result, err := do1b(toInts(lines))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(result)
}

func Test2a(t *testing.T) {
	lines := read("testdata/2.txt")
	var policies []policy
	for _, line := range lines {
		policies = append(policies, parsePolicy(line))
	}
	fmt.Println(do2a(policies))
}

func Test2b(t *testing.T) {
	lines := read("testdata/2.txt")
	var policies []policy
	for _, line := range lines {
		policies = append(policies, parsePolicy(line))
	}
	fmt.Println(do2b(policies))
}

func Test3a(t *testing.T) {
	lines := read("testdata/3.txt")
	isTree := parseHill(lines)
	fmt.Println(do3a(isTree))
}

func Test3b(t *testing.T) {
	lines := read("testdata/3.txt")
	isTree := parseHill(lines)
	fmt.Println(do3b(isTree))
}

func Test4a(t *testing.T) {
	lines := read("testdata/4.txt")
	passports := parsePassports(lines)
	fmt.Println(do4a(passports))
}

func TestHeight(t *testing.T) {
	if !checkHeight("60in") {
		t.Error()
	}
	if !checkHeight("190cm") {
		t.Error()
	}
	if checkHeight("190in") {
		t.Error()
	}
	if checkHeight("190") {
		t.Error()
	}
}

func TestHair(t *testing.T) {
	if !checkHairColor("#123abc") {
		t.Error()
	}
	if checkHairColor("#123abca") {
		t.Error()
	}
	if checkHairColor("123abc") {
		t.Error()
	}
}

func Test4b(t *testing.T) {
	lines := read("testdata/4.txt")
	passports := parsePassports(lines)
	fmt.Println(do4b(passports))
}

func TestToSeatId(t *testing.T) {
	if toSeatId("BFFFBBFRRR") != 567 {
		t.Error()
	}
	if toSeatId("FFFBBBFRRR") != 119 {
		t.Error()
	}
	if toSeatId("BBFFBBFRLL") != 820 {
		t.Error()
	}
}

func Test5a(t *testing.T) {
	lines := read("testdata/5.txt")
	fmt.Println(do5a(lines))
}

func Test5b(t *testing.T) {
	lines := read("testdata/5.txt")
	fmt.Println(do5b(lines))
}

func Test6a(t *testing.T) {
	lines := read("testdata/6.txt")
	fmt.Println(do6a(lines))
}

func Test6b(t *testing.T) {
	lines := read("testdata/6.txt")
	fmt.Println(do6b(lines))
}

func TestParseBag(t *testing.T) {
	for _, v := range []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	} {
		parseBag(v)
	}
}

func Test7a(t *testing.T) {
	lines := read("testdata/7.txt")
	fmt.Println(do7a(lines))
}

func Test7b(t *testing.T) {
	lines := read("testdata/7.txt")
	fmt.Println(do7b(lines))
}
