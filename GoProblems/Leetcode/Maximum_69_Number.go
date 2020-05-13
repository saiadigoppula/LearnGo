package main

import (
	"fmt"
)

func main() {
	//arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(maximum69Number(9999))
}

func maximum69Number(num int) int {
	count := 0
	x := 1
	y := num
	for num != 0 {
		if num%10 == 6 {
			count = 3 * x
		}
		x = x * 10
		num = num / 10
	}
	return count + y
}
