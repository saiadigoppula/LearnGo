package main

import (
	"fmt"
)

func main() {

	fmt.Println(freqAlphabets("10#11#12"))
}
func freqAlphabets(s string) string {
	x := []rune(s)
	str := ""

	for i := len(x) - 1; i >= 0; i-- {
		if x[i] == '#' {
			str = string(96+((x[i-1]-48)+(x[i-2]-48)*10)) + str
			i = i - 2
		} else {
			str = string(96+(x[i]-48)) + str
		}
	}

	fmt.Println(int(x[0]))
	fmt.Println(len(x))
	return str
}
