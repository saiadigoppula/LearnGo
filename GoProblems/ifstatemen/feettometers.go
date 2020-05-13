package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const usage = "        plese tell me a value in feet"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(strings.TrimSpace(usage))
		return
	}
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg, 64)
	meters := feet * 0.3048

	fmt.Printf("%g feet in %g meters.\n", feet, meters)

}
