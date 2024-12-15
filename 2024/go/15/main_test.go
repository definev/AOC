package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/15.txt")
	part2Input := inflateMap(input)

	got := SolutionPartOne(input)
	want := 10092

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(part2Input)
	want = 9021

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/15.txt")
	part2Input := inflateMap(input)

	got := SolutionPartOne(input)
	want := 1398947

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(part2Input)
	want = 1398947

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
