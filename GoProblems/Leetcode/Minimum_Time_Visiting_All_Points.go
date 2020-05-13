package main

import (
	"fmt"
	"math"
)

func main() {
	Points := [][]int{{3, 2}, {-2, 2}}

	//fmt.Println(Points[0][0])
	fmt.Println(minTimeToVisitAllPoints(Points))
}

func minTimeToVisitAllPoints(points [][]int) int {
	//fmt.Println(points[0][0])
	var count float64
	for i := 0; i < len(points)-1; i++ {
		if math.Abs(float64(points[i][0]-points[i+1][0])) >
			math.Abs(float64(points[i][1]-points[i+1][1])) {
			count = count + math.Abs(float64(points[i][0]-points[i+1][0]))
			//fmt.Println("if  :", count)
		} else {
			count = count + math.Abs(float64(points[i][1]-points[i+1][1]))
			//fmt.Println("if  :", count)
		}
	}
	return int(count)
}
