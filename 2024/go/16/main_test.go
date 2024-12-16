package main

import "testing"

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/16.txt")

	got := SolutionPartOne(input)
	want := 7036

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 45

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/16.txt")

	got := SolutionPartOne(input)
	want := 95444

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 513

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
