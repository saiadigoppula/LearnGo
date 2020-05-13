package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)

	fmt.Println("path:", os.Args[0])
	fmt.Println("1st path:", os.Args[1])
	fmt.Println("2nd path:", os.Args[2])
	fmt.Println("3rd path:", os.Args[3])

	fmt.Println("Number of items inside os.Args :",
		len(os.Args))
}
