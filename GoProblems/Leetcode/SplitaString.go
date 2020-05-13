package main

import "fmt"

func main() {
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
}

func balancedStringSplit(s string) int {
	val := []rune(s)
	count := 0
	x := 0
	for i := 0; i < len(val); i++ {
		if val[i] == 76 {
			x++
		} else {
			x--
		}

		if x == 0 {
			count++
		}
	}
	return count
}
