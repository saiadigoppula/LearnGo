package main

import "fmt"

func main() {

	queries := []int{3, 1, 2, 1}
	m := 5
	fmt.Println(processQueries(queries, m))

}

func processQueries(queries []int, m int) []int {

	result := []int{}
	p := []int{}
	// (from i=0 to i=queries.length-1)
	for _, v := range queries {
		for j := 1; j <= m; j++ {
			if j > len(p) {
				p = append(p, j)
				//fmt.Println("firest ----", p)
			}
			if p[j-1] == v {
				//fmt.Println(j-1, p[j-1])
				result = append(result, j-1)
				//fmt.Println("result", result)
				val := p[j-1]
				//fmt.Println("val", val)
				//fmt.Println(p[:j-1])
				//fmt.Println(p[j:])
				//p = append(p[:j-1], p[j:]...)
				//fmt.Println("p at resize", p)

				//p = append(p[:j-1], p[j:]...)
				//p = append(p, val)
				//p = append(p[len(p)-1:], p[:len(p)-1]...)
				//fmt.Println("p at resize", p)
				for x := j - 1; x > 0; x-- {
					//fmt.Println(j)
					p[x-1] = p[x] + p[x-1]
					p[x] = p[x-1] - p[x]
					p[x-1] = p[x-1] - p[x]
				}
				p[0] = val
				//fmt.Println("final p", p)
				break
			}
		}
	}
	return result
}
