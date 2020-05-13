package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "and I quote: &quot;...&quot;"
	fmt.Println(entityParser(s))

}

func entityParser(text string) string {
	str := ""
	bul := false
	val := strings.Fields(text)
	for _, v := range val {
		fmt.Println(v)
		if bul {
			str = str + fmt.Sprintf(" ")
		}
		bul = true
		switch {
		case v == "&quot;":
			str = str + `"`
			fmt.Println(str)
		case v == "&apos;":
			str = str + `'`
		case v == "&amp;":
			str = str + `&`
		case v == "&gt;":
			str = str + `>`
		case v == "&lt;":
			str = str + `<`
		case v == "&frasl;":
			str = str + `/`
		case v == "&quot;...&quot;":
			fmt.Println("Sammy says, \"Hello!\"")
			str = fmt.Sprintf(str + `\"...\"`)
		default:
			str = str + v

		}

	}
	return str
}
