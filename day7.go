package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func day7_part2() int {
	input, _ := os.ReadFile("input")
	lines := bytes.Split(bytes.TrimSpace(input), newLine)
	total := 0
	for _, line := range lines {
		parts := strings.Split(string(line), ": ")
		target := parseInt(string(parts[0]))
		nums := strings.Fields(parts[1])
		operands := make([]int, len(nums))
		for i, num := range nums {
			operands[i] = parseInt(string(num))
		}
		if validEquation(target, operands[0], operands[1:]) {
			total += target
		}
	}
	return total
}

func validEquation(target, curr int, operands []int) bool {
	if len(operands) == 0 {
		return curr == target
	}
	mul := curr * operands[0]
	add := curr + operands[0]
	con := parseInt(fmt.Sprintf("%d%d", curr, operands[0]))
	return validEquation(target, add, operands[1:]) || validEquation(target, mul, operands[1:]) || validEquation(target, con, operands[1:])
}
