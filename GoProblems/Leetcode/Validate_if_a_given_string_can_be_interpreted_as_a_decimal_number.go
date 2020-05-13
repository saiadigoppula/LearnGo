package main

import (
	"fmt"
)

func main() {
	fmt.Println(isNumber("0e1"))
}

func isNumber(s string) bool {
	space := false
	val := false
	any := false
	dot := false
	e := false
	opp := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if any == false && val == false {
				continue
			} else {
				space = true
				continue
			}

		}
		if s[i] == 'e' {
			if val == false || space == true || i == 0 || dot == true || e == true {

				return false
			} else {
				if i < len(s)-1 {
					if s[i+1] == '-' || s[i+1] == '+' {
						opp = false
						//plus = false
					}
				}
				if s[i-1] == '.' {
					any = true
					e = true
					continue
				}
				if s[i-1] < '0' || s[i-1] > '9' {
					return false
				}
				any = true
				e = true
				continue
			}
		}

		if s[i] == '-' || s[i] == '+' {
			if opp == true || any == false || val == false || space == true {
				return false
			}
			opp = true
			/*if s[i] == '+'{
			      plus = true
			  }else{
			       opp = true
			  }
			*/

			any = true
			continue

		}

		if s[i] == '.' {
			if dot == true || space == true || opp == true {
				return false
			}

			if i < len(s)-1 {
				if s[i+1] == 'e' {
					if i == 0 {
						return false
					}
					if s[i-1] < '0' || s[i-1] > '9' {
						return false
					}
				}
			}

			dot = true
			any = true
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {

			if space == true || opp == true {
				return false
			}
			val = true
			any = true
			continue
		}

		return false
	}
	if val == false {
		return false
	} else {
		return true
	}

}
