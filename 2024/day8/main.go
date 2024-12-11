package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"slices"
	"time"
)

func main() {
	start := time.Now()
	file, _ := os.Open("./task.in")
	defer file.Close()

	content, _ := io.ReadAll(file)
	m := readMap(content)
	var mirrors = make(map[int]bool)
	//var mirrors []Pos
	for _, fields := range m.Fields {
		for m := range GetMirrorPoints(fields, m.Rows, m.Cols) {
			mirrors[m.Hashcode()] = true
		}
		//mirrors = slices.AppendSeq(mirrors, GetMirrorPoints(fields, m.Rows, m.Cols))
	}
	///1045
	//printMap(m, mirrors)
	fmt.Printf("Mirrors: %d\n", len(mirrors))
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
}

func printMap(m *Map, mirrors []Pos) {
	p := ""
	for r := 0; r < m.Rows; r++ {
		for c := 0; c < m.Cols; c++ {
			mirror := slices.IndexFunc(mirrors, func(p Pos) bool {
				return p.Row == r && p.Col == c
			})
			if mirror > -1 {
				p += "#"
			} else {
				p += "."
			}
		}
		p += "\n"
	}
	fmt.Println(p)
}

func distinctMirrors(m []Pos) []Pos {
	slices.SortFunc(m, func(a, b Pos) int {
		if a.Row > b.Row {
			return 1
		}
		if a.Row < b.Row {
			return -1
		}
		if a.Row == b.Row && a.Col > b.Col {
			return 1
		}
		if a.Row == b.Row && a.Col < b.Col {
			return -1
		}
		return 0
	})
	return slices.CompactFunc(m, func(p0 Pos, p1 Pos) bool {
		return p0 == p1
	})
}

func readMap(content []byte) *Map {
	rows := bytes.Split(content, []byte{'\n'})
	fields := make(map[byte][]Pos)
	for irow, row := range rows {
		for icell, cell := range row {
			if cell != '.' {
				fields[cell] = append(fields[cell], Pos{irow, icell})
			}
		}
	}
	return &Map{
		Fields: fields,
		Rows:   len(rows),
		Cols:   len(rows[0]),
	}
}

func abs[T int | int64](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func compare(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func MirrorPoint2(a, b Pos, m int) Pos {
	diffRow := abs(a.Row - b.Row)
	diffCol := abs(a.Col - b.Col)
	return Pos{
		b.Row + (compare(a.Row, b.Row) * m * diffRow),
		b.Col + (compare(a.Col, b.Col) * m * diffCol),
	}
}

func isValid(pos Pos, rows, cols int) bool {
	return !(pos.Row < 0 || pos.Row >= rows || pos.Col < 0 || pos.Col >= cols)
}

func GetPairs(n int) [][]int {
	pairs := make([][]int, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pairs = append(pairs, []int{i, j})
		}
	}
	return pairs
}

func GetMirrorPoints(positions []Pos, rows, cols int) func(func(Pos) bool) {
	pairs := GetPairs(len(positions))
	return func(yield func(Pos) bool) {
		for _, pair := range pairs {
			a := positions[pair[0]]
			b := positions[pair[1]]
			IterMirrorPoints(yield, a, b, rows, cols, true)
			IterMirrorPoints(yield, a, b, rows, cols, false)
		}
	}
}

func IterMirrorPoints(yield func(Pos) bool, a, b Pos, rows, cols int, dir bool) {
	var i = 0
	for {
		mirror := MirrorPoint2(a, b, i)
		if !isValid(mirror, rows, cols) {
			break
		}
		if !yield(mirror) {
			return
		}
		if dir {
			i++
		} else {
			i--
		}
	}
}

type Pos struct {
	Row, Col int
}

type Map struct {
	Fields map[byte][]Pos
	Rows   int
	Cols   int
}

func (pos Pos) Hashcode() int {
	return pos.Row*1019 ^ pos.Col
}
