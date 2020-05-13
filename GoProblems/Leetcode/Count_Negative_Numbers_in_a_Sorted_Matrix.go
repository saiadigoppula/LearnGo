package main

import (
	"fmt"
)

func main() {
	//grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	grid := [][]int{{-1}}
	fmt.Println(countNegatives(grid))
}
func countNegatives(grid [][]int) int {
	count := 0
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if grid[i][j] >= 0 {
				break
			}
			count++
		}

	}
	return count
}
