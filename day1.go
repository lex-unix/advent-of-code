package main

import (
	"bytes"
	"os"
	"slices"
	"strconv"
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
