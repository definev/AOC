package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

type MemMap struct {
	raw       [][]rune
	corrupted []Position
	height    int
	width     int
}

func ParseInt(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}

func ParseInput(filepath string) MemMap {
	raw, _ := os.ReadFile(filepath)

	result := MemMap{
		raw:       [][]rune{},
		corrupted: []Position{},
	}

	maxHeight := 0
	maxWidth := 0

	for _, line := range strings.Split(string(raw), "\n") {
		parts := strings.Split(line, ",")
		x, y := ParseInt(parts[0]), ParseInt(parts[1])
		result.corrupted = append(result.corrupted, Position{x, y})
		maxHeight = int(math.Max(float64(maxHeight), float64(y)))
		maxWidth = int(math.Max(float64(maxWidth), float64(x)))
	}

	result.height = maxHeight + 1
	result.width = maxWidth + 1

	for y := 0; y < result.height; y++ {
		result.raw = append(result.raw, []rune{})
		for x := 0; x < result.width; x++ {
			result.raw[y] = append(result.raw[y], '.')
		}
	}

	return result
}

type Position [2]int

type Path [3]int

func LogRaw(input MemMap) {
	for y := 0; y < input.height; y++ {
		for x := 0; x < input.width; x++ {
			print(string(input.raw[y][x]))
		}
		println()
	}
}

func SolutionPartOne(input MemMap, maxBytes int) int {
	queue := []Path{{0, 0, 0}}
	traversed := map[Position]bool{}

	for i := 0; i < maxBytes; i++ {
		corruptedByte := input.corrupted[i]
		input.raw[corruptedByte[1]][corruptedByte[0]] = '#'
	}
	// LogRaw(input)

	for len(queue) > 0 {
		step := queue[0]
		queue = queue[1:]

		if step[0] < 0 || step[1] < 0 || step[0] >= input.width || step[1] >= input.height {
			continue
		}
		if traversed[Position{step[0], step[1]}] {
			continue
		}
		char := input.raw[step[1]][step[0]]
		if char == '#' {
			continue
		}

		traversed[Position{step[0], step[1]}] = true

		if step[0] == input.width-1 && step[1] == input.height-1 {
			return step[2]
		}

		queue = append(queue, [3]int{step[0] + 1, step[1], step[2] + 1})
		queue = append(queue, [3]int{step[0], step[1] + 1, step[2] + 1})
		queue = append(queue, [3]int{step[0] - 1, step[1], step[2] + 1})
		queue = append(queue, [3]int{step[0], step[1] - 1, step[2] + 1})
	}

	return -1
}

func SolutionPartTwo(input MemMap) Position {
	maxSteps := 1
	for maxSteps <= len(input.corrupted) && SolutionPartOne(input, maxSteps) != -1 {
		maxSteps++
	}

	return input.corrupted[maxSteps-1]
}
