package main

import "fmt"

func main() {

	//Input: "(()())(())(()(()))"
	//Output: "()()()()(())"
	s := "(()())(())(()(()))"
	fmt.Println(removeOuterParentheses(s))
}

func removeOuterParentheses(S string) string {

	//	val := []rune(S)
	str := ""
	x := 0
	//fmt.Println(val[2])
	//counter != 0 && !(counter == 1 && c == ')')
	for _, val := range S {
		if x != 0 && !(x == 1 && val == ')') {
			str = str + string(val)
			//x++
		}
		if val == '(' {
			x++
		} else {
			x--
		}

	}
	return str

	/*val := []rune(S)
	str := ""
	x := 0
	//fmt.Println(val[2])
	for i := 0; i < len(val); i++ {
		if x > 0 && val[i] == 40 {
			str = str + "("
			x++
		} else if x > 1 && val[i] == 41 {
			str = str + ")"
			x--
		} else if val[i] == 40 {
			x++
		} else {
			x--
		}

	}
	return str
	*/
}
