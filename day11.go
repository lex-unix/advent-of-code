package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func day11_part1() int {
	input, _ := os.ReadFile("input")
	nums := strings.Fields(strings.TrimSpace(string(input)))
	stones := make([]int, len(nums))
	for i, num := range nums {
		stones[i] = parseInt(num)
	}
	for range 25 {
		for i := 0; i < len(stones); i++ {
			stone := stones[i]
			if stone == 0 {
				stones[i] = 1
			} else if len(strconv.Itoa(stone))%2 == 0 {
				s := strconv.Itoa(stone)
				half := len(s) / 2
				left := s[:half]
				right := s[half:]
				stones[i] = parseInt(left)
				stones = slices.Insert(stones, i+1, parseInt(right))
				i++
			} else {
				stones[i] = stone * 2024
			}
		}

	}
	return len(stones)
}
