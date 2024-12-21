package main

import (
	"bytes"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type Point [2]int

type State interface {
	Cost() int
}

type PriorityQueue[T State] []*T

func (pq PriorityQueue[T]) Len() int { return len(pq) }
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return (*pq[i]).Cost() < (*pq[j]).Cost()
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue[T]) Push(x interface{}) {
	item := x.(*T)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Maze struct {
	Start Point
	End   Point
	Cells [][]rune
}

func ParseInput(filepath string) Maze {
	maze := Maze{
		Cells: make([][]rune, 0),
	}

	raw, _ := os.ReadFile(filepath)

	for y, line := range bytes.Split(raw, []byte("\n")) {
		cells := make([]rune, 0)

		for x, cell := range line {
			cells = append(cells, rune(cell))
			if cell == 'S' {
				maze.Start = Point{x, y}
			}
			if cell == 'E' {
				maze.End = Point{x, y}
			}
		}

		maze.Cells = append(maze.Cells, cells)
	}

	return maze
}

type MazeState struct {
	Point Point
	Paths []Point
}

func (m MazeState) Cost() int {
	return len(m.Paths)
}

func FindShortestPath(maze Maze, start Point, end Point) (*MazeState, map[Point]int) {
	q := &PriorityQueue[MazeState]{
		&MazeState{
			Point: start,
			Paths: []Point{start},
		},
	}

	dist := map[Point]int{}

	for q.Len() > 0 {
		state := heap.Pop(q).(*MazeState)

		if state.Point == end {
			return state, dist
		}

		if score, ok := dist[state.Point]; ok && score < state.Cost() {
			continue
		}

		dist[state.Point] = state.Cost()

		for _, direction := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newPoint := Point{state.Point[0] + direction[0], state.Point[1] + direction[1]}
			if newPoint[0] < 0 || newPoint[0] >= len(maze.Cells[0]) ||
				newPoint[1] < 0 || newPoint[1] >= len(maze.Cells) {
				continue
			}
			char := maze.Cells[newPoint[1]][newPoint[0]]
			if char == '#' {
				continue
			}

			newPaths := make([]Point, len(state.Paths))
			copy(newPaths, state.Paths)
			newPaths = append(newPaths, newPoint)

			heap.Push(q, &MazeState{
				Point: newPoint,
				Paths: newPaths,
			})
		}
	}

	return nil, dist
}

func FindShortestPathWithJump(
	maze Maze,
	start Point, end Point,
	preCompute map[Point]int,
	optimal int, threshold int,
) int {
	q := &PriorityQueue[MazeState]{
		&MazeState{
			Point: start,
			Paths: []Point{start},
		},
	}

	dist := map[Point]int{}
	jumpTable := map[int]map[[2]Point]bool{}

	countPossible := func() int {
		count := 0
		for _, points := range jumpTable {
			// fmt.Printf("Save: %d | %d\n", table, len(points))
			count += len(points)
		}

		return count
	}

	for q.Len() > 0 {
		state := heap.Pop(q).(*MazeState)

		if state.Point == end {
			return countPossible()
		}

		if score, ok := dist[state.Point]; ok && score < state.Cost() {
			continue
		}

		dist[state.Point] = state.Cost()

		for _, direction := range []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
			newPoint := Point{state.Point[0] + direction[0], state.Point[1] + direction[1]}
			jumpPoint := Point{newPoint[0] + direction[0], newPoint[1] + direction[1]}

			isWall := newPoint[0] < 0 || newPoint[0] >= len(maze.Cells[0]) ||
				newPoint[1] < 0 || newPoint[1] >= len(maze.Cells) ||
				maze.Cells[newPoint[1]][newPoint[0]] == '#'

			if isWall {
				if preComputeState, ok := preCompute[jumpPoint]; ok {
					estimated := state.Cost() + preComputeState + 1
					distance := optimal - estimated
					if estimated <= optimal && distance >= threshold {
						oldTable, ok := jumpTable[distance]
						if !ok {
							oldTable = map[[2]Point]bool{}
						}
						oldTable[[2]Point{state.Point, newPoint}] = true
						jumpTable[distance] = oldTable
					}
				}
				continue
			}

			newPaths := make([]Point, len(state.Paths))
			copy(newPaths, state.Paths)
			newPaths = append(newPaths, newPoint)

			heap.Push(q, &MazeState{
				Point: newPoint,
				Paths: newPaths,
			})
		}
	}

	return countPossible()
}

func SolutionPartOne(maze Maze, threshold int) int {
	mazeState, dist := FindShortestPath(maze, maze.End, maze.Start)

	count := FindShortestPathWithJump(
		maze,
		maze.Start, maze.End,
		dist,
		len(mazeState.Paths), threshold,
	)

	return count
}

func FindShortestPathWithJumpCost(
	maze Maze,
	bestState MazeState,
	preCompute map[Point]int,
	threshold int, maxSteps int,
) int {

	// saved := map[int]int{}
	pairs := map[[2]Point]int{}

	for _, p1 := range bestState.Paths {
		for _, p2 := range bestState.Paths {
			if p1 == p2 {
				continue
			}

			p1Cost := preCompute[p1]
			p2Cost := preCompute[p2]
			dx, dy := math.Abs(float64(p1[0]-p2[0])), math.Abs(float64(p1[1]-p2[1]))
			save := math.Abs(float64(p2Cost-p1Cost)) - (dx + dy)

			if dx+dy <= float64(maxSteps) && save >= math.Max(float64(threshold), 2) {
				if save == 79 {
					fmt.Printf("P1: %v | P2: %v | Save: %f\n", p1, p2, save)
				}
				pair := [2]Point{p1, p2}
				communitivePair := [2]Point{p2, p1}
				if _, ok := pairs[communitivePair]; ok {
					pair = communitivePair
				}
				pairs[pair] = int(save)
			}
		}
	}

	saved := map[int]int{}

	for _, cost := range pairs {
		saved[cost]++
	}

	result := 0
	for _, cost := range saved {
		result += cost
		// fmt.Printf("Save %d with %d times\n", save, cost)
	}

	return result
}

func SolutionPartTwo(maze Maze, threshold int) int {
	mazeState, _ := FindShortestPath(maze, maze.Start, maze.End)
	dist := map[Point]int{}
	for index, path := range mazeState.Paths {
		dist[path] = len(mazeState.Paths) - index
	}

	return FindShortestPathWithJumpCost(
		maze,
		*mazeState,
		dist, threshold, 20,
	)
}
