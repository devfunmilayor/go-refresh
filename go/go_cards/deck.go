package main

import "fmt"

type deck []string

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func (d deck) printWithIndex() {

	if len(d) == 0 {
		fmt.Println("The deck is empty.")
		return
	} else if len(d) == 1 {
		fmt.Println("The deck has only one card:", d[0])
		return
	} else if len(d) == 2 {
		fmt.Println("The deck has two cards:", d[0], "and", d[1])
		return
	} else if len(d) == 3 {
		fmt.Println("The deck has three cards:", d[0], ",", d[1], "and", d[2])
		return
	} else if len(d) == 4 {
		fmt.Println("The deck has four cards:", d[0], ",", d[1], ",", d[2], "and", d[3])
		return
	} else if len(d) == 5 {
		fmt.Println("The deck has five cards:", d[0], ",", d[1], ",", d[2], ",", d[3], "and", d[4])
		return
	} else {
		fmt.Println("The deck has more than five cards.")
	}
}

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
