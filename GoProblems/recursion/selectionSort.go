package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	selection(arr, 0, len(arr)-1)
}

func selection(arr []int, start, end int) {
	if start >= end {
		fmt.Println(arr)
		return
	}

	ind := minindex(arr, start, end, start)

	x := arr[ind]
	arr[ind] = arr[start]
	arr[start] = x
	selection(arr, start+1, end)
}

func minindex(arr []int, start, end, count int) int {
	if start > end {
		return count
	}
	if arr[count] > arr[start] {
		count = start
	}
	return minindex(arr, start+1, end, count)

}
