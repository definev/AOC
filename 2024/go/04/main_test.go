package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/04.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 18

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 9
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filename := "../../input/04.txt"

	input := ParseInput(filename)

	got := SolutionPartOne(input)
	want := 2554

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 1916
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
