package main

import (
	"fmt"
	"helpers"
	"io"
	"os"
	"strings"
)

var emptySpace int = -1

func main() {
	f, _ := os.Open("./task.in")
	defer f.Close()
	content, _ := io.ReadAll(f)
	input := helpers.MapSlice(content, helpers.ByteToInt)
	expanded := expand(input)
	//compacted := compact(expanded)
	compactedBlocks := compactBlocks(expanded)
	//print(compactedBlocks)
	fmt.Println("Part 2: ", calculateHash(compactedBlocks))
}

func expand(input []int) []int {
	res := make([]int, 0)
	f := 0
	for i, v := range input {
		if i%2 == 0 {
			// file
			res = append(res, helpers.Repeat(f, v)...)
			f++
		} else {
			// space
			res = append(res, helpers.Repeat(emptySpace, v)...)
		}
	}
	return res
}

func print(v []int) {
	str := strings.Builder{}
	for _, v := range v {
		if v == emptySpace {
			str.WriteByte('.')
		} else {
			str.WriteByte(byte(v + 48))
		}
	}
	fmt.Printf("%s\n", str.String())
}

func compact(input []int) []int {
	k := 0
	for i := len(input) - 1; i > 0; i-- {
		if input[i] != emptySpace {
			for j := k; j < len(input); j++ {
				k = j
				if i == j {
					return input
				}
				if input[j] == emptySpace {
					input[j] = input[i]
					input[i] = emptySpace
					k++
					break
				}
			}
		}
	}
	return input
}

func compactBlocks(input []int) []int {
	copy := append([]int{}, input...)
	moved := make(map[int]bool)
	for fs, fe := range fileBlocksSeq(copy) {
		fileBlock := input[fs:fe]
		for es, ee := range emptySpaceSeq(input) {
			if es >= fe {
				break
			}
			emptyBlock := input[es:ee]
			if len(emptyBlock) >= len(fileBlock) {
				n := fileBlock[0]
				if moved[n] {
					panic("Should not move block")
				}
				moved[n] = true
				i := 0
				for f := fs; f < fe; f++ {
					input[es+i] = input[f]
					input[f] = emptySpace
					i++
				}
				break
			}
		}
	}
	return input
}

func fileBlocksSeq(input []int) func(func(b, e int) bool) {
	block := len(input) - 1
	return func(yield func(b, e int) bool) {
		for i := len(input) - 1; i > 0; i-- {
			if input[i] != input[block] {
				if input[block] != emptySpace {
					if !yield(i+1, block+1) {
						return
					}
				}
				block = i
			}
		}
	}
}

func emptySpaceSeq(input []int) func(func(b, e int) bool) {
	block := 0
	return func(yield func(b, e int) bool) {
		for i := 0; i < len(input); i++ {
			if input[i] != input[block] {
				if input[block] == emptySpace {
					if !yield(block, i) {
						return
					}
				}
				block = i
			}
		}
	}
}

func calculateHash(input []int) int64 {
	sum := int64(0)
	for i, v := range input {
		if v == emptySpace {
			continue
		}
		sum += int64(i * v)
	}
	return sum
}
