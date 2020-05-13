package main

import "fmt"

func main() {

	data := []float64{10, 25, 30, 50}

	newData := []float64{99, 100, 10, 10, 10}
	copy(data, newData)
	data = append(data[:0], newData...)

	fmt.Println(data)
}
