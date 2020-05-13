package main

import (
	"fmt"
)

func main() {
	//Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
	//Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]

	A := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	fmt.Println(A)
	fmt.Println(flipAndInvertImage(A))
}

func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		//fmt.Println("", i)
		x := len(A[i]) - 1
		//fmt.Println("len", len(A[i])/2)
		for j := 0; j < len(A[i])/2; j++ {

			z := val(A[i][j])
			//fmt.Println("A[i][x]     ", A[i][x], "A[i][x]", val(A[i][x]))
			//fmt.Println("xxxxx----", x, "-----xxxxxxx------", A[i][x])
			y := val(A[i][x])
			//fmt.Println("-----z----", z, "------y-----", y)
			A[i][j] = y
			A[i][x] = z
			//fmt.Println("A[i]-----------------------", A[i])

			x--

		}
		if len(A[i])%2 == 1 {
			A[i][len(A[i])/2] = val(A[i][len(A[i])/2])
		}
		//fmt.Println("after", A[i])
	}
	return A
}

func val(x int) int {
	if x == 0 {
		return 1
	} else {
		return 0
	}
}

/*	if A[i][0] == 0 {
		A[i][0] = 1
	} else {
		A[i][0] = 0
	} */
//A[i] = append(A[i][1:], A[i][:1]...)
