package main

import "fmt"

func main() {
	fmt.Println("started")
	//i:=-100.0; i<-100.0; i+=0.1
	for i := 0.0; i < 10.0; i += 0.5 {
		for j := 0.0; j < 10.0; j += 0.5 {
			for x := 0.0; x < 10.0; x += 0.5 {
				for y := 0.0; y < 10.0; y += 0.5 {
					if (i+j) == 8 && (i+x) == 13 && (j+y) == 8 && (x-y) == 6 {
						fmt.Println(i, j, x, y)
						return
					}
				}
			}
		}
	}
}
