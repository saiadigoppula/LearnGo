package main

import "fmt"

func main() {

	fmt.Println(numJewelsInStones("z", "ZZ"))
}

func numJewelsInStones(J string, S string) int {
	js := []rune(J)
	ss := []rune(S)
	count := 0
	for _, v := range js {
		for j, y := range ss {
			if string(v) == string(y) {
				ss[j] = '.'
				count++
			}
		}
	}

	return count
}
