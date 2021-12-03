package main

import (
	"testing"
)

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
