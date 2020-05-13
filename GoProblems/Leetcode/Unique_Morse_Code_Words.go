package main

import (
	"fmt"
	"strings"
)

func main() {
	words := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words))
}
func uniqueMorseRepresentations(words []string) int {
	val := map[string]string{"a": ".-", "b": "-...", "c": "-.-.", "d": "-..", "e": ".", "f": "..-.", "g": "--.",
		"h": "....", "i": "..", "j": ".---", "k": "-.-", "l": ".-..", "m": "--", "n": "-.", "o": "---", "p": ".--.",
		"q": "--.-", "r": ".-.", "s": "...", "t": "-", "u": "..-", "v": "...-", "w": ".--", "x": "-..-", "y": "-.--", "z": "--.."}
	carry := make(map[string]int)

	for i := range words {
		//str := ""
		var b strings.Builder
		for _, j := range words[i] {
			//str = str + val[string(j)]
			b.WriteString(val[string(j)])

		}

		carry[b.String()]++

	}

	return len(carry)

}
