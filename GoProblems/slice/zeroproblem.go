package main

import "fmt"

func main() {

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)

}

func moveZeroes(nums []int) {
	count := 0
	for i := 0; i < len(nums)-count; i++ {
		if nums[i] == 0 && i != len(nums)-1-count {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			fmt.Println(nums)
			i--
			count++
		}
	}

}
