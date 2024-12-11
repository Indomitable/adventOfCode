package main

import (
	"bytes"
	"fmt"
	"helpers"
	"io"
	"os"
	"slices"
)

func main() {
	f, _ := os.Open("task.in")
	defer f.Close()
	content, _ := io.ReadAll(f)
	lines := bytes.Split(content, []byte{'\n'})
	m := createMap(lines)
	routes := findRoutes(m)
	// routes = distinctRoutes(routes)
	for _, route := range routes {
		fmt.Printf("%v\n", route)
	}
	fmt.Printf("Routes: %d\n", len(routes))
	fmt.Printf("Score: %d\n", score(routes))
}

func createMap(lines [][]byte) *helpers.Map {
	rows := len(lines)
	cols := len(lines[0])
	fields := make([][]int, rows)
	for rowIdx, rowIn := range lines {
		row := make([]int, cols)
		for colIdx, colIn := range rowIn {
			row[colIdx] = helpers.ByteToInt(colIn)
		}
		fields[rowIdx] = row
	}
	return &helpers.Map{
		Cells: fields,
		Rows:  rows,
		Cols:  cols,
	}
}

func findRoutes(m *helpers.Map) (routes [][]helpers.Cell) {
	for cell := range m.Iterate() {
		if cell.Value == 0 {
			routes = append(routes, slices.Collect(walk(m, cell, []helpers.Cell{cell}))...)
		}
	}
	return
}

func walk(m *helpers.Map, cell helpers.Cell, previousCells []helpers.Cell) func(func([]helpers.Cell) bool) {
	return func(yield func([]helpers.Cell) bool) {
		for adjacentCell := range ititerateAdjacentCells(m, cell.Pos) {
			p := helpers.Copy(previousCells)
			if adjacentCell.Value == 9 && len(p) == 9 {
				if !yield(append(p, adjacentCell)) {
					return
				}
				continue
			}
			if adjacentCell.Value == cell.Value+1 {
				c := append(p, adjacentCell)
				for route := range walk(m, adjacentCell, c) {
					yield(route)
				}
			}
		}
	}
}

// func walk1(m *helpers.Map, cell helpers.Cell, previousCells []helpers.Cell) [][]helpers.Cell {
// 	var routes [][]helpers.Cell
// 	for adjacentCell := range ititerateAdjacentCells(m, cell.Pos) {
// 		if adjacentCell.Value == 9 && len(previousCells) == 9 {
// 			routes = append(routes, append(previousCells, adjacentCell))
// 		}
// 		if adjacentCell.Value == cell.Value+1 {
// 			c := append(previousCells, adjacentCell)
// 			routes = append(routes, walk1(m, adjacentCell, c)...)
// 		}
// 	}
// 	return routes
// }

func distinctRoutes(routes [][]helpers.Cell) (distinct [][]helpers.Cell) {
	slices.SortFunc(routes, routesCompare)
	return slices.CompactFunc(routes, routesEqual)
}

func routesCompare(a, b []helpers.Cell) int {
	for i := 0; i < len(a); i++ {
		cellA := a[i]
		cellB := b[i]
		c := cellA.Pos.CompareTo(cellB.Pos)
		if c != 0 {
			return c
		}
	}
	return 0
}

func routesEqual(a, b []helpers.Cell) bool {
	return routesCompare(a, b) == 0
}

func ititerateAdjacentCells(m *helpers.Map, pos helpers.Position) func(func(cell helpers.Cell) bool) {
	cell := m.GetCell(pos)
	return func(yield func(cell helpers.Cell) bool) {
		if cell.Pos.Row > 0 {
			upCell := m.GetCell(pos.WithRow(cell.Pos.Row - 1))
			yield(upCell)
		}
		if cell.Pos.Col < m.Cols-1 {
			rightCell := m.GetCell(pos.WithCol(cell.Pos.Col + 1))
			yield(rightCell)
		}
		if cell.Pos.Row < m.Rows-1 {
			downCell := m.GetCell(pos.WithRow(cell.Pos.Row + 1))
			yield(downCell)
		}
		if cell.Pos.Col > 0 {
			leftCell := m.GetCell(pos.WithCol(cell.Pos.Col - 1))
			yield(leftCell)
		}
	}
}

func score(routes [][]helpers.Cell) int {
	res := 0
	m := make(map[helpers.Cell]*helpers.Set[helpers.Cell])
	for _, route := range routes {
		cell0 := route[0]
		cell9 := route[9]
		set, found := m[cell0]
		if found {
			set.Add(cell9)
		} else {
			m[cell0] = helpers.NewSet[helpers.Cell](cell9)
		}
	}
	for _, z := range m {
		res += z.Size()
	}
	return res
}
