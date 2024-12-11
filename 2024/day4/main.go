package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	file, _ := os.Open("./task.input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var positions [][]Position
	for row, line := range lines {
		for col, char := range line {
			if char == 'M' {
				pos := Position{Row: row, Col: col}
				//count += SearchHorizontal(lines, pos)
				//count += SearchVertical(lines, pos)
				//count += SearchDiagonal(lines, pos)
				positions = append(positions, SearchXMax(lines, pos)...)
			}
		}
	}
	var founds = DistinctPosition(positions)
	for _, found := range founds {
		for _, pos := range found {
			print(pos.String() + ";")
		}
		println()
	}
	fmt.Printf("Count of X-MAS: %d\n", len(founds))
}

type Position struct {
	Row    int
	Col    int
	Letter byte
}

func (pos Position) String() string {
	return fmt.Sprintf("%d,%d,%s", pos.Row, pos.Col, string(pos.Letter))
}

func SearchHorizontal(lines []string, pos Position) int {
	var count = 0
	line := lines[pos.Row]
	if pos.Col+3 < len(line) {
		var wordForward = line[pos.Col : pos.Col+4]
		if wordForward == "XMAS" {
			count++
		}
	}
	if pos.Col-3 >= 0 {
		var wordBackward = line[pos.Col-3 : pos.Col+1]
		if wordBackward == "SAMX" {
			count++
		}
	}
	return count
}

func SearchVertical(lines []string, pos Position) int {
	var count = 0
	if pos.Row+3 < len(lines) {
		char0 := lines[pos.Row][pos.Col]
		char1 := lines[pos.Row+1][pos.Col]
		char2 := lines[pos.Row+2][pos.Col]
		char3 := lines[pos.Row+3][pos.Col]
		var wordDownward = string([]uint8{char0, char1, char2, char3})
		if wordDownward == "XMAS" {
			count++
		}
	}
	if pos.Row-3 >= 0 {
		char0 := lines[pos.Row][pos.Col]
		char1 := lines[pos.Row-1][pos.Col]
		char2 := lines[pos.Row-2][pos.Col]
		char3 := lines[pos.Row-3][pos.Col]
		var wordUpward = string([]uint8{char0, char1, char2, char3})
		if wordUpward == "XMAS" {
			count++
		}
	}
	return count
}

func SearchDiagonal(lines []string, pos Position) int {
	var count = 0
	if pos.Row+3 < len(lines) {
		// try search down
		var line0 = lines[pos.Row]
		if pos.Col+3 < len(line0) {
			// search down right
			char0 := lines[pos.Row][pos.Col]
			char1 := lines[pos.Row+1][pos.Col+1]
			char2 := lines[pos.Row+2][pos.Col+2]
			char3 := lines[pos.Row+3][pos.Col+3]
			var word = string([]uint8{char0, char1, char2, char3})
			if word == "XMAS" {
				count++
			}
		}
		if pos.Col-3 >= 0 {
			// search down left
			char0 := lines[pos.Row][pos.Col]
			char1 := lines[pos.Row+1][pos.Col-1]
			char2 := lines[pos.Row+2][pos.Col-2]
			char3 := lines[pos.Row+3][pos.Col-3]
			var word = string([]uint8{char0, char1, char2, char3})
			if word == "XMAS" {
				count++
			}
		}
	}
	if pos.Row-3 >= 0 {
		// search up
		var line0 = lines[pos.Row]
		if pos.Col+3 < len(line0) {
			// search up right
			char0 := lines[pos.Row][pos.Col]
			char1 := lines[pos.Row-1][pos.Col+1]
			char2 := lines[pos.Row-2][pos.Col+2]
			char3 := lines[pos.Row-3][pos.Col+3]
			var word = string([]uint8{char0, char1, char2, char3})
			if word == "XMAS" {
				count++
			}
		}
		if pos.Col-3 >= 0 {
			// search up left
			char0 := lines[pos.Row][pos.Col]
			char1 := lines[pos.Row-1][pos.Col-1]
			char2 := lines[pos.Row-2][pos.Col-2]
			char3 := lines[pos.Row-3][pos.Col-3]
			var word = string([]uint8{char0, char1, char2, char3})
			if word == "XMAS" {
				count++
			}
		}
	}
	return count
}

