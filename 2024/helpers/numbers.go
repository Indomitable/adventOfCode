package helpers

import "math"

func ToDigits(n int64) []int {
	if n == 0 {
		return []int{0}
	}
	if n < 0 {
		n = -n
	}
	l := LenDigits(n)
	res := make([]int, l)
	c := l - 1
	for i := n; i > 0; {
		d := int(i % 10)
		res[c] = d
		i = i / 10
		c--
	}
	return res
}

func ToNumber(digits []int) int64 {
	if len(digits) == 0 {
		return 0
	}
	n := int64(0)
	j := 0
	for i := len(digits) - 1; i >= 0; i-- {
		n += int64(digits[i]) * int64(math.Pow10(j))
		j++
	}
	return n
}

func LenDigits(n int64) (res int) {
	return int(math.Floor(math.Log10(float64(n)))) + 1
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
