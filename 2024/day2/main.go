package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	start := time.Now()
	file, _ := os.Open("task.input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var count = 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Split(line, " ")
		list := createColumn(numbers)
		if checkColumn(list) || checkBrute(list) {
			count++
		}
	}
	fmt.Printf("Number of columns: %d\n", count)
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
}

func createColumn(numbers []string) []int {
	list := make([]int, len(numbers))
	for i, number := range numbers {
		x, er := strconv.Atoi(number)
		if er != nil {
			fmt.Printf("Unable to convert string to int: %s\n", number)
		} else {
			list[i] = x
		}
	}
	return list
}

func checkProgress(prev int, current int, direction int) bool {
	diff := current - prev
	if direction > 0 {
		return diff >= 1 && diff <= 3
	} else {
		return diff >= -3 && diff <= -1
	}

}

func checkColumn(numbers []int) bool {
	var prev = 0
	var direction = 0
	for i, number := range numbers {
		if i == 0 {
			prev = number
			continue
		}
		if i == 1 {
			direction = number - prev
		}
		good := checkProgress(prev, number, direction)
		if good {
			prev = number
		} else {
			return false
		}
	}
	return true
}

//func checkTolerance(numbers []int) bool {
//	var prev = 0
//	var direction = -1
//	var tries = 0
//	for i := 0; i < len(numbers); {
//		current := numbers[i]
//		if i == 0 {
//			prev = current
//			i++
//			continue
//		}
//		good := checkProgress(prev, current, direction)
//		if good {
//			prev = current
//			i++
//			continue
//		} else {
//			if i == len(numbers)-1 {
//				// last one is bad we can remove it
//				return tries == 0
//			} else {
//				if tries == 1 {
//					return false
//				}
//				good = checkProgress(prev, numbers[i+1], direction)
//				if good {
//					prev = numbers[i+1]
//					i = i + 2
//					tries++
//				} else {
//					return false
//				}
//			}
//		}
//	}
//	return true
//}

func checkBrute(numbers []int) bool {
	for i := 0; i < len(numbers); i++ {
		list := make([]int, 0, len(numbers)-1)
		list = append(list, numbers[:i]...)
		list = append(list, numbers[i+1:]...)
		if checkColumn(list) {
			return true
		}
	}
	return false
}
