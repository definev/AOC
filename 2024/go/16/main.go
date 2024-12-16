package main

import (
	"bytes"
	"maps"
	"math"
	"os"

	"container/heap"
)

// PriorityQueue implements heap.Interface
type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*State)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Position struct {
	X int
	Y int
}

func (p Position) Move(deltaX int, deltaY int) Position {
	return Position{
		p.X + deltaX,
		p.Y + deltaY,
	}
}

func (p Position) MoveDirection(direction Direction) Position {
	switch direction {
	case UP:
		return p.Move(0, -1)
	case DOWN:
		return p.Move(0, 1)
	case LEFT:
		return p.Move(-1, 0)
	case RIGHT:
		return p.Move(1, 0)
	}
	return p
}

func (p Position) OutOfBound(width int, height int) bool {
	return p.X < 0 || p.Y < 0 || p.X >= width || p.Y >= height
}

type Direction int

const (
	UP Direction = iota
	LEFT
	DOWN
	RIGHT
)

type World struct {
	Map   [][]rune
	Start Position
	End   Position
}

func ParseInput(filepath string) World {
	data, _ := os.ReadFile(filepath)

	result := World{
		Map: make([][]rune, 0),
	}

	for y, line := range bytes.Split(data, []byte("\n")) {
		result.Map = append(result.Map, bytes.Runes(line))
		for x, r := range bytes.Runes(line) {
			if r == 'E' {
				result.End = Position{x, y}
			}
			if r == 'S' {
				result.Start = Position{x, y}
			}
		}
	}

	return result
}

type State struct {
	Coordinate
	cost int
	path map[Position]bool
}

type Coordinate struct {
	position  Position
	direction Direction
}

func FindShortestPath(input World) (int, map[Position]bool) {
	height, width := len(input.Map), len(input.Map[0])

	maxScore := math.MaxInt
	finalPath := map[Position]bool{
		input.Start: true,
	}

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{
		cost: 0,
		Coordinate: Coordinate{
			position:  input.Start,
			direction: RIGHT,
		},
		path: map[Position]bool{},
	})

	dist := map[Coordinate]int{}

	for pq.Len() > 0 {
		state := heap.Pop(pq).(*State)
		direction, cost, coordinate := state.direction, state.cost, state.Coordinate

		if score, ok := dist[coordinate]; ok && score < state.cost {
			continue
		}

		dist[coordinate] = state.cost
		if state.position == input.End && state.cost <= maxScore {
			maxScore = state.cost
			maps.Copy(finalPath, state.path)
		}

		// move forward
		newPosition := state.position.MoveDirection(direction)

		if !newPosition.OutOfBound(width, height) && input.Map[state.position.Y][state.position.X] != '#' {
			newPath := maps.Clone(state.path)
			newPath[newPosition] = true
			heap.Push(pq, &State{
				cost: cost + 1,
				Coordinate: Coordinate{
					position:  newPosition,
					direction: direction,
				},
				path: newPath,
			})
		}

		// move clockwise
		clockwiseDirection := (state.direction + 1) % 4
		newPosition = state.position.MoveDirection(clockwiseDirection)
		if !newPosition.OutOfBound(width, height) && input.Map[newPosition.Y][newPosition.X] != '#' {
			newPath := maps.Clone(state.path)
			newPath[newPosition] = true
			heap.Push(pq, &State{
				cost: cost + 1000,
				Coordinate: Coordinate{
					position:  coordinate.position,
					direction: clockwiseDirection,
				},
				path: newPath,
			})
		}

		// move counter-clockwise
		counterClockwiseDirection := (state.direction + 3) % 4
		newPosition = state.position.MoveDirection(counterClockwiseDirection)
		if !newPosition.OutOfBound(width, height) && input.Map[newPosition.Y][newPosition.X] != '#' {
			newPath := maps.Clone(state.path)
			newPath[newPosition] = true
			heap.Push(pq, &State{
				cost: cost + 1000,
				Coordinate: Coordinate{
					position:  coordinate.position,
					direction: counterClockwiseDirection,
				},
				path: newPath,
			})
		}
	}

	return maxScore, finalPath
}

func SolutionPartOne(input World) int {
	maxScore, _ := FindShortestPath(input)
	return maxScore
}

func SolutionPartTwo(input World) int {
	_, path := FindShortestPath(input)
	return len(path)
}
