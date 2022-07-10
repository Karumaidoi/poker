package main

import "fmt"

func main() {
	// cards := newDeck()
	cards := newDeck()
	cards.shuffleCards()
	d := []string(cards)
	fmt.Println(d)
}
