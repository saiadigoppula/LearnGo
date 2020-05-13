package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	str := "నేను అను సర్వనామము"

	bytes := []byte(str)

	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))
}
