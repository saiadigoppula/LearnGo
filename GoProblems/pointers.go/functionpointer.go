package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("ARRAYS")
	arrays()

	fmt.Println("\n....SLICE")
	slices()

	fmt.Println("\n....MAPS")
	maps()

	fmt.Println("\n... STRUCE")
	structs()
}

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(&myHouse)

	fmt.Printf("%+v\n", myHouse)
}

func addRoom(h *house) {
	h.rooms++
}

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

func slices() {
	dirs := []string{"up", "down", "left", "right"}
	//up(dirs)
	//fmt.Println(dirs)
	upPtr(&dirs)
	fmt.Println(dirs)
}

func upPtr(list *[]string) {
	lv := *list
	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}
	*list = append(*list, "Hi sir")
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
	}
}

func arrays() {
	nums := [...]int{1, 2, 3}
	incr(&nums)
	fmt.Println(nums)
}

func incr(nums *[3]int) {
	for i := range nums {
		nums[i]++
	}
}
