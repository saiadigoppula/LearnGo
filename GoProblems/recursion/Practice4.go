package main

import "fmt"

func main() {
	arr := []int{12, 10, 30, 50, 100}
	fun(5)
	fmt.Println()
	fmt.Println(fun2(arr, 5))
	fmt.Println(func3(200))
}

func func3(i int) int {

	i--
	//fmt.Println(i)
	if i%2 == 0 {
		i++
		return i
	} else {
		return func3(func3(i))
	}

}

func fun2(arr []int, n int) int {
	var x int
	if n == 1 {
		return arr[0]
	}
	x = fun2(arr, n-1)

	//fmt.Println("n", n-1)
	if x > arr[n-1] {
		return x
	}
	return arr[n-1]
}

func fun(n int) {
	// /fmt.Println("start", n)
	if n <= 0 {
		//fmt.Println("return-----", n)
		return
	}
	n--
	fun(n)
	fmt.Print(n)
	n--
	fun(n)
}
