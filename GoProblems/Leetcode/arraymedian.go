package main

import "fmt"

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4, 5}

	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var x, y, xcnt, ycnt int
	for i := 0; i <= int((len(nums1)+len(nums2))/2+1)-1; i++ {
		y = x
		//fmt.Println("y", y)
		if len(nums1)-1 >= xcnt {
			if len(nums2)-1 >= ycnt {
				if nums1[xcnt] <= nums2[ycnt] {
					x = nums1[xcnt]
					xcnt++
				} else {
					x = nums2[ycnt]
					ycnt++
				}

			} else {
				x = nums1[xcnt]
				xcnt++
			}

		} else {
			x = nums2[ycnt]
			ycnt++
		}
	}
	//fmt.Println("x", x)

	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(x+y) / 2
	}
	return float64(x)

}
