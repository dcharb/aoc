package aoc

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func getInput(filename string) []int {
	// Get file contents
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading file:%s", filename))
	}

	// Convert to int slice
	s := string(b)
	s = strings.TrimSuffix(strings.ReplaceAll(s, "\n", ","), ",")
	split := strings.Split(s, ",")
	var result []int
	for _, v := range split {
		asInt, _ := strconv.Atoi(v)
		result = append(result, asInt)
	}
	return result
}

func TestCalc(t *testing.T) {
	if calc(1969) != 654 {
		t.Error()
	}
	if calc(100756) != 33583 {
		t.Error()
	}
}

func TestGetMine(t *testing.T) {
	if getMine(14) != 2 {
		t.Error()
	}
	if getMine(1969) != 966 {
		t.Error()
	}
	if getMine(100756) != 50346 {
		t.Error()
	}
}

func TestMy1(t *testing.T) {
	in := getInput("testdata/1.txt")
	one, two := day1(in)
	e := 3412531
	if one != e {
		t.Errorf("e:%d, one:%d", e, one)
	}
	e = 5115927
	if two != e {
		t.Errorf("e:%d, two:%d", e, two)
	}
}

func Test2(t *testing.T) {
	in := []int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	day2(in)
	expect := []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}
	if !reflect.DeepEqual(in, expect) {
		t.Errorf("e:%v, a:%v", expect, in)
	}

	in = []int{1, 0, 0, 0, 99}
	day2(in)
	expect = []int{2, 0, 0, 0, 99}
	if !reflect.DeepEqual(in, expect) {
		t.Errorf("e:%v, a:%v", expect, in)
	}

	in = []int{2, 3, 0, 3, 99}
	day2(in)
	expect = []int{2, 3, 0, 6, 99}
	if !reflect.DeepEqual(in, expect) {
		t.Errorf("e:%v, a:%v", expect, in)
	}

	in = []int{2, 4, 4, 5, 99, 0}
	day2(in)
	expect = []int{2, 4, 4, 5, 99, 9801}
	if !reflect.DeepEqual(in, expect) {
		t.Errorf("e:%v, a:%v", expect, in)
	}

	in = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	day2(in)
	expect = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	if !reflect.DeepEqual(in, expect) {
		t.Errorf("e:%v, a:%v", expect, in)
	}
}

func TestMy2(t *testing.T) {
	// Part 1
	in := getInput("testdata/2.txt")
	in[1] = 12
	in[2] = 2

	day2(in)
	a := in[0]
	e := 3085697
	if a != e {
		t.Errorf("e:%v, a:%v", e, a)
	}

	// Part 2
	in = getInput("testdata/2.txt")
	a = d2p2(in)
	e = 9425
	if a != e {
		t.Errorf("e:%v, a:%v", e, a)
	}
}
