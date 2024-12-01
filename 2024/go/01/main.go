package main

import (
	"bytes"
	"math"
	"os"
	"strconv"

	"golang.org/x/exp/constraints"
)

func ParseInput(filepath string) ([]int64, []int64) {
	data, _ := os.ReadFile(filepath)

	// split data
	var lefts []int64
	var rights []int64

	rows := bytes.Split(data, []byte("\n"))
	for _, row := range rows {
		nums := bytes.Split(row, []byte("   "))
		left, _ := strconv.ParseInt(string(nums[0]), 10, 0)
		right, _ := strconv.ParseInt(string(nums[1]), 10, 0)

		lefts = append(lefts, left)
		rights = append(rights, right)
	}

	return lefts, rights
}

func GenerateIndexList[T any](list []T) []int {
	indexList := make([]int, len(list))
	for i := range list {
		indexList[i] = i
	}
	return indexList
}

func Partition[T constraints.Ordered](list []T) int {
	pivot := len(list) - 1
	i := 0
	for j := 0; j < len(list); j++ {
		if list[j] < list[pivot] {
			list[i], list[j] = list[j], list[i]
			i++
		}
	}

	list[i], list[pivot] = list[pivot], list[i]
	return i
}

func QuickSort[T constraints.Ordered](list []T) {
	if len(list) <= 1 {
		return
	}

	pivot := Partition(list)
	QuickSort(list[:pivot])
	QuickSort(list[pivot+1:])
}

func SolutionPartOne(lefts, rights []int64) int {
	QuickSort(lefts)
	QuickSort(rights)

	result := 0
	for i := 0; i < len(lefts); i++ {
		result += int(math.Abs(float64(lefts[i] - rights[i])))
	}

	return result
}

func SolutionPartTwo(lefts, rights []int64) int {
	leftMap := map[int]int{}
	rightMap := map[int]int{}

	for i := 0; i < len(lefts); i++ {
		leftMap[int(lefts[i])] = leftMap[int(lefts[i])] + 1
		rightMap[int(rights[i])] = rightMap[int(rights[i])] + 1
	}

	var result int64 = 0
	for i := 0; i < len(lefts); i++ {
		num := lefts[i]
		result += num * int64(rightMap[int(num)])
	}

	return int(result)
}
