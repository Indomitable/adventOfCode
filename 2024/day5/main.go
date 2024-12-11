package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./task.input")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rules = make(map[int][]int)
	var printQueue [][]int
	var loadRules = true
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			loadRules = false
			continue
		}
		if loadRules {
			parseRules(line, rules)
		} else {
			printQueue = append(printQueue, parsePrintQueue(line))
		}
	}

	var sum int64 = 0
	var sumFixed int64 = 0
	for _, queue := range printQueue {
		if verifyLine(rules, queue) {
			sum += int64(getMiddleNumber(queue))
		} else {
			fixQueue(rules, queue)
			sumFixed += int64(getMiddleNumber(queue))
		}
	}
	fmt.Printf("Count is: %d\n", sum)
	fmt.Printf("Count is: %d\n", sumFixed)
}

func parseRules(rule string, order map[int][]int) {
	var rules = strings.Split(rule, "|")
	f, _ := strconv.Atoi(rules[0])
	n, _ := strconv.Atoi(rules[1])
	order[f] = append(order[f], n)
}

func parsePrintQueue(line string) []int {
	return mapSlice(strings.Split(line, ","), func(s string) int {
		parse, _ := strconv.Atoi(s)
		return parse
	})
}

func mapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	res := make([]R, len(slice))
	for i, item := range slice {
		res[i] = mapper(item)
	}
	return res
}

func verifyLine(rules map[int][]int, queue []int) bool {
	processed := make(map[int]bool)
	for _, n := range queue {
		if !checkViolation(n, rules, processed) {
			processed[n] = true
		} else {
			return false
		}
	}
	return true
}

func checkViolation(n int, rules map[int][]int, processed map[int]bool) bool {
	var rule = rules[n]
	if rule == nil {
		// no rules for this number, no violations
		return false
	}
	for _, prev := range rule {
		// check if any numbers that should be after our number is processed
		if processed[prev] {
			return true
		}
	}
	return false
}

func getMiddleNumber(queue []int) int {
	return queue[len(queue)/2]
}

func fixQueue(rules map[int][]int, queue []int) []int {
	slices.SortFunc(queue, func(a, b int) int {
		return compare(a, b, rules)
	})
	return queue
}

func compare(a int, b int, rules map[int][]int) int {
	aRule := rules[a]
	if aRule != nil {
		if slices.Contains(aRule, b) {
			// a should be before b so
			return -1
		}
	}
	bRule := rules[b]
	if bRule != nil {
		if slices.Contains(bRule, a) {
			return 1
		}
	}
	return 0 // no rule or for these numbers
}
