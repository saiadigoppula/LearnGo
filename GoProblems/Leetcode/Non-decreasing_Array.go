package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 2, 3} //{3, 4, 2, 3} //{4, 2, 1} //{[2,3,3,2,4]}
	fmt.Println(checkPossibility(nums))
}
func checkPossibility(nums []int) bool {
	//boo := true
	cnt := 0
	for i := len(nums) - 1; i > 0; i-- {
		//fmt.Println()
		//fmt.Println(nums)
		//fmt.Println("i --------", nums[i], "-----i-1----", nums[i-1])
		if nums[i] < nums[i-1] {
			//fmt.Println("in iff loop")
			cnt++
			if cnt >= 2 {
				return false
			}
			if i < len(nums)-1 {
				if nums[i-1] > nums[i+1] {
					//return false
					nums[i-1] = nums[i]
					//fmt.Println(nums)
				}
				//nums[1] = nums[i+1]
			} else {
				nums[i] = nums[i-1]
			}

		}

	}
	return true
}
