package main

import (
	"fmt"
)

func main() {
	city := "Lyon"

	switch city {
	case "paris", "Lyon":
		fmt.Println("French")
	case "Tokyo":
		fmt.Println("japan")

	default:
		fmt.Println("sai")
	}
	i := 10
	switch {
	case i > 0:
		fmt.Printf("%d is a positive number", i)
	case i < 0:
		fmt.Printf("%d is a negative  number", i)
	default:
		fmt.Printf("%d is zero", i)
	}

}
