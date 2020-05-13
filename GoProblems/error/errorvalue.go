package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println(s)

	i, err := strconv.Atoi(os.Args[1])

	fmt.Println("converted number  :", i)
	fmt.Println("error   :", err)

}
