package main

import "testing"

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/20.txt")

	got := SolutionPartOne(input, 0)
	want := 44

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input, 50)
	want = 3

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/20.txt")

	got := SolutionPartOne(input, 100)
	want := 1263

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input, 100)
	want = 957831

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
