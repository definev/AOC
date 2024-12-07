package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ParseInput(filepath string) map[int][]int {
	data, _ := os.ReadFile(filepath)
	lines := strings.Split(string(data), "\n")

	result := map[int][]int{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		total, _ := strconv.Atoi(parts[0])
		result[total] = make([]int, 0)

		ops := strings.Split(strings.TrimSpace(parts[1]), " ")
		for _, op := range ops {
			n, _ := strconv.Atoi(op)
			result[total] = append(result[total], n)
		}
	}

	return result
}

type OperatorFunc func(currentValue int, nextNumber int) int

func ComputeTarget(
	target int, currentValue int, remainingValues []int,
	operatorFuncList []OperatorFunc,
) bool {
	if len(remainingValues) == 0 {
		return target == currentValue
	}

	nextNumber := remainingValues[0]
	restRemainingValues := remainingValues[1:]

	for _, operatorFunc := range operatorFuncList {
		newCurrentValue := operatorFunc(currentValue, nextNumber)
		if ComputeTarget(target, newCurrentValue, restRemainingValues, operatorFuncList) {
			return true
		}
	}

	return false
}

func SolutionPartOne(input map[int][]int) int {
	result := 0

	operatorFuncList := []OperatorFunc{
		func(currentValue int, nextNumber int) int {
			return currentValue + nextNumber
		},
		func(currentValue int, nextNumber int) int {
			return currentValue * nextNumber
		},
	}

	for target, ops := range input {
		if ComputeTarget(target, ops[0], ops[1:], operatorFuncList) {
			result += target
		}
	}

	return result
}

func SolutionPartTwo(input map[int][]int) int {
	result := 0

	operatorFuncList := []OperatorFunc{
		func(currentValue int, nextNumber int) int {
			return currentValue + nextNumber
		},
		func(currentValue int, nextNumber int) int {
			return currentValue * nextNumber
		},
		func(currentValue, nextNumber int) int {
			value, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, nextNumber))
			return value
		},
	}

	for target, ops := range input {
		if ComputeTarget(target, ops[0], ops[1:], operatorFuncList) {
			result += target
		}
	}

	return result
}
