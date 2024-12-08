package main

import (
	"bytes"
	"os"
	"strings"
)

func day7_part1() int {
	input, _ := os.ReadFile("input")
	lines := bytes.Split(bytes.TrimSpace(input), newLine)
	total := 0
	for _, line := range lines {
		parts := strings.Split(string(line), ": ")
		sum := parseInt(string(parts[0]))
		operandsCh := strings.Fields(parts[1])
		operands := make([]int, len(operandsCh))
		for i, num := range operandsCh {
			operands[i] = parseInt(string(num))
		}
		if validEquation(sum, operands[0], operands[1:]) {
			total += sum
		}

	}
	return total
}

func validEquation(sum, curr int, operands []int) bool {
	if len(operands) == 0 {
		return curr == sum
	}
	return validEquation(sum, curr+operands[0], operands[1:]) || validEquation(sum, curr*operands[0], operands[1:])
}
