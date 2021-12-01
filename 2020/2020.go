package aoc2020

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func do1a(nums []int) (int, error) {
	for _, a := range nums {
		for _, b := range nums {
			if a+b == 2020 {
				return a * b, nil
			}
		}
	}

	return 0, fmt.Errorf("Not found")
}

func do1b(nums []int) (int, error) {
	for _, a := range nums {
		for _, b := range nums {
			for _, c := range nums {
				if a+b+c == 2020 {
					return a * b * c, nil
				}
			}
		}
	}

	return 0, fmt.Errorf("Not found")
}

type policy struct {
	min      int
	max      int
	letter   string
	password string
}

func parsePolicy(line string) policy {
	// 1-3 a: abcde
	pattern := `(?P<min>\d+)-(?P<max>\d+) (?P<letter>\w): (?P<password>\w+)`
	meta := regexp.MustCompile(pattern)
	matches := meta.FindStringSubmatch(line)
	names := meta.SubexpNames()
	p := policy{}
	for i, match := range matches {
		if i != 0 {
			switch names[i] {
			case "min":
				m, _ := strconv.Atoi(match)
				p.min = m
			case "max":
				m, _ := strconv.Atoi(match)
				p.max = m
			case "letter":
				p.letter = match
			case "password":
				p.password = match
			}
		}
	}

	return p
}

func isPolicyValid(p policy) bool {
	count := strings.Count(p.password, p.letter)
	return count >= p.min && count <= p.max
}

func do2a(policies []policy) int {
	count := 0
	for _, p := range policies {
		if isPolicyValid(p) {
			count++
		}
	}
	return count
}

func isValidSecondPolicy(p policy) bool {
	return (string(p.password[p.min-1]) == p.letter) != (string(p.password[p.max-1]) == p.letter)
}

func do2b(policies []policy) int {
	// min and max now mean different positions
	count := 0
	for _, p := range policies {
		if isValidSecondPolicy(p) {
			count++
		}
	}
	return count
}

func parseHill(lines []string) [][]bool {
	var isTree [][]bool
	for _, line := range lines {
		var row []bool
		for _, r := range line {
			row = append(row, r == '#')
		}
		isTree = append(isTree, row)
	}
	return isTree
}

func slide(isTree [][]bool, slopeX, slopeY int) int {
	count := 0
	x := 0
	y := 0
	for {
		if y >= len(isTree) {
			break
		}
		if x >= len(isTree[y]) {
			x = x % len(isTree[y])
		}
		if isTree[y][x] {
			count++
		}
		x += slopeX
		y += slopeY
	}
	return count
}

func do3a(isTree [][]bool) int {
	return slide(isTree, 3, 1)
}

func do3b(isTree [][]bool) int {
	count := slide(isTree, 1, 1)
	count *= slide(isTree, 3, 1)
	count *= slide(isTree, 5, 1)
	count *= slide(isTree, 7, 1)
	count *= slide(isTree, 1, 2)
	return count
}

type passport map[string]string

func parsePassports(lines []string) []passport {
	var ports []passport
	p := passport{}
	for _, line := range lines {
		if line == "" {
			ports = append(ports, p)
			p = passport{}
		}
		for _, kv := range strings.Fields(line) {
			split := strings.Split(kv, ":")
			p[split[0]] = split[1]
		}
	}
	ports = append(ports, p)
	return ports
}

func isValidPassport(port passport) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, r := range required {
		if _, ok := port[r]; !ok {
			return false
		}
	}
	return true
}

func do4a(passports []passport) int {
	count := 0
	for _, port := range passports {
		if isValidPassport(port) {
			count++
		}
	}
	return count
}

func checkYear(s string, min, max int) bool {
	if s == "" {
		return false
	}

	asInt, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if len(s) != 4 || asInt < min || asInt > max {
		return false
	}
	return true
}

func checkHeight(s string) bool {
	pattern := `^(\d+)(cm|in)$`
	meta := regexp.MustCompile(pattern)
	matches := meta.FindStringSubmatch(s)
	if len(matches) != 3 {
		return false
	}
	num, _ := strconv.Atoi(matches[1])
	switch matches[2] {
	case "cm":
		return num >= 150 && num <= 193
	case "in":
		return num >= 59 && num <= 76
	}
	return false
}

