package main

import "testing"

// func TestSample(t *testing.T) {
// 	input := ParseInput("../../sample/14.txt")

// 	got := SolutionPartOne(input)
// 	want := 12

// 	if got != want {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

func TestInput(t *testing.T) {
	input := ParseInput("../../input/14.txt")

	got := SolutionPartOne(input)
	want := 12

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 12

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
