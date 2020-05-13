package main

import "fmt"

type MinStack struct {
	data    []int
	min     []int
	minimum int
	count   int
}

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())

}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {

	if len(this.min) == 0 || this.min[this.minimum-1] >= x {
		this.min = append(this.min, x)
		this.minimum++
	}

	this.data = append(this.data, x)
	this.count++
	/*fmt.Println(this.data)
	fmt.Println(this.min)
	fmt.Println(this.count)
	fmt.Println(this.minimum)*/
	//fmt.Println(len(this.data))

}

func (this *MinStack) Pop() {
	if this.data[this.count-1] == this.min[this.minimum-1] {
		this.min = this.min[:this.minimum-1]
		this.minimum--
	}
	this.data = this.data[:this.count-1]
	this.count--
}

func (this *MinStack) Top() int {
	//fmt.Println(this.data)
	return this.data[this.count-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.minimum-1]
}
