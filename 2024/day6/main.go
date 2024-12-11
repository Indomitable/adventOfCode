package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("./task.input")
	defer file.Close()
	c, _ := io.ReadAll(file)

	var m = createMap(c)
	visited, _ := startRoute(m, -1)
	fmt.Printf("Positions visited: %d\n", len(visited))

	r := bruteForceParallel(m, visited)
	fmt.Printf("Positions obstacle to be placed: %d\n", r)
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
}

const (
	Empty     = '.'
	Obstacle  = '#'
	GuardType = '^'
)

const (
	LEFT  = 1 << 0
	RIGHT = 1 << 1
	UP    = 1 << 2
	DOWN  = 1 << 3
)

type Map struct {
	Fields    []byte //MapItem
	Length    int
	RowLength int
	Guard     Guard
}

type Orientation int

type Guard struct {
	Pos         int
	Orientation Orientation
}

type VisitedLocations map[int]Orientation

func createMap(content []byte) *Map {
	rowLength := bytes.IndexByte(content, '\n')
	rows := bytes.ReplaceAll(content, []byte{'\n'}, []byte{})
	return &Map{
		Fields:    rows,
		Length:    len(rows),
		RowLength: rowLength,
		Guard: Guard{
			Pos:         bytes.IndexByte(rows, GuardType),
			Orientation: UP,
		},
	}
}

func (m *Map) MoveGuard(oldState Guard, direction Orientation, extra int) (newState Guard, finish bool) {
	isAllowed := func(item byte, pos int) bool {
		return item != Obstacle && pos != extra
	}
	switch direction {
	case LEFT:
		if oldState.Pos%m.RowLength == 0 {
			// game over
			return oldState, true
		}
		nextPos := oldState.Pos - 1
		item := m.Fields[nextPos]
		if isAllowed(item, nextPos) {
			return Guard{Pos: nextPos, Orientation: LEFT}, false
		} else {
			// step right
			return m.MoveGuard(oldState, UP, extra)
		}
	case RIGHT:
		if (oldState.Pos+1)%m.RowLength == 0 {
			// game over
			return oldState, true
		}
		nextPos := oldState.Pos + 1
		item := m.Fields[nextPos]
		if isAllowed(item, nextPos) {
			return Guard{Pos: nextPos, Orientation: RIGHT}, false
		} else {
			// step right
			return m.MoveGuard(oldState, DOWN, extra)
		}
	case UP:
		if oldState.Pos < m.RowLength {
			return oldState, true
		}
		nextPos := oldState.Pos - m.RowLength
		item := m.Fields[nextPos]
		if isAllowed(item, nextPos) {
			return Guard{Pos: nextPos, Orientation: UP}, false
		} else {
			return m.MoveGuard(oldState, RIGHT, extra)
		}
	case DOWN:
		if oldState.Pos > m.Length-m.RowLength {
			return oldState, true
		}
		nextPos := oldState.Pos + m.RowLength
		item := m.Fields[nextPos]
		if isAllowed(item, nextPos) {
			return Guard{Pos: nextPos, Orientation: DOWN}, false
		} else {
			return m.MoveGuard(oldState, LEFT, extra)
		}
	}
	panic("Illegal move")
}

func startRoute(m *Map, extra int) (VisitedLocations, bool) {
	path := make(VisitedLocations)
	current := m.Guard
	path[current.Pos] = current.Orientation
	for {
		guard, finish := m.MoveGuard(current, current.Orientation, extra)
		if finish {
			return path, false
		}
		if inLoop(guard, path) {
			// guard is in a loop return path
			return path, true
		}
		current = guard
		path[current.Pos] |= current.Orientation
	}
}

func inLoop(guard Guard, visited VisitedLocations) bool {
	return visited[guard.Pos]&guard.Orientation == guard.Orientation
}

func bruteForce(m *Map, visited VisitedLocations) int {
	sum := 0
	for place := range visited {
		cell := m.Fields[place]
		if cell == Empty {
			_, loop := startRoute(m, place)
			if loop {
				sum++
			}
		}
	}
	return sum
}

func bruteForceParallel(m *Map, visited VisitedLocations) int {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var wg sync.WaitGroup
	results := make(chan int)
	for place := range visited {
		cell := m.Fields[place]
		if cell == Empty {
			wg.Add(1)
			go func(p int) {
				defer wg.Done()
				_, loop := startRoute(m, p)
				if loop {
					results <- 1
				}
			}(place)
		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	//return len(results)
	sum := 0
	for result := range results {
		sum += result
	}
	return sum
}
