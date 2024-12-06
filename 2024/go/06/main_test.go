package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/06.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 41

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/06.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 5086

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 1770
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
