package main

import "testing"

func TestSample(t *testing.T) {
	filepath := "../../sample/11.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 55312

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 65601038650482

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/11.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 218956

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 259593838049805

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
