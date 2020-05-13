package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 5

func main() {

	rand.Seed(time.Now().UnixNano())
	//fmt.Println(time.Now().UnixNano())

	args := os.Args[1:]
	guess, err := strconv.Atoi(args[0])

	if len(args) != 1 {
		fmt.Println("pick a number")
		return
	}

	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("please pick a number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n == guess {
			fmt.Println("You win")
			return

		}

	}
	fmt.Println("You Lost")
}
