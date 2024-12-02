package main

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filepath string) [][]int {
	result := [][]int{}

	data, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(data), "\n") {
		row := []int{}
		if line == "" {
			continue
		}
		for _, num := range strings.Split(line, " ") {
			number, _ := strconv.ParseInt(num, 10, 0)
			row = append(row, int(number))
		}

		result = append(result, row)
	}

	return result
}

func getDiffList(report []int) []int {
	diffList := []int{}
	for i := 0; i < len(report)-1; i++ {
		diffList = append(diffList, report[i+1]-report[i])
	}
	return diffList
}

func isSafe(diffList []int) bool {
	posCount := 0
	negCount := 0

	for _, diff := range diffList {
		diffAbs := int(math.Abs(float64(diff)))
		if diffAbs >= 1 && diffAbs <= 3 {
			if diff > 0 {
				posCount += 1
			} else {
				negCount += 1
			}
		} else {
			return false
		}
	}

	return !(posCount >= 1 && negCount >= 1)
}

func SolutionPartOne(input [][]int) int {
	result := 0

	for _, report := range input {
		diffList := getDiffList(report)
		if isSafe(diffList) {
			result += 1
		}

	}

	return result
}

func SolutionPartTwo(input [][]int) int {
	result := 0

	for _, report := range input {
		for i := 0; i < len(report); i++ {
			reportCopy := make([]int, len(report))
			copy(reportCopy, report)
			subReport := append(reportCopy[:i], reportCopy[i+1:]...)
			diffList := getDiffList(subReport)
			if isSafe(diffList) {
				result += 1
				break
			}
		}
	}

	return result
}