func checkHairColor(s string) bool {
	pattern := `^#([0-9]|[a-f]){6}$`
	meta := regexp.MustCompile(pattern)
	return meta.MatchString(s)
}

func checkEyeColor(s string) bool {
	switch s {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	}
	return false
}

func checkPID(s string) bool {
	pattern := `^\d{9}$`
	meta := regexp.MustCompile(pattern)
	return meta.MatchString(s)
}

func isValidSecondPassport(port passport) bool {
	if !isValidPassport(port) {
		return false
	}

	return checkYear(port["byr"], 1920, 2002) &&
		checkYear(port["iyr"], 2010, 2020) &&
		checkYear(port["eyr"], 2020, 2030) &&
		checkHeight(port["hgt"]) &&
		checkHairColor(port["hcl"]) &&
		checkEyeColor(port["ecl"]) &&
		checkPID(port["pid"])
}

func do4b(passports []passport) int {
	count := 0
	for _, port := range passports {
		if isValidSecondPassport(port) {
			count++
		}
	}
	return count
}

func toSeatId(line string) int {
	row := strings.ReplaceAll(line[:7], "F", "0")
	row = strings.ReplaceAll(row, "B", "1")
	r, _ := strconv.ParseInt(row, 2, 64)
	col := strings.ReplaceAll(line[7:], "L", "0")
	col = strings.ReplaceAll(col, "R", "1")
	c, _ := strconv.ParseInt(col, 2, 64)
	return int(r*8 + c)
}

func do5a(lines []string) int {
	highest := 0
	for _, line := range lines {
		id := toSeatId(line)
		if id > highest {
			highest = id
		}
	}
	return highest
}

func do5b(lines []string) int {
	var list []int
	for _, line := range lines {
		list = append(list, toSeatId(line))
	}
	sort.Ints(list)
	var older int
	for i, id := range list {
		if i == 0 {
			older = id
			continue
		}

		if id-older == 2 {
			return older + 1
		}

		older = id
	}
	return 0
}

func parseLine(line string) int {
	return 0
}

func do6a(lines []string) int {
	letters := make(map[rune]bool)
	count := 0
	for _, line := range lines {
		if line == "" {
			count += len(letters)
			letters = make(map[rune]bool)
			continue
		}
		for _, letter := range line {
			letters[letter] = true
		}
	}
	count += len(letters)
	return count
}

func do6b(lines []string) int {
	var has string
	var tmp string
	reset := true
	count := 0
	for _, line := range lines {
		if line == "" {
			reset = true
			count += len(has)
			continue
		}
		if reset {
			reset = false
			has = line
			continue
		}
		for _, r := range line {
			if strings.Contains(has, string(r)) {
				tmp += string(r)
			}
		}
		has = tmp
		tmp = ""
	}
	count += len(has)
	return count
}

type bag struct {
	color      string
	canContain map[string]int
}

func getContains(line string) map[string]int {
	// " x colorB bag(s)[, y colorC bag(s)]."
	// " no other bags."
	clean := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(line, "bags", ""), "bag", ""), ".", "")
	sep := strings.Split(clean, ",")
	result := make(map[string]int)
	for _, one := range sep {
		if one == " no other " {
			continue
		}
		f := strings.Fields(one)
		i, _ := strconv.Atoi(f[0])
		result[fmt.Sprintf("%s %s", f[1], f[2])] = i
	}
	return result
}

func parseBag(line string) bag {
	// colorA bags contain x colorB bag(s)[, y colorC bag(s)].
	pattern := `(\w+\s\w+) bags contain(.*)`
	meta := regexp.MustCompile(pattern)
	matches := meta.FindStringSubmatch(line)
	b := bag{}
	for i, match := range matches {
		//fmt.Printf("i:%d, match:%s\n", i, match)
		switch i {
		case 0:
			continue
		case 1:
			b.color = match
		case 2:
			if match != "" {
				b.canContain = getContains(match)
			}
		}
	}
	return b
}

func do7a(lines []string) int {
	var bags []bag
	for _, line := range lines {
		bags = append(bags, parseBag(line))
	}
	for _, bag := range bags {
		hasrecursiveBag(bag)
	}
	return 0
}

func do7b(lines []string) int {
	return 0
}
