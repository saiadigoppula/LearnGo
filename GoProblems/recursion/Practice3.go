package main

import "fmt"

func main() {
	pattern(10)
	pattern2(10)

}

func pattern2(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" * ")
	}
	fmt.Println()
	if n > 0 {
		pattern2(n - 1)
	}
}

func pattern(n int) {

	if n > 0 {
		pattern(n - 1)
	}
	for i := 0; i < n; i++ {
		fmt.Print(" * ")
	}
	fmt.Println()
}
