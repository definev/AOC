package main

import (
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Tuple[F any, S any] struct {
	first  F
	second S
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func parseInput(fileName string) []Tuple[Tuple[int, int], Tuple[int, int]] {
	file, _ := os.ReadFile(fileName)
	rawList := strings.Split(string(file), "\n")
	result := make([]Tuple[Tuple[int, int], Tuple[int, int]], 0)

	for _, line := range rawList {
		list := strings.Split(line, ",")

		firstNumber := strings.Split(list[0], "-")
		firstNumberFirst, _ := strconv.Atoi(firstNumber[0])
		firstNumberSecond, _ := strconv.Atoi(firstNumber[1])

		secondNumber := strings.Split(list[1], "-")
		secondNumberFirst, _ := strconv.Atoi(secondNumber[0])
		secondNumberSecond, _ := strconv.Atoi(secondNumber[1])

		result = append(result, Tuple[Tuple[int, int], Tuple[int, int]]{
			first: Tuple[int, int]{
				first:  firstNumberFirst,
				second: firstNumberSecond,
			},
			second: Tuple[int, int]{
				first:  secondNumberFirst,
				second: secondNumberSecond,
			},
		})
	}

	return result
}

func main() {
	input := parseInput("../../sample/day04.txt")

	firstHalfProblem(input)
	lastHalfProblem(input)

}

func firstHalfProblem(input []Tuple[Tuple[int, int], Tuple[int, int]]) {
	total := 0

	for _, ele := range input {
		isSecondInFirst := ele.first.first <= ele.second.first && ele.first.second >= ele.second.second
		isFirstInSecond := ele.second.first <= ele.first.first && ele.second.second >= ele.first.second
		if isFirstInSecond || isSecondInFirst {
			total += 1
		}
	}

	print(total)
}

func lastHalfProblem(input []Tuple[Tuple[int, int], Tuple[int, int]]) {
	total := 0

	for _, ele := range input {
		inputRange := min(ele.first.second, ele.second.second) - max(ele.first.first, ele.second.first)
		if inputRange >= 0 {
			total += 1
		}
	}

	print(total)
}
