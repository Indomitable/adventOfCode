package main

import (
	"io"
	"os"
)

func main() {
	f, _ := os.Open("test.in")
	defer f.Close()
	content, _ := io.ReadAll(f)

}

