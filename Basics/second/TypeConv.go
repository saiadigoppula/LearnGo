package main

import (
	"fmt"
)

func main() {
	var (
		apple  int   = 1
		orange int32 = 65
	)

	apple = int(orange)
	fmt.Println(apple)
	color := string(orange)
	fmt.Println(color)
	fmt.Println(string([]byte{104, 105}))
}
