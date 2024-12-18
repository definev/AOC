package main

import "testing"

func TestSample(t *testing.T) {
	input := ParseInput("../../sample/17.txt")

	got := SolutionPartOne(input)
	want := "4,6,3,5,6,3,5,2,1,0"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got2 := SolutionPartTwo(input)
	want2 := 117440

	if got2 != want2 {
		t.Errorf("got %v want %v", got2, want2)
	}

}

func TestInput(t *testing.T) {
	input := ParseInput("../../input/17.txt")

	got := SolutionPartOne(input)
	want := "2,1,4,7,6,0,3,1,4"

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got2 := SolutionPartTwo(input)
	want2 := 266932601404433

	if got2 != want2 {
		t.Errorf("got %v want %v", got2, want2)
	}
}
