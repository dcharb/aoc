package main

import (
	"fmt"
	"testing"
)

func Test3b(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010",
		"01010"}
	actual := do20213b(input)
	fmt.Println(actual)
	if actual != 230 {
		t.Fail()
	}
}
