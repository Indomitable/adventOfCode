package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	leftList := make([]int, 0)
	rightList := make([]int, 0)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, "   ")
		leftNum, _ := strconv.Atoi(numbers[0])
		rightNum, _ := strconv.Atoi(numbers[1])
		leftList = placeItem(leftList, leftNum)
		rightList = placeItem(rightList, rightNum)
		lineNum++
	}

	handleLists1(leftList, rightList)
	handleLists2(leftList, rightList)
}

func placeItem(list []int, item int) []int {
	pos, _ := slices.BinarySearch(list, item)
	return slices.Insert(list, pos, item)
}

func handleLists1(leftList []int, rightList []int) {
	sum := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		sum += distance
	}
	fmt.Printf("Distance is: %d\n", sum)
}

func handleLists2(leftList []int, rightList []int) {
	rightMap := make(map[int]int)
	for i := 0; i < len(rightList); i++ {
		rightMap[rightList[i]]++
	}

	var similarity int64 = 0
	for i := 0; i < len(leftList); i++ {
		similarity += int64(leftList[i] * rightMap[leftList[i]])
	}
	fmt.Printf("Similarity is: %d\n", similarity)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
