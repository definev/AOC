package main

import "testing"

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/19.txt")

	got := SolutionPartOne(input)
	want := 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 16

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/19.txt")

	got := SolutionPartOne(input)
	want := 260

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 639963796864990

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
