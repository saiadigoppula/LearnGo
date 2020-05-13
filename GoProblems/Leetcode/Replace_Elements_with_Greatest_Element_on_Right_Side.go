package main

import (
	"fmt"
)

func main() {
	arr := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(arr))
}
func replaceElements(arr []int) []int {
	val := arr[len(arr)-1]
	for i := len(arr) - 1; i >= 0; i-- {
		//if i != len(arr)-1 {
		if val < arr[i] {
			val = val + arr[i]
			arr[i] = val - arr[i]
			val = val - arr[i]
		} else {
			arr[i] = val
		}
		/*} else {
			val = arr[i]
			arr[i] = -1
		} */
	}
	arr[len(arr)-1] = -1
	return arr
}
