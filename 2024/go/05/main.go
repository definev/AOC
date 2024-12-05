package main

import (
	"bytes"
	"os"
	"strconv"
)

type Input struct {
	Rules   [][]int
	Updates [][]int
}

func ParseInput(filepath string) Input {
	data, _ := os.ReadFile(filepath)

	slices := bytes.Split(data, []byte{'\n', '\n'})

	rules := [][]int{}
	for _, line := range bytes.Split(slices[0], []byte{'\n'}) {
		rule := []int{}
		for _, n := range bytes.Split(line, []byte{'|'}) {
			num, _ := strconv.ParseInt(string(n), 10, 64)
			rule = append(rule, int(num))
		}
		rules = append(rules, rule)
	}

	updates := [][]int{}
	for _, line := range bytes.Split(slices[1], []byte{'\n'}) {
		update := []int{}
		for _, n := range bytes.Split(line, []byte{','}) {
			num, _ := strconv.ParseInt(string(n), 10, 64)
			update = append(update, int(num))
		}
		updates = append(updates, update)
	}

	return Input{rules, updates}
}

func calculateOrderedRuleMap(input Input) map[int][]int {
	orderedRuleMap := map[int][]int{}

	for _, rule := range input.Rules {
		value := orderedRuleMap[rule[0]]
		orderedRuleMap[rule[0]] = append(value, rule[1])
	}
	return orderedRuleMap
}

func findUpdate(input Input, valid bool) [][]int {
	result := [][]int{}

	orderedRuleMap := calculateOrderedRuleMap(input)

	for _, update := range input.Updates {
		isValid := true

		for pieceIndex, piece := range update {
			for _, rulePiece := range orderedRuleMap[piece] {
				for i := pieceIndex - 1; i >= 0; i-- {
					if rulePiece == update[i] {
						isValid = false
						goto validate
					}
				}
			}
		}

	validate:
		if isValid == valid {
			result = append(result, update)
		}
	}

	return result
}

func calculateMiddleSum(input [][]int) int {
	result := 0
	for _, update := range input {
		result += update[len(update)/2]
	}
	return result
}

func SolutionPartOne(input Input) int {
	updates := findUpdate(input, true)
	return calculateMiddleSum(updates)
}

func SolutionPartTwo(input Input) int {
	orderedRuleMap := calculateOrderedRuleMap(input)
	updates := findUpdate(input, false)
	for updateIndex, update := range updates {
		orderedUpdate := make([]int, len(update))
		copy(orderedUpdate, update)

	order_update:
		for pieceIndex, piece := range orderedUpdate {
			for _, rulePiece := range orderedRuleMap[piece] {
				for i := pieceIndex - 1; i >= 0; i-- {
					if rulePiece == orderedUpdate[i] {
						orderedUpdate[i], orderedUpdate[pieceIndex] = orderedUpdate[pieceIndex], orderedUpdate[i]
						goto order_update
					}
				}
			}
		}

		updates[updateIndex] = orderedUpdate
	}
	return calculateMiddleSum(updates)
}
