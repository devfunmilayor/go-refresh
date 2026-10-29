package main

import "fmt"

func main() {
	card := newCard()
	cardNumber := newCardNumber()
	fmt.Println(card)
	fmt.Println("Card Number:", cardNumber)

	cardsValue := deck{"Ace of Spades", "Two of Hearts", "Three of Diamonds", newCard()}
	cardsValue.print()
	cardsValue = append(cardsValue, "Four of Clubs")
	fmt.Println("Updated Cards Value:", cardsValue)

	cardsValue.printWithIndex()

	for index, card := range cardsValue {
		fmt.Println(index, card)
	}

	colors := []string{"Red", "Yellow", "Blue"}
	fmt.Println("Colors:", colors)

	colors = append(colors, "Green")
	fmt.Println("Updated Colors:", colors)

	for index, color := range colors {
		fmt.Println(index, color)
	}

}

func newCard() string {
	return "Five of Diamonds"
}

func newCardNumber() int {
	return 5
}
