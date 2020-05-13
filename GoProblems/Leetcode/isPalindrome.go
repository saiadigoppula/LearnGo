package main

import "fmt"

func main() {

	s := -121
	fmt.Println(isPalindrome(s))
}

func isPalindrome(x int) bool {
	if x <= 0 {
		return false
	}
	n := x
	rev := 0
	for n > 0 {
		rev = rev * 10
		rev = rev + n%10
		n = n / 10
	}
	//fmt.Println(rev)
	if rev == x {
		return true
	}

	return false
}
