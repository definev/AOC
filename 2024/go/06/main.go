package main

import (
	"fmt"
	"os"
	"strings"
)

type Position struct {
	X int
	Y int
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

type Maze struct {
	InitialPosition Position
	Obstacles       map[Position]bool
	Raw             [][]rune
	Height          int
	Width           int
}

func ParseInput(filepath string) Maze {
	data, _ := os.ReadFile(filepath)

	result := Maze{
		Obstacles: make(map[Position]bool),
	}

	lines := strings.Split(string(data), "\n")
	result.Height = len(lines)
	result.Width = len(lines[0])
	result.Raw = make([][]rune, result.Height)

	for x, line := range lines {
		for y, char := range line {
			result.Raw[x] = append(result.Raw[x], char)

			if char == '#' {
				result.Obstacles[Position{x, y}] = true
			}
			if char == '^' {
				result.InitialPosition = Position{x, y}
			}
		}
	}

	return result
}

func SolutionPartOne(input Maze) int {
	direction := UP
	currentPosition := input.InitialPosition

	stepsSet := make(map[Position]bool)
	stepsSet[currentPosition] = true

	for {
		switch direction {
		case UP:
			currentPosition.X--
		case RIGHT:
			currentPosition.Y++
		case DOWN:
			currentPosition.X++
		case LEFT:
			currentPosition.Y--
		}

		if currentPosition.X < 0 || currentPosition.X >= input.Height || currentPosition.Y < 0 || currentPosition.Y >= input.Width {
			break
		}

		stepsSet[currentPosition] = true
		for obstacle := range input.Obstacles {
			if currentPosition.X == obstacle.X && currentPosition.Y == obstacle.Y {
				delete(stepsSet, currentPosition)

				switch direction {
				case UP:
					direction = RIGHT
					currentPosition.X++
				case RIGHT:
					direction = DOWN
					currentPosition.Y--
				case DOWN:
					direction = LEFT
					currentPosition.X--
				case LEFT:
					direction = UP
					currentPosition.Y++
				}
				break
			}
		}
	}

	return len(stepsSet)
}

func ComputeDirection(a Position, b Position) Direction {
	if a.X < b.X {
		return UP
	}
	if a.X > b.X {
		return DOWN
	}
	if a.Y < b.Y {
		return LEFT
	}
	return RIGHT
}

type ObstaclePair struct {
	Obstacle  Position
	Direction Direction
}

func CheckDuplicate(path []Position) bool {
outer:
	for i := len(path) / 2; i >= 3; i-- {
		subPath := path[len(path)-i:]
		comparePath := path[len(path)-2*i : len(path)-i]
		if len(subPath) != len(comparePath) {
			continue
		}
		for j := 0; j < len(subPath); j++ {
			if subPath[j] != comparePath[j] {
				continue outer
			}
		}

		return true
	}

	return false
}

func IsLoopPath(input Maze, AdderObstacle Position) bool {
	direction := UP
	obstacles := make(map[Position]bool, len(input.Obstacles))
	for k, v := range input.Obstacles {
		obstacles[k] = v
	}
	obstacles[AdderObstacle] = true

	path := []Position{}
	currentPosition := input.InitialPosition
	path = append(path, currentPosition)
	for {
		currentPosition = updatePosition(currentPosition, direction)

		if currentPosition.X < 0 || currentPosition.X >= input.Height || currentPosition.Y < 0 || currentPosition.Y >= input.Width {
			break
		}

		if obstacles[currentPosition] {
			oppositeDirection := (direction + 2) % 4
			currentPosition = updatePosition(currentPosition, oppositeDirection)
			direction = (direction + 1) % 4
		} else {
			path = append(path, currentPosition)
		}

		if CheckDuplicate(path) {
			fmt.Printf("Adding obstacle: %v\n", AdderObstacle)
			return true
		}
	}

	return false
}

func updatePosition(position Position, direction Direction) Position {
	newPosition := Position{
		X: position.X,
		Y: position.Y,
	}
	switch direction {
	case UP:
		newPosition.X -= 1
	case RIGHT:
		newPosition.Y += 1
	case DOWN:
		newPosition.X += 1
	case LEFT:
		newPosition.Y -= 1
	}
	return newPosition
}

func SolutionPartTwo(input Maze) int {
	direction := UP
	currentPosition := input.InitialPosition

	obstaclePair := []ObstaclePair{}

	for {
		currentPosition = updatePosition(currentPosition, direction)

		if currentPosition.X < 0 || currentPosition.X >= input.Height || currentPosition.Y < 0 || currentPosition.Y >= input.Width {
			break
		}

		for input.Obstacles[currentPosition] {
			obstacle := currentPosition
			oppositeDirection := (direction + 2) % 4
			currentPosition = updatePosition(currentPosition, oppositeDirection)
			direction = (direction + 1) % 4

			obstaclePair = append(obstaclePair, ObstaclePair{obstacle, direction})
			break
		}
	}

	obstacleAdditional := map[Position]bool{}

	fmt.Printf("Length of obstaclePair: %d\n", len(obstaclePair))

	for i := -1; i < len(obstaclePair); i++ {
		first := Position{}
		if i == -1 {
			first = input.InitialPosition
		} else {
			direction = (obstaclePair[i].Direction - 1 + 2) % 4
			first = updatePosition(obstaclePair[i].Obstacle, direction)
		}
		fmt.Printf("At %d | result: %d\n", i, len(obstacleAdditional))

		second := Position{}
		if i == len(obstaclePair)-1 {
			switch obstaclePair[i].Direction {
			case UP:
				second.X = 0
				second.Y = first.Y
			case RIGHT:
				second.X = first.X
				second.Y = input.Width - 1
			case DOWN:
				second.X = input.Height - 1
				second.Y = first.Y
			case LEFT:
				second.X = first.X
				second.Y = 0
			}
		} else {
			second = obstaclePair[i+1].Obstacle
		}

		direction := UP
		if i != -1 {
			direction = obstaclePair[i].Direction
		}

		switch direction {
		case UP:
			for j := first.X - 1; j >= second.X; j-- {
				additionalObstacle := Position{X: j, Y: first.Y}
				if IsLoopPath(input, additionalObstacle) {
					obstacleAdditional[additionalObstacle] = true
				}
			}
		case DOWN:
			for j := first.X + 1; j <= second.X; j++ {
				additionalObstacle := Position{X: j, Y: first.Y}
				if IsLoopPath(input, additionalObstacle) {
					obstacleAdditional[additionalObstacle] = true
				}
			}
		case LEFT:
			for j := first.Y - 1; j >= second.Y; j-- {
				additionalObstacle := Position{X: first.X, Y: j}
				if IsLoopPath(input, additionalObstacle) {
					obstacleAdditional[additionalObstacle] = true
				}
			}
		case RIGHT:
			for j := first.Y + 1; j <= second.Y; j++ {
				additionalObstacle := Position{X: first.X, Y: j}
				if IsLoopPath(input, additionalObstacle) {
					obstacleAdditional[additionalObstacle] = true
				}
			}
		}
	}

	return len(obstacleAdditional)
}
