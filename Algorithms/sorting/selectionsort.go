package main

import "fmt"

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	selection(arr)
	fmt.Println(arr)
}

func selection(arr []int) {

	for i := 0; i < len(arr); i++ {
		x := arr[i]
		index := i
		for j := i + 1; j < len(arr); j++ {
			if x > arr[j] {
				x = arr[j]
				index = j
			}

		}
		//fmt.Println("arr[i]   ", arr[i], "     arr[index]    ", arr[index])
		val := arr[i]
		arr[i] = arr[index]
		arr[index] = val

	}
}
