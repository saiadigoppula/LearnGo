package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type hold struct {
	head *Node
}

func main() {
	var head hold
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, v := range arr {
		head.append(v)
	}
	head.print()
	head.appendFirst(10)
	head.print()
	head.appendPosition(4, 20)
	head.print()
	head = head.removeFirst()
	head.print()
	head.removeTail()
	head.print()
	head.removePosition(4)
	head.print()
}

func (h *hold) removePosition(p int) {
	if h.head == nil {
		fmt.Println("empty list")
		return
	}
	copy := h.head
	for copy.next != nil {
		if p == 1 {
			break
		}
		p--
		copy = copy.next
	}
	if p == 1 {
		remove := copy.next
		copy.next = copy.next.next
		remove = nil
		_ = remove
		return
	}
	fmt.Println("invalid position")
	return

}

func (h *hold) removeTail() {
	if h.head == nil {
		fmt.Println("empty list")
		return
	}
	copy := h.head
	for copy.next.next != nil {
		copy = copy.next
	}
	copy.next = nil
}

func (h *hold) removeFirst() hold {
	copy := h.head
	h.head = h.head.next
	copy.next = nil
	fmt.Println(copy.val)
	return *h
}

func (h hold) appendPosition(p, v int) {

	data := Node{val: v}

	if h.head == nil {
		fmt.Println("empty list")
		return
	}
	copy := h.head
	for copy.next != nil {
		if p == 0 {
			break
		}
		p--
		copy = copy.next
	}

	if p == 0 {
		data.next = copy.next
		copy.next = &data
	} else {
		fmt.Println("not a correct position")
		return
	}

}

func (h *hold) appendFirst(v int) {
	data := Node{val: v}

	if h.head == nil {
		h.head = &data
		return
	}

	data.next = h.head
	h.head = &data
}

func (h *hold) append(v int) {
	//fmt.Print("in append")
	data := Node{val: v}
	if h.head == nil {
		h.head = &data
	} else {
		copy := h.head
		for copy.next != nil {
			copy = copy.next
		}

		copy.next = &data
	}

}

func (h hold) print() {
	if h.head == nil {
		fmt.Println("empty list")
		return
	}
	copy := h.head
	for copy != nil {
		fmt.Print(copy.val, " ")
		copy = copy.next
	}
	fmt.Println()
}
