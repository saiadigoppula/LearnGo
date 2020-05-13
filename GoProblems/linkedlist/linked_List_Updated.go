package main

import "fmt"

//Node...
type Node struct {
	val  int
	next *Node
}

var head *Node

func main() {

	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range A {
		insert(v)
	}
	display()
	insertHead(10)
	display()
	insertPosition(8, 100)
	display()
	removeHead()
	display()
}

func removeHead() {
	if head == nil {
		fmt.Println("list is empty")
		return
	}

	head = head.next

}

func insertPosition(p, v int) {
	if head == nil {
		fmt.Println("list is empty")
		return
	}
	copy := head

	for copy.next != nil {
		if p == 1 {
			break
		}
		copy = copy.next
		p--
	}
	if p == 1 {
		var data Node
		data.val = v

		data.next = copy.next
		copy.next = &data

	} else {
		fmt.Println("position is wrong")
	}
}

func insertHead(v int) {
	var data Node
	data.val = v
	if head == nil {
		head = &data
	} else {
		data.next = head
		head = &data
	}
}

func insert(v int) {
	var data Node
	data.val = v
	if head == nil {
		head = &data
	} else {
		copy := head

		for copy.next != nil {
			copy = copy.next
		}
		copy.next = &data
	}

}

func display() {

	if head == nil {
		return
	}

	copy := head

	for copy != nil {
		fmt.Print(copy.val, " ")
		copy = copy.next
	}
	fmt.Println()
}
