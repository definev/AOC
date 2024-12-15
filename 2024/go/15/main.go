package main

import (
	"bytes"
	"fmt"
	"os"
)

type Position struct {
	X int
	Y int
}

func (p Position) Move(deltaX int, deltaY int) Position {
	return Position{
		p.X + deltaX,
		p.Y + deltaY,
	}
}

func (p Position) MoveDirection(direction Direction) Position {
	switch direction {
	case UP:
		return p.Move(0, -1)
	case DOWN:
		return p.Move(0, 1)
	case LEFT:
		return p.Move(-1, 0)
	case RIGHT:
		return p.Move(1, 0)
	}
	return p
}

type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

type FishMap struct {
	Map             [][]rune
	InitialPosition Position
	Directions      []Direction
}

type InflateMap struct {
	FishMap
	tanglePositionMap map[Position]Position
}

func inflateMap(input FishMap) InflateMap {
	tanglePositionMap := make(map[Position]Position)
	newFishMap := FishMap{
		Map:        make([][]rune, len(input.Map)),
		Directions: input.Directions,
	}

	for y, row := range input.Map {
		newRow := make([]rune, len(row)*2)

		for x, char := range row {
			switch char {
			case 'O':
				newRow[x*2] = '['
				newRow[x*2+1] = ']'
				tanglePositionMap[Position{x * 2, y}] = Position{x*2 + 1, y}
				tanglePositionMap[Position{x*2 + 1, y}] = Position{x * 2, y}
			case '.':
				newRow[x*2] = '.'
				newRow[x*2+1] = '.'
			case '#':
				newRow[x*2] = '#'
				newRow[x*2+1] = '#'
			case '@':
				newRow[x*2] = '@'
				newRow[x*2+1] = '.'

				newFishMap.InitialPosition = Position{x * 2, y}
			}
		}

		newFishMap.Map[y] = newRow
	}

	return InflateMap{
		FishMap:           newFishMap,
		tanglePositionMap: tanglePositionMap,
	}
}

func ParseInput(filepath string) FishMap {
	result := FishMap{
		Map:        make([][]rune, 0),
		Directions: make([]Direction, 0),
	}

	data, _ := os.ReadFile(filepath)

	part := bytes.Split(data, []byte{'\n', '\n'})

	mapRaw := bytes.Split(part[0], []byte{'\n'})
	// map
	for y, line := range mapRaw {
		charLine := make([]rune, len(line))
		for x, char := range line {
			charLine[x] = rune(char)
			if charLine[x] == '@' {
				result.InitialPosition = Position{x, y}
			}
		}
		result.Map = append(result.Map, charLine)
	}

	// direction
	directionsRaw := part[1]
	for _, char := range directionsRaw {
		if rune(char) == '<' {
			result.Directions = append(result.Directions, LEFT)
		}
		if rune(char) == 'v' {
			result.Directions = append(result.Directions, DOWN)
		}
		if rune(char) == '>' {
			result.Directions = append(result.Directions, RIGHT)
		}
		if rune(char) == '^' {
			result.Directions = append(result.Directions, UP)
		}
	}

	return result
}

func IsValidPosition(nextPosition Position, input FishMap) bool {
	if nextPosition.X < 0 || nextPosition.X >= len(input.Map[0]) || nextPosition.Y < 0 || nextPosition.Y >= len(input.Map) {
		return false
	}
	return true
}

func resolveMap(input FishMap) FishMap {
	for _, direction := range input.Directions {
		nextPosition := input.InitialPosition.MoveDirection(direction)
		if !IsValidPosition(nextPosition, input) {
			continue
		}

		target := input.Map[nextPosition.Y][nextPosition.X]
		switch target {
		case '.':
			input.Map[input.InitialPosition.Y][input.InitialPosition.X] = '.'
			input.Map[nextPosition.Y][nextPosition.X] = '@'
			input.InitialPosition = nextPosition
		case '#':
			continue
			// Hit the wall
		case 'O':
			// Hit the tank
			potentialChangePosition := []Position{input.InitialPosition, nextPosition}
			potentialNextPosition := nextPosition.MoveDirection(direction)

			for IsValidPosition(potentialNextPosition, input) &&
				input.Map[potentialNextPosition.Y][potentialNextPosition.X] == 'O' {
				potentialChangePosition = append(potentialChangePosition, potentialNextPosition)
				potentialNextPosition = potentialNextPosition.MoveDirection(direction)
			}

			if IsValidPosition(potentialNextPosition, input) &&
				input.Map[potentialNextPosition.Y][potentialNextPosition.X] == '.' {
				potentialChangePosition = append(potentialChangePosition, potentialNextPosition)
			}
			if IsValidPosition(potentialNextPosition, input) &&
				input.Map[potentialNextPosition.Y][potentialNextPosition.X] == '#' {
				continue
			}
			if len(potentialChangePosition) < 3 {
				continue
			}

			for i := len(potentialChangePosition) - 1; i > 0; i -= 1 {
				positionI := potentialChangePosition[i]
				positionPrevI := potentialChangePosition[i-1]

				input.Map[positionI.Y][positionI.X] = input.Map[positionPrevI.Y][positionPrevI.X]
			}

			input.Map[input.InitialPosition.Y][input.InitialPosition.X] = '.'
			input.Map[nextPosition.Y][nextPosition.X] = '@'
			input.InitialPosition = nextPosition
		}

	}

	return input
}

