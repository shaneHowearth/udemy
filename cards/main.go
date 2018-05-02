package main

import "fmt"

func main() {
	/*
		cards := newDeck()

		hand, remainingCards := deal(cards, 5)
		remainingCards.print()
		hand.print()
	*/
	cards := newDeck()
	//fmt.Println([]byte(cards.toString()))
	cards.saveToFile("my_cards")
	nd := newDeckFromFile("my_cards")
	//nd.shuffle()
	fmt.Println(nd)

}
