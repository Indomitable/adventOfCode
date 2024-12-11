package helpers

import "fmt"

type Map struct {
	Cells [][]int
	Rows  int
	Cols  int
}

type Position struct {
	Row, Col int
}

type Cell struct {
	Value int
	Pos   Position
}

const (
	LEFT  = 1 << 0
	RIGHT = 1 << 1
	UP    = 1 << 2
	DOWN  = 1 << 3
)

type DIRECTION int

func (m *Map) GetRow(idx int) []int {
	return m.Cells[idx]
}

func (m *Map) GetColumn(idx int) (ret []int) {
	for _, row := range m.Cells {
		ret = append(ret, row[idx])
	}
	return
}

func (m *Map) GetCell(pos Position) Cell {
	return Cell{
		Value: m.Cells[pos.Row][pos.Col],
		Pos:   pos,
	}
}

func (m *Map) Iterate() func(func(cell Cell) bool) {
	return func(yield func(cell Cell) bool) {
		for rIdx, row := range m.Cells {
			for cIdx, value := range row {
				if !yield(Cell{value, Position{rIdx, cIdx}}) {
					return
				}
			}
		}
	}
}

func (m *Map) Print() {
	for _, row := range m.Cells {
		for _, col := range row {
			fmt.Printf("%d", col)
		}
		fmt.Println()
	}
}

func (p Position) WithRow(row int) Position {
	return Position{
		Row: row,
		Col: p.Col,
	}
}

func (p Position) WithCol(col int) Position {
	return Position{
		Row: p.Row,
		Col: col,
	}
}

func (a Position) CompareTo(b Position) int {
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
}
