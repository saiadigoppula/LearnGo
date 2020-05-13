package main

import "fmt"

func main() {

	type (
		Gram     int
		Kilogram Gram
		Ton      Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n",
		salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(apples)
	fmt.Printf("salt: %d, apples: %d, truck: %d\n",
		salt, apples, truck)

}
