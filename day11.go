package main

import (
	"fmt"
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

func day11_part2() int {
	input, _ := os.ReadFile("input")
	stones := strings.Fields(string(input))
	memo := make(map[string]int)
	var breakStone func(stone string, blinks int) int
	breakStone = func(stone string, blinks int) int {
		key := fmt.Sprintf("%s:%d", stone, blinks)
		if count, ok := memo[key]; ok {
			return count
		}
		if blinks == 0 {
			memo[key] = 1
			return memo[key]
		}
		var count int
		if stone == "0" {
			count = breakStone("1", blinks-1)
		} else if len(stone)%2 == 0 {
			half := len(stone) / 2
			left := strconv.Itoa(parseInt(stone[:half]))
			right := strconv.Itoa(parseInt(stone[half:]))
			count = breakStone(left, blinks-1) + breakStone(right, blinks-1)
		} else {
			count = breakStone(strconv.Itoa(parseInt(stone)*2024), blinks-1)
		}
		memo[key] = count
		return memo[key]
	}

	stonesLen := 0
	for _, stone := range stones {
		stonesLen += breakStone(stone, 75)
	}
	return stonesLen
}
