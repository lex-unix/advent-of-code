package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	newLine = []byte("\n")
	space   = []byte(" ")
)

func day1_part1() int {
	lists, _ := os.ReadFile("input")
	listA := make([]int, 0, 100)
	listB := make([]int, 0, 100)
	for _, line := range bytes.Split(lists, newLine) {
		parts := bytes.Split(line, space)
		numA, _ := strconv.Atoi(string(parts[0]))
		numB, _ := strconv.Atoi(string(parts[len(parts)-1]))
		listA = append(listA, numA)
		listB = append(listB, numB)
	}

	slices.Sort(listA)
	slices.Sort(listB)

	distance := 0
	for i := range listA {
		diff := int(listA[i]) - int(listB[i])
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}

	return distance
}

func day1_part2() int {
	lists, _ := os.ReadFile("input")

	rightFreq := make(map[int]int)
	leftNums := make([]int, 0)

	for _, line := range bytes.Split(lists, newLine) {
		parts := bytes.Split(line, space)
		numA, _ := strconv.Atoi(string(parts[0]))
		numB, _ := strconv.Atoi(string(parts[len(parts)-1]))

		leftNums = append(leftNums, numA)
		rightFreq[numB]++
	}

	simScore := 0
	for _, leftNum := range leftNums {
		simScore += leftNum * rightFreq[leftNum]
	}

	return simScore
}

func day2_part1() int {
	reports, _ := os.ReadFile("input")
	var safeReports int

	isValidReport := func(report []byte) bool {
		levels := strings.Split(string(report), " ")
		if len(levels) < 2 {
			return false
		}
		isIncreasing := parseInt(levels[1]) > parseInt(levels[0])
		for i := 1; i < len(levels); i++ {
			curr := parseInt(levels[i])
			prev := parseInt(levels[i-1])
			diff := curr - prev
			if abs(diff) < 1 || abs(diff) > 3 {
				return false
			}
			if isIncreasing && diff <= 0 || !isIncreasing && diff >= 0 {
				return false
			}
		}
		return true
	}

	for _, report := range bytes.Split(reports, newLine) {
		if isValidReport(report) {
			safeReports++
		}
	}
	return safeReports
}

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

func main() {
	fmt.Println(day2_part1())
}
