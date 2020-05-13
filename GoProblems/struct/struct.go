package main

import "fmt"

func main() {

	type picasso struct {
		name, lastname string
		age            int
	}

	sai := picasso{
		"sai",
		"kick",
		30,
	}
	fmt.Println(sai)

}
