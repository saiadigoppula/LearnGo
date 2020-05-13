package main

import "fmt"

/*node */
type Node struct {
	data int
	next *Node
}

var head *Node
var count int

func main() {

	A := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range A {
		insert(v)
	}
	//display()
	insertPosition(1000, 3)
	display()

}

func insertPosition(A, p int) {

	if p > count || p <= 0 {
		fmt.Println("invalid position")
		return
	}
	copy := head

	for i := 0; i < p-2; i++ {
		//println("true")
		copy = copy.next
	}
	//fmt.Println("val", copy.data)
	var sai Node
	sai.data = A
	sai.next = copy.next
	copy.next = &sai

	count++

}

func display() {
	copy := head
	if copy.next == nil {
		fmt.Println("empty list")
		return
	}
	for true {
		if copy.next != nil {
			println(copy.data)
			copy = copy.next
		} else {
			println(copy.data)
			break
		}
	}
	fmt.Println("count", count)
	fmt.Println()

}

func insert(A int) {

	if head == nil {
		var sai Node
		sai.data = A
		head = &sai
		count++
	} else {
		copy := head
		for true {
			if copy.next == nil {
				break
			} else {
				copy = copy.next
			}
		}
		var sai Node
		sai.data = A
		copy.next = &sai
		count++

	}

}
