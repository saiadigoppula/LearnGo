package main

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	for _, it := range l {
		it.print()
	}
}
