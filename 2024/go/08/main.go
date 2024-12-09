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

type AntennaMap struct {
	Raw      [][]rune
	Antennas map[rune][]Position
}

func ParseInput(filepath string) AntennaMap {
	data, _ := os.ReadFile(filepath)
	lines := strings.Split(string(data), "\n")

	result := AntennaMap{
		Raw:      make([][]rune, 0),
		Antennas: map[rune][]Position{},
	}

	for x, line := range lines {
		result.Raw = append(result.Raw, []rune(line))

		for y, r := range line {
			if r != '.' {
				if _, ok := result.Antennas[r]; !ok {
					result.Antennas[r] = make([]Position, 0)
				}

				result.Antennas[r] = append(result.Antennas[r], Position{x, y})
			}
		}
	}

	return result
}

func SolutionPartOne(input AntennaMap) int {
	waveMap := map[Position]bool{}

	validWave := func(p Position) bool {
		if p.X < 0 || p.X >= len(input.Raw) {
			return false
		}
		if p.Y < 0 || p.Y >= len(input.Raw[p.X]) {
			return false
		}

		return true
	}

	for _, positions := range input.Antennas {
		for i := 0; i < len(positions)-1; i += 1 {
			for j := i + 1; j < len(positions); j += 1 {
				p1 := positions[i]
				p2 := positions[j]

				diff := Position{p2.X - p1.X, p2.Y - p1.Y}

				o1 := Position{p1.X - diff.X, p1.Y - diff.Y}
				o2 := Position{p2.X + diff.X, p2.Y + diff.Y}

				if validWave(o1) {
					waveMap[o1] = true
				}
				if validWave(o2) {
					waveMap[o2] = true
				}
			}
		}
	}

	for key := range waveMap {
		input.Raw[key.X][key.Y] = '#'
	}

	return len(waveMap)
}

func SolutionPartTwo(input AntennaMap) int {
	waveMap := map[Position]bool{}

	for _, positions := range input.Antennas {
		for i := 0; i < len(positions)-1; i += 1 {
			for j := i + 1; j < len(positions); j += 1 {
				p1 := positions[i]
				p2 := positions[j]

				diff := Position{p2.X - p1.X, p2.Y - p1.Y}
				o1 := Position{p1.X, p1.Y}
				waveMap[o1] = true

				for {
					o1 = Position{o1.X - diff.X, o1.Y - diff.Y}
					if o1.X < 0 || o1.X >= len(input.Raw) || o1.Y < 0 || o1.Y >= len(input.Raw[o1.X]) {
						break
					}

					waveMap[o1] = true
				}

				o2 := Position{p2.X, p2.Y}
				waveMap[o2] = true

				for {
					o2 = Position{o2.X + diff.X, o2.Y + diff.Y}
					if o2.X < 0 || o2.X >= len(input.Raw) || o2.Y < 0 || o2.Y >= len(input.Raw[o2.X]) {
						break
					}

					waveMap[o2] = true
				}

			}
		}
	}

	for key := range waveMap {
		input.Raw[key.X][key.Y] = '#'
	}

	return len(waveMap)
}

func PrintMap(input [][]rune) {
	for _, row := range input {
		for _, r := range row {
			fmt.Print(string(r))
		}
		fmt.Println()
	}
}
