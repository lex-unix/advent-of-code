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

func isValidOrder(seq []string, beforeAfterRules map[string][]string) bool {
	for i, before := range seq {
		rule := beforeAfterRules[before]
		for _, after := range seq[i+1:] {
			found := slices.Index(rule, after)
			if found < 0 {
				return false
			}
		}
	}
	return true
}
