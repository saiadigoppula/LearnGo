package main

import "fmt"

func main() {
	val := []int{1, 2, 3, 4, 5, 6, 7, 8}
	x := val[2]
	val = append(val[:2], val[2+1:]...)
	val = append(val, x)
	val = append(val[len(val)-1:], val[:len(val)-1]...)
	fmt.Println(val, x)
}
