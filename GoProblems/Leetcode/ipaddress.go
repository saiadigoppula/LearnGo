package main

import "fmt"

func main() {

	s := "1.1.1.1"
	fmt.Println(defangIPaddr(s))
}

func defangIPaddr(address string) string {

	ads := []rune(address)
	str := ""
	for _, v := range ads {
		if string(v) == "." {
			//ads = append(ads[:i], '[')
			str = str + "[.]"
		} else {
			str = str + string(v)
		}
		//fmt.Println(i, string(v))

	}
	return str
}
