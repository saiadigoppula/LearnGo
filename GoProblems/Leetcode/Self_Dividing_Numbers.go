package main

import (
	"fmt"
)

func main() {
	fmt.Println(selfDividingNumbers(1, 22))
}

func selfDividingNumbers(left int, right int) []int {

	val := []int{}
FirstLoop:
	for i := left; i <= right; i++ {
		x := i
		for x != 0 {
			if (x % 10) == 0 {
				continue FirstLoop
			}
			if i%(x%10) != 0 {
				//fmt.Println("i", i, "i/(x%10) ", i/(x%10))
				continue FirstLoop
			}
			x = x / 10
		}
		val = append(val, i)
	}
	return val
}
