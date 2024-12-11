package main

import (
	"fmt"
	"testing"
)

func TestConvertion(t *testing.T) {
	v := []byte("123458799789797")

	if BytesToLong(v) != 123458799789797 {
		t.Errorf("want %d, got %d", 123458799789797, BytesToLong(v))
	}
}

//func TestGenerateOperations2(t *testing.T) {
//	comb := GenerateCombinations(4, 3)
//	for _, r := range comb {
//		fmt.Printf("%v\n", r)
//	}
//}

func TestGenerateOperations3(t *testing.T) {
	combinations := GenerateCombinations(4, 2)
	for r := range combinations {
		fmt.Printf("%v\n", r)
	}
	fmt.Printf("Exit")
}

func TestConcat(t *testing.T) {
	//x := ConcatNumbers(7505, 2)
	//if x != 75052 {
	//	t.Errorf("want %d, got %d", 75052, x)
	//}

	x := ConcatNumbers(7520424975, 2)
	if x != 75204249752 {
		t.Errorf("want %d, got %d", 75204249752, x)
	}
}

//func TestConvertNumber(t *testing.T) {
//	//x := ConvertNumber(7, 3)
//	//if x != 21 {
//	//	t.Errorf("want %d, got %d", 21, x)
//	//}
//	for x := range 100 {
//		fmt.Printf("%v\n", ConvertNumber(x, 3))
//	}
//
//}
