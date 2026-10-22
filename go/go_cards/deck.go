package main

type deck []string

func (d deck) print() {
	for index, card := range d {
		println(index, card)
	}
}
