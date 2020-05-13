package main

import (
	"fmt"
	"math"
)

func main() {
	x := 100100111000000
	count := 0
	var res float64
	val := []int{}
	for x != 0 {
		if x%10 == 1 {
			val = append(val, count)
		}
		x = x / 10
		count++
	}
	fmt.Println(val)
	for _, v := range val {
		res = res + math.Pow(2, float64(v))
	}

	fmt.Println(int(res), count)
}
