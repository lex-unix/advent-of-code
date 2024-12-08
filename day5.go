package main

import (
	"os"
	"slices"
	"strings"
)

func day5_part1() (sum int) {
	content, _ := os.ReadFile("input")
	parts := strings.Split(string(content), "\n\n")
	rules := strings.Split(parts[0], "\n")
	updates := strings.Split(parts[1], "\n")
	beforeAfterRules := make(map[string][]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before := parts[0]
		after := parts[1]
		beforeAfterRules[before] = append(beforeAfterRules[before], after)
	}
	for _, update := range updates {
		seq := strings.Split(update, ",")
		if isValidOrder(seq, beforeAfterRules) {
			m := len(seq) / 2
			sum += parseInt(seq[m])
		}
	}
	return sum
}

func day5_part2() (sum int) {
	content, _ := os.ReadFile("input")
	parts := strings.Split(string(content), "\n\n")
	rules := strings.Split(parts[0], "\n")
	updates := strings.Split(parts[1], "\n")
	beforeAfterRules := make(map[string][]string)
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		before := parts[0]
		after := parts[1]
		beforeAfterRules[before] = append(beforeAfterRules[before], after)
	}
	for _, update := range updates {
		seq := strings.Split(update, ",")
		if !isValidOrder(seq, beforeAfterRules) {
			sum += findMiddlePage(seq, beforeAfterRules)
		}
	}
	return sum
}

func isValidOrder(seq []string, beforeAfterRules map[string][]string) bool {
	for i, before := range seq {
		validAfterPages := beforeAfterRules[before]
		for _, after := range seq[i+1:] {
			if !slices.Contains(validAfterPages, after) {
				return false
			}
		}
	}
	return true
}

func findMiddlePage(seq []string, beforeAfterRules map[string][]string) int {
	m := len(seq) / 2
	for i, before := range seq {
		foundRules := 0
		validAfterPages := beforeAfterRules[before]
		for j, after := range seq {
			if slices.Contains(validAfterPages, after) && j != i {
				foundRules++
			}
		}
		if foundRules == m {
			return parseInt(before)
		}
	}
	return 0
}
