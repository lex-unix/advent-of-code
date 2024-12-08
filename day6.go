package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func intialGuardPosition(m [][]byte) (int, int) {
	guardMarks := []string{"^", ">", "<", "v"}
	for i, row := range m {
		for j, cell := range row {
			if slices.Contains(guardMarks, string(cell)) {
				return i, j
			}
		}
	}
	return 0, 0
}

func day6_part1() int {
	content, _ := os.ReadFile("input")
	m := bytes.Split(bytes.TrimSpace(content), newLine)

	inBounds := func(x, y int) bool {
		h := len(m)
		w := len(m[0])
		return x < w && x >= 0 && y < h && y >= 0
	}

	y, x := intialGuardPosition(m)
	dx := 0
	dy := -1

	set := NewSet[string]()
	for inBounds(x, y) {
		set.Add(fmt.Sprintf("%d:%d", x, y))
		nx := x + dx
		ny := y + dy
		if inBounds(nx, ny) && m[ny][nx] == '#' {
			dx, dy = -dy, dx
		} else {
			x = nx
			y = ny
		}
	}

	return set.Size()
}
