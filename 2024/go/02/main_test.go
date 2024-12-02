package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/02.txt"
	input := ParseInput(filepath)
	got := SolutionPartOne(input[:])
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SolutionPartTwo(input[:])
	want = 4
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/02.txt"
	input := ParseInput(filepath)
	got := SolutionPartOne(input[:])
	want := 624
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SolutionPartTwo(input[:])
	want = 658
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
