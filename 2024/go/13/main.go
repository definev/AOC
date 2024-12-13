package main

import (
	"bytes"
	"math"
	"os"
	"strconv"
)

type Button struct {
	X int
	Y int
}

type Pair struct {
	A    Button
	B    Button
	X, Y int
}

func ParseInput(filepath string) []Pair {
	result := []Pair{}

	data, _ := os.ReadFile(filepath)

	for _, chunk := range bytes.Split(data, []byte{'\n', '\n'}) {
		lines := bytes.Split(chunk, []byte{'\n'})

		parseButton := func(line []byte) Button {
			splits := bytes.Split(
				bytes.Split(line, []byte(": "))[1],
				[]byte(", "),
			)

			x, _ := strconv.Atoi(string(bytes.Split(splits[0], []byte{'+'})[1]))
			y, _ := strconv.Atoi(string(bytes.Split(splits[1], []byte{'+'})[1]))
			return Button{x, y}
		}

		buttonA := parseButton(lines[0])
		buttonB := parseButton(lines[1])

		var goalX, goalY int
		splits := bytes.Split(
			bytes.Split(lines[2], []byte(": "))[1],
			[]byte(", "),
		)
		goalX, _ = strconv.Atoi(
			string(bytes.Split(splits[0], []byte{'='})[1]),
		)
		goalY, _ = strconv.Atoi(
			string(bytes.Split(splits[1], []byte{'='})[1]),
		)

		result = append(result, Pair{buttonA, buttonB, goalX, goalY})
	}

	return result
}

func solve(input Pair) int {
	minACount := -1
	minBCount := 1

	isValid := func(aCount, bCount int) bool {
		if minACount == -1 {
			return true
		} else if aCount*3+bCount < minACount*3+minBCount {
			return true
		}
		return false
	}

	maxA := input.X / input.A.X

	aCount := 1
	bCount := -1

	for ; aCount <= maxA; aCount += 1 {
		axData, ayData := input.A.X*aCount, input.A.Y*aCount
		bxData, byData := input.X-axData, input.Y-ayData

		if bxData%input.B.X == 0 && byData%input.B.Y == 0 {
			if bxData/input.B.X == byData/input.B.Y {
				bCount = bxData / input.B.X
				if isValid(aCount, bCount) {
					minACount = aCount
					minBCount = bCount
					continue
				}
			}
		}
	}

	if minACount == -1 {
		return 0
	}

	return minACount*3 + minBCount
}

// 94 * a + 22 * b = 8400
// 34 * a + 67 * b = 5400
// x1 * a + y1 * b = m
// y1 * a + y2 * b = n
//
// x1*y2 * a + y1*y2 * b = m*y2
// x2*y1 * a + y2*y1 * b = n*y1

// a = (m*y2 - n*y1) / (x1*y2 - x2*y1)
// b = (m*x2 - n*x1) / (y1*x2 - y2*x1)
func solveMath(input Pair) int {
	a := float64(input.X*input.B.Y-input.Y*input.B.X) / float64(input.A.X*input.B.Y-input.B.X*input.A.Y)
	b := (float64(input.X) - float64(input.A.X)*a) / float64(input.B.X)

	if math.Round(a) == a && math.Round(b) == b {
		return int(a)*3 + int(b)
	}

	return 0

}

func SolutionPartOne(input []Pair) int {
	result := 0

	for _, pair := range input {
		result += solve(pair)
	}

	return result
}

func SolutionPartTwo(input []Pair) int {
	result := 0

	for _, pair := range input {
		pair.X += 10000000000000
		pair.Y += 10000000000000
		result += solveMath(pair)
	}

	return result
}
