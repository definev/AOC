package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	X int
	Y int
}

type Robot struct {
	P Pair
	V Pair
}

func ForceInt(raw string) int {
	result, _ := strconv.Atoi(raw)
	return result
}

func NewLine(raw string) Robot {
	splits := strings.Split(strings.Split(raw, "p=")[1], " v=")

	P := Pair{
		X: ForceInt(strings.Split(splits[0], ",")[0]),
		Y: ForceInt(strings.Split(splits[0], ",")[1]),
	}
	V := Pair{
		X: ForceInt(strings.Split(splits[1], ",")[0]),
		Y: ForceInt(strings.Split(splits[1], ",")[1]),
	}

	return Robot{P, V}
}

func ParseInput(filepath string) []Robot {
	data, _ := os.ReadFile(filepath)

	result := []Robot{}

	for _, line := range bytes.Split(data, []byte{'\n'}) {
		result = append(result, NewLine(string(line)))
	}

	return result
}

func Compute(robots []Robot, maxX int, maxY int, step int) int {
	newRobots := make([]Robot, len(robots))
	copy(newRobots, robots)

	stepCount := 0

	for stepCount < step {
		for index, robot := range newRobots {
			robot.P.X += robot.V.X
			robot.P.Y += robot.V.Y

			if robot.P.X < 0 {
				robot.P.X += maxX
			} else if robot.P.X >= maxX {
				robot.P.X -= maxX
			}

			if robot.P.Y < 0 {
				robot.P.Y += maxY
			} else if robot.P.Y >= maxY {
				robot.P.Y -= maxY
			}

			newRobots[index] = robot
		}
		stepCount += 1
	}

	return CountQuadrant(newRobots, maxX, maxY)
}

func VisualizeMap(maxX int, maxY int, robots []Robot) {
	visualize := make([][]int, 0)
	for y := 0; y < maxY; y += 1 {
		row := make([]int, maxX)
		visualize = append(visualize, row)
	}

	for _, robot := range robots {
		visualize[robot.P.Y][robot.P.X] += 1
	}

	for _, row := range visualize {
		for _, char := range row {
			if char == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}

		fmt.Println()
	}
}

func CountQuadrant(robots []Robot, maxX int, maxY int) int {
	quarter := [4]int{}

	halfX := maxX / 2
	halfY := maxY / 2

	for _, robot := range robots {
		if robot.P.X < halfX && robot.P.Y < halfY {
			quarter[0] += 1
		}
		if robot.P.X > halfX && robot.P.Y > halfY {
			quarter[1] += 1
		}
		if robot.P.X < halfX && robot.P.Y > halfY {
			quarter[2] += 1
		}
		if robot.P.X > halfX && robot.P.Y < halfY {
			quarter[3] += 1
		}
	}

	return quarter[0] * quarter[1] * quarter[2] * quarter[3]
}

func SolutionPartOne(input []Robot) int {
	maxX := 0
	maxY := 0
	for _, robot := range input {
		if robot.P.X > maxX {
			maxX = robot.P.X
		}
		if robot.P.Y > maxY {
			maxY = robot.P.Y
		}
	}
	maxX += 1
	maxY += 1

	return Compute(input, maxX, maxY, 100)
}

func ComputeTwo(robots []Robot, maxX int, maxY int) int {
	newRobots := make([]Robot, len(robots))
	copy(newRobots, robots)

	stepCount := 0

	for stepCount < 10000 {
		fmt.Printf("At step %d\n", stepCount)
		VisualizeMap(maxX, maxY, newRobots)
		for index, robot := range newRobots {
			robot.P.X += robot.V.X
			robot.P.Y += robot.V.Y

			if robot.P.X < 0 {
				robot.P.X += maxX
			} else if robot.P.X >= maxX {
				robot.P.X -= maxX
			}

			if robot.P.Y < 0 {
				robot.P.Y += maxY
			} else if robot.P.Y >= maxY {
				robot.P.Y -= maxY
			}

			newRobots[index] = robot
		}
		fmt.Println()
		stepCount += 1
	}

	return CountQuadrant(newRobots, maxX, maxY)
}

func SolutionPartTwo(input []Robot) int {
	maxX := 0
	maxY := 0
	for _, robot := range input {
		if robot.P.X > maxX {
			maxX = robot.P.X
		}
		if robot.P.Y > maxY {
			maxY = robot.P.Y
		}
	}
	maxX += 1
	maxY += 1

	return ComputeTwo(input, maxX, maxY)
}
