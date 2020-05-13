package main

import (
	"fmt"
	"strconv"
)

func main() {
	name, last := "sai", "krishna"

	fmt.Println(name + last)
	one := 10
	fmt.Println(strconv.Itoa(one) + name)

	eq := strconv.FormatBool(true) +
		strconv.FormatBool(false)
	fmt.Println(eq)
}
