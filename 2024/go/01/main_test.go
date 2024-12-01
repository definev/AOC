package main

import (
	"testing"
)

func TestSample(t *testing.T) {
	file := "../../sample/01.txt"
	lefts, rights := ParseInput(file)

	got := SolutionPartOne(lefts[:], rights[:])
	want := 11

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SolutionPartTwo(lefts[:], rights[:])
	want = 31
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	file := "../../input/01.txt"
	lefts, rights := ParseInput(file)

	got := SolutionPartOne(lefts, rights)
	want := 2164381

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

	got = SolutionPartTwo(lefts[:], rights[:])
	want = 20719933
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
