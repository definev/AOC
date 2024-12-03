package main

import (
	"os"
	"regexp"
	"strconv"
)

func ParseInput(filepath string) string {
	data, _ := os.ReadFile(filepath)
	return string(data)
}

func parseMatch(match string) (int, int) {
	// mul(2,3)
	re := regexp.MustCompile(`\d+`)

	matches := re.FindAllString(match, -1)

	left, _ := strconv.Atoi(matches[0])
	right, _ := strconv.Atoi(matches[1])

	return left, right
}

func SolutionPartOne(input string) int {
	result := 0

	matches := mulMatch(input)

	for _, match := range matches {
		left, right := parseMatch(input[match[0]:match[1]])
		result += left * right
	}

	return result
}

func doMatch(input string) [][]int {
	re := regexp.MustCompile(`do\(()\)`)
	return re.FindAllStringIndex(input, -1)
}

func dontMatch(input string) [][]int {
	re := regexp.MustCompile(`don't\(()\)`)
	return re.FindAllStringIndex(input, -1)
}

func mulMatch(input string) [][]int {
	re := regexp.MustCompile(`(mul\((\d\d?\d?,\d\d?\d?)\))`)
	return re.FindAllStringIndex(input, -1)
}

func SolutionPartTwo(input string) int {
	result := 0

	matches := mulMatch(input)

	enable := true

	for _, match := range matches {

		doMatches := doMatch(input[:match[0]])
		dontMatches := dontMatch(input[:match[0]])
		if len(doMatches) == 0 && len(dontMatches) == 0 {
			enable = true
		} else if len(doMatches) == 0 {
			enable = false
		} else if len(dontMatches) == 0 {
			enable = true
		} else {
			doMatch := doMatches[len(doMatches)-1]
			dontMatch := dontMatches[len(dontMatches)-1]
			if doMatch[1] < dontMatch[1] {
				enable = false
			} else {
				enable = true
			}
		}

		if enable {
			left, right := parseMatch(input[match[0]:match[1]])
			result += left * right
		}
	}

	return result
}
