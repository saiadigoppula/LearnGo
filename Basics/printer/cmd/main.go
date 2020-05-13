package main

import (
	"fmt"
	"learngo/printer"
)

func main() {
	printer.Hello()
	fmt.Println(
		-200, -100, 0, 50, 100,
	)

	fmt.Println(
		-50.5, -20.5, -0., 1.0,
	)

	fmt.Println(
		true, false,
	)
	fmt.Println(0x0, 0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9)
	fmt.Println(0xa, 0xb, 0xc, 0xd, 0xe, 0xf)
	fmt.Println(0x11) // 17
	fmt.Println(0x19) // 25
	fmt.Println(0x32) // 50
	fmt.Println(0x64) // 100
}
