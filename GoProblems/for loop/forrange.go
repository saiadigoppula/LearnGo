package main

import (
	"fmt"
	"strings"
)

func main() {

	words := strings.Fields(
		"lazy sai in in home",
	)

	for i, v := range words {
		fmt.Printf("#%-2d: %q\n",
			i+1, v)
	}

	/*for _, v := range os.Args[1:] {

		fmt.Printf("%q\n", v)
	}*/

}
