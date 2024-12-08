package main

import (
	"bytes"
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
	m[y][x] = 'X'
	uniquePositions := 1

	dx := 0
	dy := -1

	for {
		for inBounds(x+dx, y+dy) && m[y+dy][x+dx] != '#' {
			x += dx
			y += dy
			if m[y][x] != 'X' {
				m[y][x] = 'X'
				uniquePositions++
			}
		}
		if !inBounds(x+dx, y+dy) {
			break
		}
		oldDx := dx
		dx = -dy
		dy = oldDx
	}

	return uniquePositions
}
