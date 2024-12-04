package main

import (
	"bytes"
	"fmt"
	"os"
)

func ParseInput(filepath string) [][]rune {
	data, _ := os.ReadFile(filepath)
	result := [][]rune{}

	for _, line := range bytes.Split(data, []byte("\n")) {
		result = append(result, []rune(string(line)))
	}

	return result
}

func reverseString(s string) string {
	// Convert string to rune slice to handle multi-byte characters
	runes := []rune(s)
	// Reverse the rune slice
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// Convert rune slice back to string
	return string(runes)
}

func scanPattern(input [][]rune, pattern string, positions [][]int) int {
	frame := len(pattern)

	result := 0
	for start := 0; start < len(positions); start++ {
		assemble := ""
		end := 0
		for inc := 0; inc < frame; inc++ {
			end = start + inc
			if end >= len(positions) {
				end = start + len(assemble)
				break
			}

			position := positions[end]
			if position[0] >= len(input) || position[1] >= len(input[0]) || position[0] < 0 || position[1] < 0 {
				break
			}
			assemble += string(input[position[0]][position[1]])
		}

		if assemble == pattern || reverseString(assemble) == pattern {
			// fmt.Printf("Found pattern at %v %v: %s\n", positions[start], positions[end], assemble)
			result++
		}
	}

	return result
}

func positionInDiagonal(inputLength int) [][][]int {
	result := [][][]int{}
	for i := 0; i < inputLength; i++ {
		positions := [][]int{}
		startPosition := []int{0, i}

		positions = append(positions, startPosition)
		for j := 1; j <= i; j++ {
			currentPosition := []int{j, i - j}
			positions = append(positions, currentPosition)
		}

		result = append(result, positions)
	}

	for i := inputLength - 2; i >= 0; i-- {
		positions := [][]int{}
		startPosition := []int{inputLength - 1, inputLength - i - 1}

		positions = append(positions, startPosition)
		for j := 1; j <= i; j++ {
			currentPosition := []int{inputLength - 1 - j, inputLength - i - 1 + j}
			positions = append(positions, currentPosition)
		}

		result = append(result, positions)
	}

	for i := inputLength - 1; i >= 0; i-- {
		positions := [][]int{}
		startPosition := []int{0, i}

		positions = append(positions, startPosition)
		for j := 1; j <= inputLength-1; j++ {
			currentPosition := []int{startPosition[0] + j, startPosition[1] + j}
			positions = append(positions, currentPosition)
		}

		result = append(result, positions)
	}

	for i := 1; i < inputLength; i++ {
		positions := [][]int{}
		startPosition := []int{inputLength - 1, inputLength - i - 1}

		positions = append(positions, startPosition)
		for j := 1; j <= inputLength-1; j++ {
			currentPosition := []int{startPosition[0] - j, startPosition[1] - j}
			positions = append(positions, currentPosition)
		}

		result = append(result, positions)
	}

	return result
}

func DebugInput(input [][]rune) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			fmt.Printf("%c", input[i][j])
		}
		fmt.Println()
	}
}

func DebugPositions(input [][]rune, positions [][]int) {
	for _, position := range positions {
		if (position[0] >= len(input)) || (position[1] >= len(input[0])) || (position[0] < 0) || (position[1] < 0) {
			continue
		}
		input[position[0]][position[1]] = '_'
	}

	DebugInput(input)
	fmt.Println()
}

func CopyMatrix(input [][]rune) [][]rune {
	result := [][]rune{}
	for i := 0; i < len(input); i++ {
		temp := make([]rune, len(input[i]))
		copy(temp, input[i])
		result = append(result, temp)
	}

	return result
}

func SolutionPartOne(input [][]rune) int {
	result := 0
	var positions [][]int

	// Scan horizontal
	for i := 0; i < len(input); i++ {
		positions = [][]int{}
		for j := 0; j < len(input[i]); j++ {
			positions = append(positions, []int{i, j})
		}
		result += scanPattern(input, "XMAS", positions)
	}

	// Scan vertical
	for j := 0; j < len(input[0]); j++ {
		positions = [][]int{}
		for i := 0; i < len(input); i++ {
			positions = append(positions, []int{i, j})
		}
		result += scanPattern(input, "XMAS", positions)
	}

	// Scan diagonal
	for _, positions := range positionInDiagonal(len(input)) {
		result += scanPattern(input, "XMAS", positions)
	}

	return result
}

func checkCrossXMAX(input [][]rune, pattern string, aPosition []int) bool {
	topLeft := []int{aPosition[0] - 1, aPosition[1] - 1}
	topRight := []int{aPosition[0] - 1, aPosition[1] + 1}
	bottomLeft := []int{aPosition[0] + 1, aPosition[1] - 1}
	bottomRight := []int{aPosition[0] + 1, aPosition[1] + 1}

	if topLeft[0] < 0 || topLeft[1] < 0 {
		return false
	}
	if topRight[0] < 0 || topRight[1] >= len(input[0]) {
		return false
	}
	if bottomLeft[0] >= len(input) || bottomLeft[1] < 0 {
		return false
	}
	if bottomRight[0] >= len(input) || bottomRight[1] >= len(input[0]) {
		return false
	}

	firstDiagonal := string(input[topLeft[0]][topLeft[1]]) + string(input[aPosition[0]][aPosition[1]]) + string(input[bottomRight[0]][bottomRight[1]])
	secondDiagonal := string(input[topRight[0]][topRight[1]]) + string(input[aPosition[0]][aPosition[1]]) + string(input[bottomLeft[0]][bottomLeft[1]])

	if (firstDiagonal == pattern || reverseString(firstDiagonal) == pattern) && (secondDiagonal == pattern || reverseString(secondDiagonal) == pattern) {
		return true
	}

	return false
}

func SolutionPartTwo(input [][]rune) int {
	result := 0

	aPositions := [][]int{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 'A' {
				aPositions = append(aPositions, []int{i, j})
			}
		}
	}

	for _, aPosition := range aPositions {
		if checkCrossXMAX(input, "MAS", aPosition) {
			result++
		}
	}

	return result
}
