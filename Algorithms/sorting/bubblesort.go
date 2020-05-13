package main

import "fmt"

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	bubble(arr)
	fmt.Println(arr)
}

func bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				x := arr[i]
				arr[i] = arr[j]
				arr[j] = x
			}
		}
	}
}
