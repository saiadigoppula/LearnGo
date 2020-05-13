package main

import "fmt"

func main() {

	fmt.Println(sum(10, 10))
}

func sum(x, y int) int {
	if x == 0 {
		return y
	}
	fmt.Println(x, y)
	return sum(x-1, x+y)
}
