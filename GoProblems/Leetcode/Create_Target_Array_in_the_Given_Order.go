package main

import "fmt"

func main() {
	//nums = [0,1,2,3,4], index = [0,1,2,2,1]
	nums := []int{1, 2, 3, 4, 0}
	index := []int{0, 1, 2, 3, 0}
	fmt.Println(createTargetArray(nums, index))
}

func createTargetArray(nums []int, index []int) []int {
	//val := make([]int, len(index))
	val := []int{}
	for i := 0; i < len(index); i++ {
		//fmt.Println(val[i])
		//val[index[i]] = nums[i]
		if len(val) == index[i] {
			val = append(val, nums[i])
		} else {
			buff := []int{}
			buff = append(buff, val[index[i]:]...)
			val = append(val[:index[i]], nums[i])
			val = append(val, buff...)

		}
	}
	return val
}
