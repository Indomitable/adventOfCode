package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"time"
)

func main() {
	start := time.Now()
	tasks := readTasks()
	var sum int64 = 0
	for _, task := range tasks {
		if verify(&task) {
			sum += task.Result
		}
	}
	fmt.Printf("%d\n", sum)
	fmt.Printf("Time elapsed: %v\n", time.Since(start))
}

type Task struct {
	Result  int64
	Numbers []int64
}

func readTasks() []Task {
	file, _ := os.Open("task.input")
	defer file.Close()
	content, _ := io.ReadAll(file)
	lines := bytes.Split(content, []byte{'\n'})
	tasks := make([]Task, len(lines))
	for i, line := range lines {
		parts := bytes.Split(line, []byte{':'})
		numbers := bytes.Split(bytes.TrimSpace(parts[1]), []byte{' '})
		tasks[i] = Task{
			Result:  BytesToLong(parts[0]),
			Numbers: mapSlice(numbers, BytesToLong),
		}
	}
	return tasks
}

func BytesToLong(buffer []byte) int64 {
	var result int64 = 0
	start := len(buffer) - 1
	for i := start; i >= 0; i-- {
		digit := int64(buffer[i] - 48)
		result += digit * int64(math.Pow10(start-i))
	}
	return result
}

func mapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	res := make([]R, len(slice))
	for i, item := range slice {
		res[i] = mapper(item)
	}
	return res
}

func GenerateCombinations(n, k int) func(func([]int) bool) {
	return func(yield func([]int) bool) {
		combination := make([]int, n)
		var generate func(int) bool
		generate = func(index int) bool {
			if index == n {
				if !yield(append([]int{}, combination...)) {
					return true // fast exit
				}
				return false
			}
			for i := 0; i < k; i++ {
				combination[index] = i
				if generate(index + 1) {
					return true
				}
			}
			return false
		}

		generate(0)
	}
}

func GenerateCombinations1(n, k int) [][]int {
	var result = make([][]int, int(math.Pow(float64(k), float64(n))))
	combination := make([]int, n)
	var generate func(int)
	var generation = 0

	generate = func(index int) {
		if index == n {
			result[generation] = append([]int{}, combination...) // copy into new array
			generation++
			return
		}
		for i := 0; i < k; i++ {
			combination[index] = i + 1
			generate(index + 1)
		}
	}

	generate(0)
	return result
}

func testCombination(numbers []int64, operations []int, expected int64) bool {
	current := numbers[0]
	for i, operation := range operations {
		switch operation {
		case 0: // multiply
			current *= numbers[i+1]
			break
		case 1: // sum
			current += numbers[i+1]
			break
		case 2: // concat
			current = ConcatNumbers(current, numbers[i+1])
			break
		}
		if current > expected {
			return false
		}
	}
	return current == expected
}

func ConcatNumbers(a, b int64) int64 {
	var m int64 = 1
	var sum int64 = 0
	for ld := range ExtractDigits(b) {
		sum += int64(ld) * m
		m *= 10
	}
	for rd := range ExtractDigits(a) {
		sum += int64(rd) * m
		m *= 10
	}
	return sum
}

func ExtractDigits(x int64) func(func(int64) bool) {
	var a int64 = 10
	var b int64 = 1
	return func(yield func(int64) bool) {
		for {
			digit := x % a / b
			if b > x {
				return
			}
			if !yield(digit) {
				return
			}
			a *= 10
			b *= 10
		}
	}
}

func verify(t *Task) bool {
	operations := GenerateCombinations(len(t.Numbers)-1, 3)
	for operation := range operations {
		if testCombination(t.Numbers, operation, t.Result) {
			return true
		}
	}
	return false
}

func ConvertNumber(n int, base int) int {
	result := 0
	multiplier := 1

	for n > 0 {
		remainder := n % base
		result += remainder * multiplier
		multiplier *= 10
		n /= base
	}
	return result
}

//func verifyParallel(t *Task) bool {
//	operations := GenerateCombinations(len(t.Numbers)-1, 3)
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	var wg sync.WaitGroup
//	results := make(chan bool)
//
//	for operation := range operations {
//		wg.Add(1)
//		go func(op []int) {
//			defer wg.Done()
//			if res := executeOperationsWithContext(t.Numbers, op, ctx) == t.Result; res {
//				select {
//				case results <- true:
//				default:
//				}
//			}
//		}(operation)
//	}
//	go func() {
//		wg.Wait()
//		close(results)
//	}()
//
//	result := <-results
//	cancel()
//	return result
//}
//
//func executeOperationsWithContext(numbers []int64, operations []int, ctx context.Context) int64 {
//	current := numbers[0]
//	for i, operation := range operations {
//		switch operation {
//		case 0: // multiply
//			current *= numbers[i+1]
//			break
//		case 1: // sum
//			current += numbers[i+1]
//			break
//		case 2: // concat
//			current = ConcatNumbers(current, numbers[i+1])
//			break
//		}
//	}
//	select {
//	case <-ctx.Done(): // cancelled
//		return -1
//	default:
//		return current
//	}
//}
