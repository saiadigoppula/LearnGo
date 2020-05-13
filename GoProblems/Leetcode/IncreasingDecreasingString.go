package main

import (
	"fmt"
	"strings"
)

func main() {
	//Input: s = "leetcode"
	//Output: "cdelotee"

	fmt.Println(sortString("leetcode"))
}

func sortString(s string) string {

	if len(s) <= 1 {
		return s
	}
	valkey := []rune{}
	valval := []int8{}

	for _, v := range s {
		x := 0
		for i := 0; i < len(valkey); i++ {
			if valkey[i] == v {
				valval[i]++
				//fmt.Println("valkey", valkey)
				//fmt.Println("val ", valval)
				x = 1
				break
			}
		}
		if x == 0 {
			valkey = append(valkey, v)
			valval = append(valval, 1)
		}

	}
	fmt.Println("valkey", valkey)
	fmt.Println("val ", valval)
	for i := range valkey {
		for j := i + 1; j < len(valkey); j++ {
			if valkey[i] > valkey[j] {
				x := valkey[i]
				y := valval[i]
				valkey[i] = valkey[j]
				valval[i] = valval[j]
				valkey[j] = x
				valval[j] = y
				fmt.Println("valkey", string(valkey[i]))
				fmt.Println("val ", valval[i])
			}
		}
	}
	fmt.Println("valkey", valkey)
	fmt.Println("val ", valval)

	fmt.Println("valkey", string(valkey[1]))
	fmt.Println("val ", valval[1])
	//str := ""
	var str strings.Builder
	for true {
		z := 0
		for i := 0; i < len(valkey); i++ {
			if valval[i] > 0 {
				str.WriteString(string(valkey[i]))
				valval[i]--
				z++
			}
		}

		for j := len(valkey) - 1; j >= 0; j-- {
			if valval[j] > 0 {
				str.WriteString(string(valkey[j]))
				valval[j]--
				z++
			}
		}

		if z == 0 {
			return str.String()
		}
	}
	return str.String()
}
