package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/07.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 3749

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 11387

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/07.txt"
	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 28730327770375

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 424977609625985

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
