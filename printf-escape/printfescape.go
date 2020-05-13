package main

import "fmt"

func main() {
	var speed int = 20
	var heat float64 = 50.35
	var off bool = true
	var brand string = "sai"

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	fmt.Printf("%v\n", speed)
	fmt.Printf("%.1f so %v and %[2]v  %[1]v\n", heat, off)
	fmt.Printf("%v\n", off)
	fmt.Printf("%q\n", brand)

}
