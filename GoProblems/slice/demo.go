package main

func main() {
	slice := [][]string{}

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	x, y := -1, 1
	Cmax := 0
	Rmax := 0
	for i := 0; i < len(strs); i++ {
		y--
		x++
		slice = append(slice, []string{strs[i]})

		if y == 0 {
			Cmax++
			y = Cmax + 1
		}
		if x == Rmax {
			Rmax++
			x = -1
		}

	}
	println(slice)
}
