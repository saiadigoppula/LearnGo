package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]

	feet, _ := strconv.ParseFloat(arg, 64)
	meter := feet * 0.3048
	fmt.Printf("%g feet is\n %g meeter.\n", feet, meter)

}
