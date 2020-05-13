package main

import "fmt"

func main() {

	fmt.Println(fact(5))
}

func fact(n int) int {
	if n == 1 {
		return 1
	}
	//fmt.Println(n)
	return n * fact(n-1)
}
