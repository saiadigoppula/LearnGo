package main

import "fmt"

func main() {
	nums := []int{8, 1, 2, 2, 3}
	fmt.Println(smallerNumbersThanCurrent(nums))
}

func smallerNumbersThanCurrent(nums []int) []int {
	val := []int{}
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		val = append(val, count)
	}

	return val

}
