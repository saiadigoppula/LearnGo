package main

import "fmt"

func main() {

	spendings := make([][]int, 0, 5)
	spendings = append(spendings, []int{200, 300})
	spendings = append(spendings, []int{202, 303})
	spendings = append(spendings, []int{220, 330})
	spendings = append(spendings, []int{110, 120})
	spendings = append(spendings, []int{20, 30}, []int{40, 50})

	for i, daily := range spendings {
		var total int

		for _, spending := range daily {
			total += spending

		}
		fmt.Printf("daily %d  spendings %d\n", i+1, total)
	}
	fmt.Println(spendings)
}
