package main

import "fmt"

func main() {
	num := []int{1, 1, 2, 3}
	fmt.Println(decompressRLElist(num))
}

func decompressRLElist(nums []int) []int {
	val := []int{}
	for i := 0; i < len(nums); i += 2 {
		//fmt.Println("sai")
		for j := 0; j < nums[i]; j++ {
			//fmt.Println(nums[i+1])
			val = append(val, nums[i+1])
		}
	}
	return val
}
