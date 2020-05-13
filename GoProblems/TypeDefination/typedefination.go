package main

import "fmt"

type gram float64
type ounce float64

func main() {

	var g gram = 1000
	var o ounce
	o = ounce(g) * 0.035274
	fmt.Printf("%g grams is %.2f ounces\n", g, o)
}
