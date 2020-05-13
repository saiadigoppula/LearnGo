package main

import "fmt"

func main() {
	var (
		counter int
		p       *int
		v       int
	)
	counter = 100
	p = &counter
	v = *p

	fmt.Printf("value %d\n", v)
	passVal(&counter)
	fmt.Printf("value %d\n", counter)
	fmt.Printf("value %p\n", &counter)
}
func passVal(n *int) {
	*n = 50
	fmt.Printf("value %p\n", n)
	*n++

}
