package main

import "fmt"

func main() {
	var s string
	s = "how are you?"
	fmt.Println(s)
	s = `how are you?`
	fmt.Println(s)

	s = `
<html>
	<body>"hello"</body>
</html>`
	fmt.Println(s)
	fmt.Println(`c:\my\dir\file`)
}
