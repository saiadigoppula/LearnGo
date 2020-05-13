package main

import "fmt"

func main() {

	dict := map[string]string{
		"a": "sai",
		"b": "rick",
		"c": "kick",
	}
	turkish := make(map[string]string, len(dict))

	for k, v := range dict {
		turkish[k] = v
	}

	fmt.Printf("Zero Value : %#v\n", turkish)
}
