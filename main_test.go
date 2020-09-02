package aoc

import "testing"

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

func Test1(t *testing.T) {
	one, two := day1("testdata/1.txt")
	e := 3412531
	if one != e {
		t.Errorf("e:%d, one:%d", e, one)
	}
	e = 5115927
	if two != e {
		t.Errorf("e:%d, two:%d", e, two)
	}
}
