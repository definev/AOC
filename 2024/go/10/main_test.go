package main

import "testing"

func TestSample(t *testing.T) {
	filepath := "../../sample/10.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 36

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 81

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/10.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 737

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 1619

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
