package main

import "fmt"

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quicksort(arr []int, low, high int) {
	if low < high {

		pi := partition(arr, low, high)

		quicksort(arr, low, pi-1)
		quicksort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	i := low - 1
	pi := arr[high]
	for j := low; j < high; j++ {

		if pi > arr[j] {
			i++
			swap(&arr[i], &arr[j])
			//fmt.Println("after swap  pi", arr[i], "----arr----", arr[j])
		}
	}
	swap(&arr[i+1], &arr[high])
	return i + 1
}

func swap(pi, arr *int) {
	//fmt.Println("pi", *pi, "----arr----", *arr)
	*pi = *pi + *arr
	*arr = *pi - *arr
	*pi = *pi - *arr
	//fmt.Println("after pi", *pi, "----arr----", *arr)
}
