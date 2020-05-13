package main

import "fmt"

func main() {

	fmt.Println(subtractProductAndSum(234))
}

func subtractProductAndSum(n int) int {
	x := 1
	y := 0

	for n != 0 {
		val := n % 10
		n = n / 10
		x = x * val
		y = y + val
		//fmt.Println(x, y)
	}

	return x - y

}
