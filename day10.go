package main

import (
	"fmt"
	"os"
	"strings"
)

func day10_part1() int {
	input, _ := os.ReadFile("input")
	rows := strings.Split(strings.TrimSpace(string(input)), "\n")
	hikingMap := make([][]int, len(rows))
	for i, row := range rows {
		trailSegment := make([]int, len(row))
		for j, cell := range row {
			trailSegment[j] = (parseInt(string(cell)))
		}
		hikingMap[i] = trailSegment
	}
	totalScore := 0
	for i, row := range hikingMap {
		for j, cell := range row {
			if cell == 0 {
				totalScore += trailHeadScore(i, j, hikingMap)
			}
		}
	}
	return totalScore
}

func trailHeadScore(i, j int, hikingMap [][]int) int {
	n := len(hikingMap)
	m := len(hikingMap[0])
	visited := NewSet[string]()
	var passableTrails func(i, j, next int) int
	passableTrails = func(i, j, next int) int {
		if i >= n || i < 0 || j >= m || j < 0 {
			return 0
		}
		if hikingMap[i][j] != next {
			return 0
		}
		key := fmt.Sprintf("%d:%d", i, j)
		if next == 9 && !visited.Exists(key) {
			visited.Add(key)
			return 1
		}
		next++
		return passableTrails(i+1, j, next) +
			passableTrails(i, j+1, next) +
			passableTrails(i-1, j, next) +
			passableTrails(i, j-1, next)
	}
	return passableTrails(i, j, 0)
}
