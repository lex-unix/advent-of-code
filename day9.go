package main

import (
	"os"
	"strings"
)

type FileLayout struct {
	fileLen  int
	spaceLen int
}

func day9_part1() int {
	input, _ := os.ReadFile("input")
	diskMap := string(input)
	diskMap = strings.TrimSpace(diskMap)
	fileLayouts := make([]FileLayout, 0, len(diskMap)/2)
	for i := 0; i < len(diskMap)-2; i += 2 {
		f := FileLayout{
			fileLen:  parseInt(string(diskMap[i])),
			spaceLen: parseInt(string(diskMap[i+1])),
		}
		fileLayouts = append(fileLayouts, f)
	}
	fileLayouts = append(fileLayouts, FileLayout{
		fileLen: parseInt(string(diskMap[len(diskMap)-1])),
	})

	blocks := make([]int, 0, len(diskMap))
	for i, f := range fileLayouts {
		for range f.fileLen {
			blocks = append(blocks, i)
		}
		for range f.spaceLen {
			blocks = append(blocks, -1)
		}
	}

	i, j := 0, len(blocks)-1
	for i < j {
		for i < len(blocks) && blocks[i] != -1 {
			i++
		}
		if i < j {
			blocks[i] = blocks[j]
			blocks[j] = 0
			j--
		}
	}

	checksum := 0
	for i := range blocks {
		if blocks[i] != -1 {
			checksum += i * blocks[i]
		}
	}
	return checksum
}

func day9_part2() int {
	input, _ := os.ReadFile("input")
	diskMap := string(input)
	diskMap = strings.TrimSpace(diskMap)
	fileLayouts := make([]FileLayout, 0, len(diskMap)/2)

	for i := 0; i < len(diskMap)-2; i += 2 {
		f := FileLayout{
			fileLen:  parseInt(string(diskMap[i])),
			spaceLen: parseInt(string(diskMap[i+1])),
		}
		fileLayouts = append(fileLayouts, f)
	}
	fileLayouts = append(fileLayouts, FileLayout{
		fileLen: parseInt(string(diskMap[len(diskMap)-1])),
	})

	blocks := make([]int, 0, len(diskMap))
	for i, f := range fileLayouts {
		for range f.fileLen {
			blocks = append(blocks, i)
		}
		for range f.spaceLen {
			blocks = append(blocks, -1)
		}
	}

	for fileID := len(fileLayouts) - 1; fileID >= 0; fileID-- {
		fileStart := -1
		fileLen := 0
		for i := 0; i < len(blocks); i++ {
			if blocks[i] == fileID {
				if fileStart == -1 {
					fileStart = i
				}
				fileLen++
			}
		}

		if fileStart == -1 || fileLen == 0 {
			continue
		}
		for i := 0; i < fileStart; i++ {
			if blocks[i] == -1 {
				freeLen := 0
				freeStart := i
				for j := i; j < len(blocks) && blocks[j] == -1; j++ {
					freeLen++
				}
				if freeLen >= fileLen {
					for k := range fileLen {
						blocks[freeStart+k] = fileID
						blocks[fileStart+k] = -1
					}
					break
				}
				i += freeLen - 1
			}
		}
	}

	checksum := 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] != -1 {
			checksum += i * blocks[i]
		}
	}

	return checksum
}
