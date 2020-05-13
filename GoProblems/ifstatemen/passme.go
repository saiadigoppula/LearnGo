package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q.\n"
	errPwd   = "Invalid password for %q.\n"
	accessOk = "Access granted to %q.\n"

	user, user2 = "jack", "sai"
	pass, pass2 = "1888", "1997"
)

func main() {

	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user && u != user2 {
		fmt.Printf(errUser, u)
	} else if u == user && p == pass {
		fmt.Printf(accessOk, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOk, u)
	} else {
		fmt.Printf(errPwd, u)
	}

}
