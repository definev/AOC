package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

func ParseInput(filepath string) []int {
	data, _ := os.ReadFile(filepath)

	result := []int{}

	for _, number := range bytes.Split(data, []byte(" ")) {
		num, _ := strconv.Atoi(string(number))
		result = append(result, num)
	}

	return result
}

func ApplyRules(stone int) Output {
	var result Output

	if stone == 0 {
		result = Output{Left: 1, Right: nil}
	} else {
		digitCount := int(math.Floor(math.Log10(float64(stone)))) + 1
		if digitCount%2 == 0 {
			left := stone / int(math.Pow10(digitCount/2))
			right := stone % int(math.Pow10(digitCount/2))

			result = Output{Left: left, Right: &right}
		} else {
			result = Output{Left: stone * 2024, Right: nil}
		}
	}

	return result
}

type Output struct {
	Left  int
	Right *int
}

func Expand(cache *sync.Map, stone int, iteration int, targetIteration int, total chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	runnedIteration := targetIteration - iteration
	output := ApplyRules(stone)

	if runnedIteration > 4 {
		if value, ok := cache.Load(Iteration{Stone: stone, Iteration: runnedIteration}); ok {
			// fmt.Printf("Found in cache, Iteration %d, Stone %d, Total %d\n", iteration, stone, value.(int))
			total <- value.(int)
			return
		}
	}

	if iteration == targetIteration-1 {
		if output.Right == nil {
			total <- 1
		} else {
			total <- 2
		}
		return
	}

	subWg := sync.WaitGroup{}
	subTotal := make(chan int, 2)

	subWg.Add(1)
	go Expand(cache, output.Left, iteration+1, targetIteration, subTotal, &subWg)
	if output.Right != nil {
		subWg.Add(1)
		go Expand(cache, *output.Right, iteration+1, targetIteration, subTotal, &subWg)
	}

	subWg.Wait()
	close(subTotal)
	result := 0
	for total := range subTotal {
		result += total
	}

	if runnedIteration > 4 {
		cache.Store(Iteration{Stone: stone, Iteration: runnedIteration}, result)
	}

	if iteration < 1 {
		fmt.Printf("Iteration %d, Stone %d, Total %d\n", iteration, stone, result)
	}

	total <- result
}

type Iteration struct {
	Stone     int
	Iteration int
}

func Solve(input []int, iteration int) int {
	totalCh := make(chan int, len(input))
	wg := sync.WaitGroup{}
	wg.Add(len(input))

	cache := sync.Map{}

	for _, number := range input {
		go Expand(&cache, number, 0, iteration, totalCh, &wg)
	}
	wg.Wait()
	close(totalCh)

	result := 0
	for total := range totalCh {
		result += total
	}

	return result
}

func SolutionPartOne(input []int) int {
	return Solve(input, 25)
}

func SolutionPartTwo(input []int) int {
	return Solve(input, 75)
}
