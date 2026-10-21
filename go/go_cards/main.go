package main

import "fmt"

func main() {
	card := newCard()
	cardNumber := newCardNumber()
	fmt.Println(card)
	fmt.Println("Card Number:", cardNumber)

	cardsValue := []string{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	fmt.Println("Cards Value:", cardsValue)

	cardsValue = append(cardsValue, "Four of Clubs")
	fmt.Println("Updated Cards Value:", cardsValue)
}

func newCard() string {
	return "Five of Diamonds"
}

func newCardNumber() int {
	return 5
}
