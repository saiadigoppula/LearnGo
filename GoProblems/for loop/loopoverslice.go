package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields(
		"lazy cat is near my bed")
	for i := 0; i < len(words); i++ {
		fmt.Printf("#%-2d: %q\n", i+1, words[i])
	}
}
