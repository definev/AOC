package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Code int

const (
	ADV Code = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type Op struct {
	Code Code
	Inst int
}

type Input struct {
	RegA    uint
	RegB    uint
	RegC    uint
	Program []int
}

func ParseInt(raw string) int {
	n, _ := strconv.Atoi(raw)
	return n
}

func ParseInput(filepath string) Input {
	input := Input{}

	data, _ := os.ReadFile(filepath)
	lines := strings.Split(string(data), "\n")

	input.RegA = uint(ParseInt(strings.Split(lines[0], ": ")[1]))
	input.RegB = uint(ParseInt(strings.Split(lines[1], ": ")[1]))
	input.RegC = uint(ParseInt(strings.Split(lines[2], ": ")[1]))

	input.Program = make([]int, 0)
	raw := strings.Split(strings.Split(lines[4], ": ")[1], ",")
	for i := 0; i < len(raw); i += 1 {
		input.Program = append(input.Program, ParseInt(raw[i]))
	}

	return input
}

func GetLiteralFromInst(input Input, inst int) uint {
	return uint(inst)
}

func GetComboFromInst(input Input, inst int) uint {
	switch inst {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 3
	case 4:
		return input.RegA
	case 5:
		return input.RegB
	case 6:
		return input.RegC
	case 7:
		return 7
	}
	return 0
}

func Interpret(input *Input, debug bool) []string {
	var ip uint = 0
	output := make([]string, 0)

	for ip < uint(len(input.Program)-1) {
		op := input.Program[ip]
		combo := GetComboFromInst(*input, input.Program[ip+1])
		literal := GetLiteralFromInst(*input, input.Program[ip+1])

		switch Code(op) {
		case BXL:
			input.RegB = input.RegB ^ literal
			ip += 2
		case BST:
			input.RegB = combo % 8
			ip += 2
		case JNZ:
			if input.RegA != 0 {
				ip = literal
			} else {
				ip += 2
			}
		case BXC:
			input.RegB = input.RegB ^ input.RegC
			ip += 2
		case OUT:
			if debug {
				DumpVM(*input)
			}
			output = append(output, fmt.Sprintf("%d", combo%8))
			ip += 2
		case ADV:
			input.RegA = input.RegA >> combo
			ip += 2
		case BDV:
			input.RegB = input.RegA >> combo
			ip += 2
		case CDV:
			input.RegC = input.RegA >> combo
			ip += 2
		}
	}

	return output
}

func DumpVM(input Input) {
	fmt.Printf("RegA = %20d | RegB = %20d | RegC = %20d\n", input.RegA, input.RegB, input.RegC)
}

func SolutionPartOne(input Input) string {
	return strings.Join(Interpret(&input, false), ",")
}

func FindValidSolution(input Input, possible *[]int, digit int) {
	program := input.Program[len(input.Program)-digit]

	var i uint
	for i = 0; i < 8; i += 1 {
		newInput := &Input{
			RegA:    input.RegA + uint(i),
			RegB:    0,
			RegC:    0,
			Program: input.Program,
		}
		output := Interpret(newInput, false)

		if output[0] == fmt.Sprintf("%d", program) {
			if digit == len(newInput.Program) {
				*possible = append(*possible, int(input.RegA+i))
			}
			if digit < len(newInput.Program) {
				newInput.RegA = (input.RegA + i) << 3
				FindValidSolution(*newInput, possible, digit+1)
			}
		}
	}
}

func SolutionPartTwo(input Input) int {
	possible := make([]int, 0)
	FindValidSolution(Input{
		RegA:    0,
		RegB:    0,
		RegC:    0,
		Program: input.Program,
	}, &possible, 1)
	min := math.MaxInt
	for _, p := range possible {
		if p < min {
			min = p
		}
	}
	return min

}
