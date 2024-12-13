package main

import "testing"

func TestSample(t *testing.T) {
	filepath := "../../sample/12.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 140

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 80

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/12.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 1457298

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 9

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
