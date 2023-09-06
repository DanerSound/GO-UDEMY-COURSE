package main

func main() {

	cards := newDeck()
	cards.print()
	cards.shuffle()
	cards.print()
	// hand, remainCards := deal(cards, 5)
	// hand.print()
	// remainCards.print()
	// cards := newDeckFromFile("my_.txt")
	// cards.print()
	// cards.saveToFile("my_cards.txt")
}
