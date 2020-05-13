package main

import "fmt"

func main() {
	S, T := "a##c", "#a#c"
	//S, T := "xywrrmp", "xywrrm#p"
	val := backspaceCompare(S, T)
	fmt.Print(val)

}

func backspaceCompare(S string, T string) bool {
	s := []rune(S)
	t := []rune(T)
	_ = t
	for i := 0; i < len(s); i++ {
		//fmt.Println(i)
		if s[i] == '#' {
			if i != 0 {
				s = append(s[:i-1], s[i+1:]...)
				//fmt.Println("first", string(s))
				i = i - 2
			} else {
				s = append(s[:0], s[i+1:]...)
				//fmt.Println("second", string(s))
				i = i - 1
			}
			//fmt.Println(string(s[i]))
		}
	}
	//fmt.Println(len(s))

	for i := 0; i < len(t); i++ {
		//fmt.Println(i)
		if t[i] == '#' {
			if i != 0 {
				t = append(t[:i-1], t[i+1:]...)
				//fmt.Println("first", string(s))
				i = i - 2
			} else {
				t = append(t[:0], t[i+1:]...)
				//fmt.Println("second", string(s))
				i = i - 1
			}
			//fmt.Println(string(s[i]))
		}
	}
	//fmt.Println(len(t))
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if s[i] != t[i] {
			return false
		}
	}

	return true
}