func SearchXMax(lines []string, pos Position) [][]Position {
	var positions = make([][]Position, 0, 4)
	if pos.Row+2 < len(lines) {
		// try search down
		var line0 = lines[pos.Row]
		if pos.Col+2 < len(line0) {
			// search down right
			m := lines[pos.Row][pos.Col]
			a := lines[pos.Row+1][pos.Col+1]
			s := lines[pos.Row+2][pos.Col+2]
			ms := lines[pos.Row+2][pos.Col]
			sm := lines[pos.Row][pos.Col+2]
			if m == 'M' && a == 'A' && s == 'S' && ((ms == 'M' && sm == 'S') || (ms == 'S' && sm == 'M')) {
				positions = append(positions, []Position{
					{pos.Row, pos.Col, m},
					{pos.Row + 1, pos.Col + 1, a},
					{pos.Row + 2, pos.Col + 2, s},
					{pos.Row + 2, pos.Col, ms},
					{pos.Row, pos.Col + 2, sm}})
			}
		}
		if pos.Col-2 >= 0 {
			// search down left
			m := lines[pos.Row][pos.Col]
			a := lines[pos.Row+1][pos.Col-1]
			s := lines[pos.Row+2][pos.Col-2]
			ms := lines[pos.Row+2][pos.Col]
			sm := lines[pos.Row][pos.Col-2]
			if m == 'M' && a == 'A' && s == 'S' && ((ms == 'M' && sm == 'S') || (ms == 'S' && sm == 'M')) {
				positions = append(positions, []Position{
					{pos.Row, pos.Col, m},
					{pos.Row + 1, pos.Col - 1, a},
					{pos.Row + 2, pos.Col - 2, s},
					{pos.Row + 2, pos.Col, ms},
					{pos.Row, pos.Col - 2, sm}})
			}
		}
	}
	if pos.Row-2 >= 0 {
		// search up
		var line0 = lines[pos.Row]
		if pos.Col+2 < len(line0) {
			// search up right
			m := lines[pos.Row][pos.Col]
			a := lines[pos.Row-1][pos.Col+1]
			s := lines[pos.Row-2][pos.Col+2]
			ms := lines[pos.Row-2][pos.Col]
			sm := lines[pos.Row][pos.Col+2]
			if m == 'M' && a == 'A' && s == 'S' && ((ms == 'M' && sm == 'S') || (ms == 'S' && sm == 'M')) {
				positions = append(positions, []Position{
					{pos.Row, pos.Col, m},
					{pos.Row - 1, pos.Col + 1, a},
					{pos.Row - 2, pos.Col + 2, s},
					{pos.Row - 2, pos.Col, ms},
					{pos.Row, pos.Col + 2, sm}})
			}
		}
		if pos.Col-2 >= 0 {
			// search up left
			m := lines[pos.Row][pos.Col]
			a := lines[pos.Row-1][pos.Col-1]
			s := lines[pos.Row-2][pos.Col-2]
			ms := lines[pos.Row-2][pos.Col]
			sm := lines[pos.Row][pos.Col-2]
			if m == 'M' && a == 'A' && s == 'S' && ((ms == 'M' && sm == 'S') || (ms == 'S' && sm == 'M')) {
				positions = append(positions, []Position{
					{pos.Row, pos.Col, m},
					{pos.Row - 1, pos.Col - 1, a},
					{pos.Row - 2, pos.Col - 2, s},
					{pos.Row - 2, pos.Col, ms},
					{pos.Row, pos.Col - 2, sm}})
			}
		}
	}
	return positions
}

func DistinctPosition(positions [][]Position) [][]Position {
	slices.SortFunc(positions, func(a, b []Position) int {
		a0 := a[1]
		a1 := b[1]
		if a0.Row > a1.Row {
			return 1
		}
		if a0.Row < a1.Row {
			return -1
		}
		if a0.Row == a1.Row && a0.Col > a1.Col {
			return 1
		}
		if a0.Row == a1.Row && a0.Col < a1.Col {
			return -1
		}
		return 0
	})
	return slices.CompactFunc(positions, func(p0 []Position, p1 []Position) bool {
		return p0[1] == p1[1]
	})
}

//.M.S......
//..A..MSMS.
//.M.S.MAA..
//..A.ASMSM.
//.M.S.M....
//..........
//S.S.S.S.S.
//.A.A.A.A..
//M.M.M.M.M.
//..........

//MMMSXXMASM
//MSAMXMSMSA
//AMXSXMAAMM
//MSAMASMSMX
//XMASAMXAMM
//XXAMMXXAMA
//SMSMSASXSS
//SAXAMASAAA
//MAMMMXMMMM
//MXMXAXMASX
