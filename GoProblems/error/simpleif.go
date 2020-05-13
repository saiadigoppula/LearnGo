package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if a := os.Args; len(a) != 2 {
		fmt.Println("give a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		fmt.Println("cannot convert %q.\n", a[1])
	} else {
		fmt.Printf("converted int is %d\n", n)
	}
}
