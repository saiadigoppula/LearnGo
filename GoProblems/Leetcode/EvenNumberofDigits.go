package main

import "fmt"

func main() {
	nums := []int{555, 901, 482, 1771}
	fmt.Println(findNumbers(nums))
}

func findNumbers(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		x := 0
		for nums[i] != 0 {
			nums[i] = nums[i] / 10
			x++
		}
		if x%2 == 0 {
			count++
		}
	}
	return count
}
