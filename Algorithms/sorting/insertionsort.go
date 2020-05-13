package main

import "fmt"

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	insertion(arr)
	fmt.Println(arr)
}

func insertion(arr []int) {

	for i := 0; i < len(arr); i++ {
		key := arr[i]

		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
