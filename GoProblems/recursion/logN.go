package main

import "fmt"

func main() {
	fmt.Println(reverse(21, 0))

	//log2(21)
}

func reverse(n, x int) int {
	if n == 0 {
		return x
	}
	//fmt.Println(n, x)
	//x = x * 10
	//x = x + n%10
	return reverse(n/10, (x*10 + n%10))

}

func log2(n int) {
	if n == 0 {
		return
	}
	log2(n / 2)
	fmt.Println("n", n)
	fmt.Println(n % 2)

}

func log(n int) int {
	if n == 1 {
		return 0
	}

	return 1 + log(n/2)
}
