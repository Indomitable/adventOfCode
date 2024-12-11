package main

import (
	"fmt"
	"testing"
)

func TestGetPairs(t *testing.T) {
	pairs := GetPairs(4)
	for _, pair := range pairs {
		fmt.Printf("%v\n", pair)
	}
}

func TestMap(t *testing.T) {
	pairs := GetPairs(4)
	for _, pair := range pairs {
		fmt.Printf("%v\n", pair)
	}
}
