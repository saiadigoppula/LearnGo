package main

import "fmt"

func main() {

	s := "ccccdd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) int {
	val := []rune(s)
	mapp := map[rune]int{}
	count := 0
	for _, v := range val {

		if mapp[v] == 1 {
			count = count + 2
			delete(mapp, v)
		} else {
			mapp[v] = 1
		}
	}
	if len(val) > count {
		return count + 1
	}
	return count
}