func PrintMap(input [][]rune) {
	for _, row := range input {
		for _, char := range row {
			fmt.Printf("%c", char)
		}
		fmt.Println()
	}
}

func SolutionPartOne(input FishMap) int {
	result := 0
	newMap := resolveMap(input)
	for y, row := range newMap.Map {
		for x, char := range row {
			if char == 'O' {
				result += x + y*100
			}
		}
	}
	return result
}

func floodFill(
	result map[Position]rune,
	input InflateMap, position Position,
	allowanceChar []rune,
	traversed map[Position]bool,
	direction Direction,
) {
	if !IsValidPosition(position, input.FishMap) {
		return
	}
	char := input.Map[position.Y][position.X]
	isAllow := false
	for _, allowChar := range allowanceChar {
		if char == allowChar {
			isAllow = true
			break
		}
	}
	if !isAllow {
		return
	}
	if traversed[position] {
		return
	}

	tanglePosition := input.tanglePositionMap[position]
	tangleChar := input.Map[tanglePosition.Y][tanglePosition.X]

	result[position] = char
	result[tanglePosition] = tangleChar
	traversed[position] = true
	traversed[tanglePosition] = true

	floodFill(
		result,
		input, position.MoveDirection(direction),
		allowanceChar,
		traversed,
		direction,
	)
	floodFill(
		result,
		input, tanglePosition.MoveDirection(direction),
		allowanceChar,
		traversed,
		direction,
	)
}

func resolveInflateMap(input InflateMap) FishMap {
	for _, direction := range input.Directions {
		nextPosition := input.InitialPosition.MoveDirection(direction)
		if !IsValidPosition(nextPosition, input.FishMap) {
			continue
		}
		target := input.Map[nextPosition.Y][nextPosition.X]

		boxMove := func() bool {
			potentialChangePositionMap := map[Position]rune{
				input.InitialPosition: '@',
			}

			floodFill(
				potentialChangePositionMap,
				input, nextPosition,
				[]rune{'[', ']'},
				map[Position]bool{},
				direction,
			)

			changePositionMap := map[Position]rune{}
			for position, char := range potentialChangePositionMap {
				nextPosition := position.MoveDirection(direction)
				if !IsValidPosition(nextPosition, input.FishMap) ||
					input.Map[nextPosition.Y][nextPosition.X] == '#' {
					return false
				}

				changePositionMap[nextPosition] = char
			}

			replaceTanglePositionMap := map[Position]Position{}

			for position := range potentialChangePositionMap {
				if input.Map[position.Y][position.X] == '[' || input.Map[position.Y][position.X] == ']' {
					oldTanglePosition := input.tanglePositionMap[position]
					delete(input.tanglePositionMap, position)
					replaceTanglePositionMap[position.MoveDirection(direction)] = oldTanglePosition.MoveDirection(direction)
				}

				input.Map[position.Y][position.X] = '.'
			}

			for position, tanglePosition := range replaceTanglePositionMap {
				input.tanglePositionMap[position] = tanglePosition
			}

			for position, char := range changePositionMap {
				input.Map[position.Y][position.X] = char
				if char == '@' {
					input.InitialPosition = position
				}
			}
			// fmt.Printf("Move %s\n", directionString(direction))
			// PrintMap(input.Map)

			return true
		}

		switch target {
		case '#':
			continue
		case '.':
			input.Map[input.InitialPosition.Y][input.InitialPosition.X] = '.'
			input.Map[nextPosition.Y][nextPosition.X] = '@'
			input.InitialPosition = nextPosition

			// fmt.Printf("Move %s\n", directionString(direction))
			// PrintMap(input.Map)
		case '[':
			boxMove()
		case ']':
			boxMove()
		}
	}

	return input.FishMap
}

func SolutionPartTwo(input InflateMap) int {
	result := 0
	// PrintMap(input.Map)
	newMap := resolveInflateMap(input)
	for y, row := range newMap.Map {
		for x, char := range row {
			if char == '[' {
				result += x + y*100
			}
		}
	}
	return result
}
