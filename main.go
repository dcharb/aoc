package aoc

import (
	"math"
)

func day1(in []int) (int, int) {
	p1 := 0
	for _, v := range in {
		p1 += calc(v)
	}

	p2 := 0
	for _, v := range in {
		p2 += getMine(v)
	}

	return p1, p2
}

func day2(in []int) {
	// opcode 99 halt
	// opcode 1,a,b,c add values located at positions a,b and put into pos c
	// opcode 2,a,b,c multiplies values as above
	// after processing, step forward 4 positions
	for i := 0; i < len(in); i += 4 {
		switch in[i] {
		case 99:
			return
		case 1:
			in[in[i+3]] = in[in[i+1]] + in[in[i+2]]
		case 2:
			in[in[i+3]] = in[in[i+1]] * in[in[i+2]]
		}
	}
}

func d2p2(in []int) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			// make copy first so we can redo
			var t []int
			t = append(t, in...)
			t[1] = noun
			t[2] = verb
			day2(t)
			if t[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}

	panic("Didn't find answer")
}

func getMine(input int) int {
	total := 0
	new := calc(input)
	for new > 0 {
		total += new
		new = calc(new)
	}
	return total
}

func calc(input int) int {
	return int(math.Floor(float64(input)/3) - 2)
}
