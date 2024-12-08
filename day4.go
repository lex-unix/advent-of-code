package main

import (
	"os"
	"strings"
)

func day4_part1() int {
	content, _ := os.ReadFile("input")
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")
	searchWord := "XMAS"
	searchWordReverse := "SAMX"
	searchWordLen := len(searchWord)
	found := 0

	maxRow := len(lines)
	maxCol := len(lines[0])

	directions := [][2]int{{0, 1}, {1, 0}, {1, 1}, {-1, 1}}

	for i := 0; i < maxRow; i++ {
		for j := 0; j < maxCol; j++ {
			for _, dir := range directions {
				for _, word := range []string{searchWord, searchWordReverse} {
					match := true
					for k := 0; k < searchWordLen; k++ {
						ni := i + k*dir[0]
						nj := j + k*dir[1]
						if ni < 0 || nj < 0 || ni >= maxRow || nj >= maxCol || lines[ni][nj] != word[k] {
							match = false
							break
						}
					}
					if match {
						found++
					}
				}
			}
		}
	}
	return found
}
