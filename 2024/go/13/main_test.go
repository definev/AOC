package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/13.txt")

	got := SolutionPartOne(input)
	want := 480

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/13.txt")

	got := SolutionPartOne(input)
	want := 37680

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 0

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
