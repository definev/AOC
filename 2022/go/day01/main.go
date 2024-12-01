package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Tuple[T, U any] struct {
	First  T
	Second U
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput() [][]int {
	f, err := os.ReadFile("../../input/day01.txt")
	check(err)

	input := string(f)
	rawElf := strings.Split(input, "\n\n")

	elfs := make([][]int, 0)
	for _, raw := range rawElf {
		elf := make([]int, 0)
		elfRaw := strings.Split(raw, "\n")
		for _, elfCalRaw := range elfRaw {
			elfCal, err := strconv.ParseInt(strings.TrimSpace(elfCalRaw), 10, 0)
			check(err)
			elf = append(elf, int(elfCal))
		}

		elfs = append(elfs, elf)
	}

	return elfs
}

func main() {
	elfs := parseInput()
	firstHalfProblem(elfs)
	lastHalfProblem(elfs)
}
func firstHalfProblem(elfs [][]int) {
	best := 0

	for _, elf := range elfs {
		elfCals := 0
		for _, cal := range elf {
			elfCals += cal
		}

		if elfCals > best {
			best = elfCals
		}
	}

	fmt.Println(best)
}

func lastHalfProblem(elfs [][]int) {
	bestElfs := make([]int, 0)

	for _, elf := range elfs {
		elfCals := 0
		for _, cal := range elf {
			elfCals += cal
		}
		if len(bestElfs) < 3 {
			bestElfs = append(bestElfs, elfCals)
			continue
		}

		sort.Sort(sort.IntSlice(bestElfs))

		canReplace := Tuple[bool, int]{First: false, Second: -1}

		for i := 0; i < 3; i++ {
			if bestElfs[i] < elfCals {
				canReplace = Tuple[bool, int]{First: true, Second: i}
				break
			}
		}

		if canReplace.First {
			bestElfs[canReplace.Second] = elfCals
		}
	}

	total := 0
	for _, cal := range bestElfs {
		total += cal
	}

	fmt.Println(total)
}
