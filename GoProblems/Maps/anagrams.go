package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	strs := [...]string{"eat", "tea", "tan",
		"ate", "nat", "bat"}

	dict := map[string][]string{}
	res := make([][]string, 0, 5)

	for _, v := range strs {
		x := SortString(v)
		if dict[x] == nil {
			dict[x] = []string{v}
		} else {
			dict[x] = append(dict[x], v)
		}
	}

	fmt.Println(dict)
	for _, v := range dict {
		res = append(res, v)
	}
	fmt.Println(res)

}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

/*m := map[string][]string{
	"cat": {"orange", "grey"},
	"dog": {"black"},
}

// Add a string at the dog key.
// ... Append returns the new string slice.
res := append(m["dog"], "brown")

    // Add a key for fish.
	m["fish"] = []string{"orange", "red"}


*/
