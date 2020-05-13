package main

import (
	"fmt"
	"strings"
)

func main() {
	s := make([]int, 0, 5)
	s = append(s, 42)
	fmt.Println(s, len(s))

	tasks := []string{"jump", "hump", "dump"}

	upTasks := make([]string, 0, len(tasks))

	for _, task := range tasks {
		upTasks = append(upTasks, strings.ToUpper(task))
		fmt.Println(upTasks)
	}
	fmt.Println(upTasks, len(upTasks))

}
