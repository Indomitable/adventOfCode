package helpers

import (
	"slices"
	"testing"
)

func TestToDigits(t *testing.T) {
	if !slices.Equal(ToDigits(1234), []int{1, 2, 3, 4}) {
		t.Error("Expected 1,2,3,4")
	}
}

func TestToDigitsZero(t *testing.T) {
	if !slices.Equal(ToDigits(0), []int{0}) {
		t.Error("Expected 0")
	}
}

func TestToDigits1(t *testing.T) {
	if !slices.Equal(ToDigits(10), []int{1, 0}) {
		t.Error("Expected 0")
	}
}

func TestToNumber0(t *testing.T) {
	if ToNumber([]int{1, 2, 3, 4}) != int64(1234) {
		t.Error("Expected 1234")
	}
}

func TestToNumber1(t *testing.T) {
	if ToNumber([]int{0, 0}) != int64(0) {
		t.Error("Expected 0")
	}
	if ToNumber([]int{0}) != int64(0) {
		t.Error("Expected 0")
	}
}

func TestToNumber2(t *testing.T) {
	if ToNumber([]int{1, 0}) != int64(10) {
		t.Error("Expected 10")
	}
}
