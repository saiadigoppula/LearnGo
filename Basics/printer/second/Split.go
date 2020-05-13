package split

import (
	"fmt"
	"path"
)

func main() {

	//var file string
	_, file := path.Split("css/main.css")

	//fmt.Println("dir", dir)
	fmt.Println("file", file)

}
