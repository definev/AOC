package main

import (
	"os"
	"strings"
)

func ParseInput(filepath string) [][]int {
	data, _ := os.ReadFile(filepath)

	var grid [][]int
	for _, line := range strings.Split(string(data), "\n") {
		lineInt := []int{}
		for _, cell := range line {
			lineInt = append(lineInt, int(cell-'0'))
		}
		grid = append(grid, lineInt)
	}

	return grid
}

type Position struct {
	X int
	Y int
}

func Traverse(traversed map[Position]bool, input [][]int, position Position, target int) {
	if position.X < 0 || position.X >= len(input) || position.Y < 0 || position.Y >= len(input[0]) {
		return
	}

	if input[position.X][position.Y] == 9 && target == 9 {
		traversed[position] = true
		return
	}

	if input[position.X][position.Y] == target {
		Traverse(traversed, input, Position{position.X - 1, position.Y}, target+1)
		Traverse(traversed, input, Position{position.X + 1, position.Y}, target+1)
		Traverse(traversed, input, Position{position.X, position.Y - 1}, target+1)
		Traverse(traversed, input, Position{position.X, position.Y + 1}, target+1)
	}
}

func ComputeHeadPoint(input [][]int, headPosition Position) int {
	traversed := map[Position]bool{}
	Traverse(traversed, input, Position{headPosition.X - 1, headPosition.Y}, 1)
	Traverse(traversed, input, Position{headPosition.X + 1, headPosition.Y}, 1)
	Traverse(traversed, input, Position{headPosition.X, headPosition.Y - 1}, 1)
	Traverse(traversed, input, Position{headPosition.X, headPosition.Y + 1}, 1)

	return len(traversed)
}

func TraverseWithPath(
	input [][]int,
	position Position, target int,
	knownPaths *[][]Position,
	path *[]Position,
) {
	if position.X < 0 || position.X >= len(input) || position.Y < 0 || position.Y >= len(input[0]) {
		// fmt.Print("| out of bound\n")
		return
	}

	if input[position.X][position.Y] == target {
		// fmt.Printf("-> %v ", position)
		*path = append(*path, position)

		if target == 9 {
			// fmt.Print(" | FINISH")
			unique := len(*knownPaths) == 0
		outer:
			for _, knownPath := range *knownPaths {
				for i, p := range knownPath {
					if p != (*path)[i] {
						unique = true
						break outer
					}
				}
			}
			if unique {
				// fmt.Print(" | UNIQUE")
				copyPath := make([]Position, len(*path))
				copy(copyPath, *path)
				*knownPaths = append(*knownPaths, copyPath)
			}
			// fmt.Println()
		} else {
			TraverseWithPath(input, Position{position.X - 1, position.Y}, target+1, knownPaths, path)
			TraverseWithPath(input, Position{position.X + 1, position.Y}, target+1, knownPaths, path)
			TraverseWithPath(input, Position{position.X, position.Y - 1}, target+1, knownPaths, path)
			TraverseWithPath(input, Position{position.X, position.Y + 1}, target+1, knownPaths, path)
		}

		*path = (*path)[:len(*path)-1]
	}

}

func ComputeHeadPointWithPath(input [][]int, headPosition Position) int {
	knownPaths := [][]Position{}

	path := []Position{headPosition}
	TraverseWithPath(input, Position{headPosition.X - 1, headPosition.Y}, 1, &knownPaths, &path)
	TraverseWithPath(input, Position{headPosition.X + 1, headPosition.Y}, 1, &knownPaths, &path)
	TraverseWithPath(input, Position{headPosition.X, headPosition.Y - 1}, 1, &knownPaths, &path)
	TraverseWithPath(input, Position{headPosition.X, headPosition.Y + 1}, 1, &knownPaths, &path)

	return len(knownPaths)
}

func SolutionPartOne(grid [][]int) int {
	heads := []Position{}

	for x, row := range grid {
		for y, cell := range row {
			if cell == 0 {
				heads = append(heads, Position{x, y})
			}
		}
	}

	result := 0
	for _, head := range heads {
		result += ComputeHeadPoint(grid, head)
	}

	return result
}

func SolutionPartTwo(grid [][]int) int {
	heads := []Position{}

	for x, row := range grid {
		for y, cell := range row {
			if cell == 0 {
				heads = append(heads, Position{x, y})
			}
		}
	}

	result := 0
	for _, head := range heads {
		result += ComputeHeadPointWithPath(grid, head)
	}

	return result
}
