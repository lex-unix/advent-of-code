package main

import (
	"os"
	"strconv"
	"strings"
)

func day3_part1() int {
	input, _ := os.ReadFile("input")
	text := string(input)

	sum, i := 0, 0
	for i < len(text) {
		if i+3 < len(text) && text[i:i+4] == "mul(" {
			if result, newPos, ok := parseMul(text[i:]); ok {
				sum += result
				i += newPos
				continue
			}
		}
		i++
	}

	return sum
}

func parseMul(text string) (int, int, bool) {
	closePos := strings.Index(text[4:], ")")
	if closePos == -1 {
		return 0, 0, false
	}
	closePos += 4
	content := text[4:closePos]

	nums := strings.Split(content, ",")
	if len(nums) != 2 {
		return 0, 0, false
	}

	x, err := strconv.Atoi(strings.TrimSpace(nums[0]))
	if err != nil || x < 0 || x > 999 {
		return 0, 0, false
	}

	y, err := strconv.Atoi(strings.TrimSpace(nums[1]))
	if err != nil || y < 0 || y > 999 {
		return 0, 0, false
	}

	return x * y, closePos + 1, true
}
