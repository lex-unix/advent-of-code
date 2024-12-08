package main

import "strconv"

var (
	newLine = []byte("\n")
	space   = []byte(" ")
)

func parseInt(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func count[S ~[]E, E comparable](s S, target E) (total int) {
	for _, v := range s {
		if v == target {
			total++
		}
	}
	return
}
