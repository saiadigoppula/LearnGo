package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello go")
	bye()
	if 5 > 1 {
		//fmt.Println("in if block" + "!")
		//fmt.Println("World")
		fmt.Println(runtime.NumCPU() + 1)
	}
}
