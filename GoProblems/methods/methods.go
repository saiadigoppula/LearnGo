package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 100}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	var store list

	store = append(store, &minecraft, &tetris, mobydick)
	store.print()

}
