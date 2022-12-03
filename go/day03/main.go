package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var scores map[rune]int = map[rune]int{}

func initScores() {
	for i := 97; i < 97+26; i++ {
		scores[rune(i)] = i - 96
	}
	for i := 65; i < 65+26; i++ {
		scores[rune(i)] = i - 64 + 26
	}
}

func check(err error) {
	if err != nil {
		log.Fatalf("Failed : %v", err)
	}
}

func parseInput() []string {
	input, _ := os.ReadFile("../../input/day03.txt")
	return strings.Split(strings.TrimSpace(string(input)), "\n")
}

func parseInLineProblemInput(input []string) [][]string {
	result := make([][]string, 0)

	for i := 0; i < len(input); i++ {
		raw := input[i]
		line := make([]string, 0)
		line = append(line, raw[0:len(raw)/2])
		line = append(line, raw[len(raw)/2:])
		result = append(result, line)
	}

	return result
}

func parseLinesProblemInput(input []string, length int) [][]string {
	result := make([][]string, 0)

	for i := 0; i < len(input)/length; i++ {
		line := make([]string, length)
		for j := 0; j < length; j++ {
			line[j] = input[j+length*i]
		}
		result = append(result, line)
	}

	return result
}

func main() {
	initScores()
	input := parseInput()
	firstInput := parseInLineProblemInput(input)
	secondInput := parseLinesProblemInput(input, 3)
	firstHalfProblem(firstInput)
	lastHalfProblem(secondInput)
}

func total(str string) int {
	total := 0
	for _, char := range str {
		total += scores[char]
	}

	return total
}

func isSameRune(list []string, newList []string) bool {
	result := true
	for i := 0; i < len(list); i++ {
		result = len(list[i]) != len(newList[i]) && result
	}
	return result
}

func replaceAll(line []string, k rune) []string {
	newList := []string{}

	for _, str := range line {
		str = strings.ReplaceAll(str, string(k), "")
		newList = append(newList, str)
	}

	return newList
}

func solve(input [][]string) {
	keys := make([]rune, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sameStr := ""
	for _, line := range input {
		for _, k := range keys {
			newLine := replaceAll(line, k)
			if isSameRune(line, newLine) {
				sameStr = fmt.Sprintf("%s%s", sameStr, string(k))
			}
		}
	}
	fmt.Println(total(sameStr))
}

func firstHalfProblem(input [][]string) {
	solve(input)
}

func lastHalfProblem(input [][]string) {
	solve(input)
}
