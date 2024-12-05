package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/05.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 143

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 123

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/05.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 4462

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 6767

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
