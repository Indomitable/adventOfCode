package main

import (
	"bytes"
	"fmt"
	"helpers"
	"io"
	"os"
	"time"
)

func main() {
	start := time.Now()
	f, _ := os.Open("task.in")
	defer f.Close()
	content, _ := io.ReadAll(f)
	stones := bytes.Split(content, []byte{' '})
	numbers := helpers.MapSlice(stones, helpers.BytesToLong)

	m := group(numbers)
	for range 75 {
		m = processStep2(m) //slices.Collect(process(numbers))
	}
	fmt.Printf("Stones count: %d\n", countStones(m))
	fmt.Printf("Time: %v\n", time.Since(start))
}

func processStep1(stones []int64) func(func(int64) bool) {
	return func(yield func(x int64) bool) {
		for _, stone := range stones {
			if stone == 0 {
				yield(1)
				continue
			}
			cnt := helpers.LenDigits(stone)
			if cnt%2 == 0 {
				digits := helpers.ToDigits(stone)
				left := digits[:cnt/2]
				right := digits[cnt/2:]
				yield(helpers.ToNumber(left))
				yield(helpers.ToNumber(right))
				continue
			}
			yield(stone * 2024)
		}
	}
}

func processStep2(stones map[int64]int) map[int64]int {
	res := make(map[int64]int)
	for k, v := range stones {
		if k == 0 {
			res[1] += v
			continue
		}
		cnt := helpers.LenDigits(k)
		if cnt%2 == 0 {
			digits := helpers.ToDigits(k)
			left := helpers.ToNumber(digits[:cnt/2])
			right := helpers.ToNumber(digits[cnt/2:])
			res[left] += v
			res[right] += v
			continue
		}
		res[k*2024] += v
	}
	return res
}

func group(stones []int64) map[int64]int {
	res := make(map[int64]int)
	for _, s := range stones {
		res[s] += 1
	}
	return res
}

func countStones(stones map[int64]int) int64 {
	cnt := int64(0)
	for _, v := range stones {
		cnt += int64(v)
	}
	return cnt
}
