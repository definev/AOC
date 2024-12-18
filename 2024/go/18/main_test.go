package main

import "testing"

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/18.txt")

	got := SolutionPartOne(input, 12)
	want := 22

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got2 := SolutionPartTwo(input)
	want2 := Position{6, 1}

	if got2 != want2 {
		t.Errorf("got %v want %v", got2, want2)
	}
}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/18.txt")

	got := SolutionPartOne(input, 1024)
	want := 22

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got2 := SolutionPartTwo(input)
	want2 := Position{20, 64}

	if got2 != want2 {
		t.Errorf("got %v want %v", got2, want2)
	}
}
