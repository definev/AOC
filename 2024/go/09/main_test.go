package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	filepath := "../../sample/09.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 1928

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 2858

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/09.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 6310675819476

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 6335972980679
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
