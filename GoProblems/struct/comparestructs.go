package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string

	songs []song
}

func main() {

	song1 := song{"first", "first1"}
	song2 := song{"second", "second1"}

	fmt.Printf("song1:  %+v\n  song2:  %+v\n", song1, song2)

	rock := playlist{genre: "indie rock", songs: []song{
		song{title: "wonderwall", artist: "oasis"},
		song{"second", "second1"},
		song1,
	}}

	fmt.Printf("rock %+v", rock)

}
