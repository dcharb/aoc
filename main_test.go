package main

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
