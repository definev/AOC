package main

import "testing"

func TestSample(t *testing.T) {
	filepathOne := "../../sample/03_1.txt"
	filepathTwo := "../../sample/03_2.txt"

	input := ParseInput(filepathOne)

	got := SolutionPartOne(input)
	want := 161

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	input = ParseInput(filepathTwo)

	got = SolutionPartTwo(input)
	want = 48
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestInput(t *testing.T) {
	filepath := "../../input/03.txt"

	input := ParseInput(filepath)

	got := SolutionPartOne(input)
	want := 174561379

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	got = SolutionPartTwo(input)
	want = 174521395
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
