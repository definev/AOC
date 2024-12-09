package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/08.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 14

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 34

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/08.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 354

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 1263

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
