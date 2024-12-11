package helpers

import "math"

func BytesToLong(buffer []byte) int64 {
	var result int64 = 0
	start := len(buffer) - 1
	for i := start; i >= 0; i-- {
		digit := int64(buffer[i] - 48)
		result += digit * int64(math.Pow10(start-i))
	}
	return result
}

func ByteToInt(b byte) int {
	return int(b - 48)
}
