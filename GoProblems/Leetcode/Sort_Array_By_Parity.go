package main

import (
	"fmt"
)

func main() {
	A := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(A))
}

func sortArrayByParity(A []int) []int {
	x := 0
	for i := 0; i < len(A)-x; i++ {
		if A[i]%2 != 0 {
			z := A[i]
			A = append(A[:i], A[i+1:]...)
			A = append(A, z)
			i--
			x++
		}
	}
	return A
}
