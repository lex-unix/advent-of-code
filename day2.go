package main

import (
	"bytes"
	"os"
	"strings"
)

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
