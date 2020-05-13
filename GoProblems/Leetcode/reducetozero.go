package main

import "fmt"

func main() {

	fmt.Println(numberOfSteps(8))

}

func numberOfSteps(num int) int {

	count := 0
	for num > 0 {
		if num%2 == 0 {
			num = num / 2
			count++
		} else {
			num--
			count++
		}
	}
	return count
}
