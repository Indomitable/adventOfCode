package main

import (
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	content, _ := io.ReadAll(file)
	doFunc := "do()"
	dontFunc := "don't()"

	text := string(content)
	var sum int64 = 0
	var doSum = true
	for len(text) > 0 {
		doIndex := strings.Index(text, doFunc)
		dontIndex := strings.Index(text, dontFunc)
		funcIndex := getMinBiggerThanMinusOne(doIndex, dontIndex)
		if doSum {
			if funcIndex == -1 {
				sum += calc(text)
				break
			} else {
				sum += calc(text[:funcIndex])
			}
		}
		text = text[funcIndex+1:] // cut 'd' letter so not to find it next time.
		doSum = funcIndex == doIndex
	}
	println(sum)
}

func getMinBiggerThanMinusOne(first int, second int) int {
	if first == -1 {
		return second
	}
	if second == -1 {
		return first
	}
	if first < second {
		return first
	}
	return second
}

func calc(text string) int64 {
	var sum int64 = 0
	reg := regexp.MustCompile(`mul\((?P<X>\d{1,3}),(?P<Y>\d{1,3})\)`)
	matches := reg.FindAllStringSubmatch(string(text), -1)
	for _, match := range matches {
		println(match[0])
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += int64(x * y)
	}
	return sum
}
