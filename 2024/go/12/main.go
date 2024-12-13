package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func ParseInput(filepath string) [][]rune {
	data, _ := os.ReadFile(filepath)

	var grid [][]rune

	for _, line := range strings.Split(string(data), "\n") {
		grid = append(grid, []rune(line))
	}

	return grid
}

type Position struct {
	X int
	Y int
}

func (p Position) Move(x int, y int) Position {
	return Position{p.X + x, p.Y + y}
}

type Area struct {
	positions map[Position]bool
	char      rune
}

func (a Area) Inflate(unit int) Area {
	newArea := Area{
		positions: map[Position]bool{},
		char:      a.char,
	}

	for key := range a.positions {
		for i := 0; i < unit; i++ {
			for j := 0; j < unit; j++ {

				newArea.positions[Position{
					X: key.X*unit + i,
					Y: key.Y*unit + j,
				}] = true
			}
		}
	}

	return newArea
}

func countRelate(position Position, areas map[Position]bool) int {
	up := Position{position.X - 1, position.Y}
	left := Position{position.X, position.Y - 1}
	down := Position{position.X + 1, position.Y}
	right := Position{position.X, position.Y + 1}

	existSide := 0
	for _, side := range []Position{up, left, down, right} {
		if areas[side] {
			existSide += 1
		}
	}

	return existSide
}

func ComputePerimeter(area Area) int {
	count := 0

	for key := range area.positions {
		existSide := countRelate(key, area.positions)
		count += 4 - existSide
	}

	return count
}

func CheckOutOfBound(input [][]rune, position Position) bool {
	if position.X >= len(input) || position.X < 0 || position.Y >= len(input[0]) || position.Y < 0 {
		return true
	}
	return false
}

func FloodFill(
	input [][]rune,
	marked [][]bool,
	area *Area, target rune,
	position Position,
	nextPosition *Position,
) {
	if CheckOutOfBound(input, position) {
		return
	}
	cell := input[position.X][position.Y]
	if cell == target && !marked[position.X][position.Y] {
		marked[position.X][position.Y] = true
		(*area).positions[position] = true
		FloodFill(input, marked, area, target, position.Move(-1, 0), nextPosition)
		FloodFill(input, marked, area, target, position.Move(1, 0), nextPosition)
		FloodFill(input, marked, area, target, position.Move(0, -1), nextPosition)
		FloodFill(input, marked, area, target, position.Move(0, 1), nextPosition)
	} else {
		if !marked[position.X][position.Y] {
			*nextPosition = position
		}
	}
}

func FindAreas(input [][]rune) []Area {
	var areas []Area

	var marked [][]bool = make([][]bool, len(input))
	for m := 0; m < len(input); m++ {
		markedLine := make([]bool, len(input[0]))
		for n := 0; n < len(input[0]); n++ {
			markedLine[n] = false
		}
		marked[m] = markedLine
	}

	scanned := 0
	currentPosition := Position{0, 0}
	missCacheCount := 0

	for scanned < len(input)*len(input[0]) {
		if marked[currentPosition.X][currentPosition.Y] {
			missCacheCount += 1
		outer:
			for i := 0; i < len(input); i++ {
				for j := 0; j < len(input[0]); j++ {
					if !marked[i][j] {
						currentPosition = Position{i, j}
						break outer
					}
				}
			}
			if marked[currentPosition.X][currentPosition.Y] {
				break
			}
		}

		area := Area{
			positions: map[Position]bool{},
			char:      input[currentPosition.X][currentPosition.Y],
		}

		FloodFill(
			input,
			marked,
			&area, input[currentPosition.X][currentPosition.Y],
			currentPosition, &currentPosition,
		)

		if len(area.positions) > 0 {
			areas = append(areas, area)
			scanned += len(area.positions)
		}
	}

	fmt.Printf("Miss cache: %d\n", missCacheCount)

	return areas
}

func SolutionPartOne(input [][]rune) int {
	result := 0
	areas := FindAreas(input)

	for _, area := range areas {
		perimeter := ComputePerimeter(area)
		// fmt.Printf("Char: %c Area %d Perimeter: %d\n", area.char, len(area.position), perimeter)
		result += len(area.positions) * perimeter
	}

	return result
}

func countContinousLine(positions []Position, axis func(Position) int) int {
	slices.SortFunc(positions, func(a, b Position) int {
		return axis(a) - axis(b)
	})

	if len(positions) < 3 {
		return 0
	}

	count := 0
	lineCount := 1
	for i := 0; i < len(positions)-1; {
		diff := axis(positions[i+1]) - axis(positions[i])
		if diff == 1 {
			lineCount += 1
			i += 1
		} else {
			if lineCount >= 3 {
				count += 1
			}
			lineCount = 1
			i += 1
		}
	}

	if lineCount >= 3 {
		count += 1
	}

	return count
}

func cutLine(positionList []Position, axis func(Position) int) [][]Position {
	slices := [][]Position{}
	current := positionList[0]
	index := 1
	slice := []Position{current}
	for index < len(positionList) {
		if axis(positionList[index]) == axis(current) {
			slice = append(slice, positionList[index])
		} else {
			current = positionList[index]
			newSlice := make([]Position, len(slice))
			copy(newSlice, slice)
			slices = append(slices, newSlice)
			slice = []Position{current}
		}
		index += 1
	}
	slices = append(slices, slice)

	return slices
}

func computeSize(leanArea Area) int {
	// fmt.Printf("Count area of: %c\n", leanArea.char)

	vert := 0
	hori := 0

	positionList := make([]Position, 0)
	for position := range leanArea.positions {
		positionList = append(positionList, position)
	}
	// Count verticle line
	slices.SortFunc(positionList, func(a, b Position) int { return a.Y - b.Y })
	sl := cutLine(positionList, func(p Position) int { return p.Y })

	for _, slice := range sl {
		newVertLine := countContinousLine(slice, func(p Position) int { return p.X })
		vert += newVertLine
		// fmt.Printf("New vert line: %v | %v\n", slice, newVertLine)
	}

	// Count horizontal line
	slices.SortFunc(positionList, func(a, b Position) int { return a.X - b.X })
	sl = cutLine(positionList, func(p Position) int { return p.X })

	for _, slice := range sl {
		newHoriLine := countContinousLine(slice, func(p Position) int { return p.Y })
		vert += newHoriLine
		// fmt.Printf("New hori line: %v | %v\n", slice, newHoriLine)
	}

	return hori + vert
}

func InflateAndLeanAreas(areas []Area) []Area {
	inflatedAreas := []Area{}
	for _, area := range areas {
		inflate := area.Inflate(3)
		inflatedAreas = append(inflatedAreas, inflate)
	}

	leanAreas := []Area{}
	for _, area := range inflatedAreas {
		leanArea := Area{
			positions: map[Position]bool{},
			char:      area.char,
		}
		for position := range area.positions {
			existSide := countRelate(position, area.positions)
			if existSide != 4 {
				leanArea.positions[position] = true
			}
		}

		leanAreas = append(leanAreas, leanArea)
	}

	return leanAreas
}

func SolutionPartTwo(input [][]rune) int {
	result := 0
	areas := FindAreas(input)
	leanAreas := InflateAndLeanAreas(areas)

	for index := range leanAreas {
		area := areas[index]
		side := computeSize(leanAreas[index])
		result += side * len(area.positions)
	}

	return result
}
