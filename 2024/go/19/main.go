package main

import (
	"os"
	"strings"
)

type Store struct {
	Patterns []string
	Orders   []string
}

func ParseInput(filepath string) Store {
	raw, _ := os.ReadFile(filepath)
	parts := strings.Split(string(raw), "\n\n")

	store := Store{
		Patterns: strings.Split(parts[0], ", "),
		Orders:   strings.Split(parts[1], "\n"),
	}
	return store
}

type Data struct {
	Index int
}

func match(patterns []string, order string) int {
	q := []Data{{
		Index: 0,
	}}

	traversed := map[int]bool{}

	count := 0

	for len(q) > 0 {
		lastIndex := q[0].Index
		q = q[1:]

		if lastIndex == len(order) {
			count++
			continue
		}
		if lastIndex > len(order) {
			continue
		}
		if traversed[lastIndex] {
			continue
		}

		traversed[lastIndex] = true

		for _, pattern := range patterns {
			if pattern[0] != order[lastIndex] {
				continue
			}
			if lastIndex+len(pattern) > len(order) {
				continue
			}
			subString := order[lastIndex : lastIndex+len(pattern)]
			if subString == pattern {
				q = append(q, Data{
					Index: lastIndex + len(pattern),
				})
			}
		}
	}

	return count
}

func matchDFS(
	patterns []string, order string,
	count map[int]int, index int,
) int {
	if index == len(order) {
		return 1
	}
	if index > len(order) {
		return 0
	}
	if val, ok := count[index]; ok {
		return val
	}

	total := 0

	for _, pattern := range patterns {
		if pattern[0] != order[index] {
			continue
		}
		if index+len(pattern) > len(order) {
			continue
		}
		subString := order[index : index+len(pattern)]
		if subString == pattern {
			total += matchDFS(patterns, order, count, index+len(pattern))
		}
	}

	count[index] = total
	return total
}

func SolutionPartOne(store Store) int {
	count := 0

	for _, order := range store.Orders {
		if match(store.Patterns, order) > 0 {
			count++
		}
	}

	return count
}

func SolutionPartTwo(store Store) int {
	count := 0
	orderCandidates := []string{}

	for _, order := range store.Orders {
		orderCount := match(store.Patterns, order)
		if orderCount > 0 {
			orderCandidates = append(orderCandidates, order)
		}
	}

	for _, order := range orderCandidates {
		countMap := map[int]int{}
		count += matchDFS(store.Patterns, order, countMap, 0)
	}

	return count
}
