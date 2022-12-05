package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tuple[F, S any] struct {
	first  F
	second S
}

type CrateStack = []string

type Command struct {
	quantity int
	from     int
	to       int
}

func ParseCommand(input string) Command {
	raw := strings.Split(input, " ")
	quantity, _ := strconv.Atoi(raw[1])
	from, _ := strconv.Atoi(raw[3])
	to, _ := strconv.Atoi(raw[5])

	return Command{
		quantity: quantity,
		from:     from - 1,
		to:       to - 1,
	}
}

func parseCrates(input string) []CrateStack {
	raw := strings.Split(input, "\n")
	lengthStr := strings.Split(strings.TrimSpace(raw[len(raw)-1]), " ")
	length, _ := strconv.Atoi(lengthStr[len(lengthStr)-1])
	raw = raw[:len(raw)-1]

	results := make([]CrateStack, length)
	for i := len(raw) - 1; i >= 0; i -= 1 {
		line := raw[i]
		for j := 0; j < length; j += 1 {
			crate := line[j*4 : j*4+3]
			if crate != "   " {
				results[j] = append(results[j], crate)
			}
		}
	}

	return results
}

func parseCommands(input string) []Command {
	raw := strings.Split(input, "\n")

	results := make([]Command, 0)
	for _, v := range raw {
		results = append(results, ParseCommand(v))
	}

	return results
}

func parseInput(fileName string) Tuple[[]CrateStack, []Command] {
	file, _ := os.ReadFile(fileName)
	raw := strings.Split(string(file), "\n\n")

	crates := parseCrates(raw[0])
	commands := parseCommands(raw[1])

	return Tuple[[]CrateStack, []Command]{
		first:  crates,
		second: commands,
	}
}

func main() {
	firstHalfProblem("../../input/day05.txt")
	lastHalfProblem("../../input/day05.txt")
}

func reverse[T any](numbers []T) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

func _execute(stacks []CrateStack, command Command, reversed bool) {
	from := stacks[command.from]
	moveStack := from[len(from)-command.quantity:]
	if reversed {
		reverse(moveStack)
	}

	stacks[command.from] = from[:len(from)-command.quantity]
	stacks[command.to] = append(stacks[command.to], moveStack...)
}

func getHead(stacks []CrateStack) string {
	results := make(CrateStack, 0)
	for _, v := range stacks {
		results = append(results, string(v[len(v) - 1][1]))
	}

	return strings.Join(results, "")
}

func firstHalfProblem(fileName string) {
	input := parseInput(fileName)

	for _, v := range input.second {
		_execute(input.first, v, true)
	}

	fmt.Println(getHead(input.first))
}

func lastHalfProblem(fileName string) {
	input := parseInput(fileName)

	for _, v := range input.second {
		_execute(input.first, v, false)
	}

	fmt.Println(getHead(input.first))
}
