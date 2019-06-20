package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	cards.saveToFile("my_cards")
	loadedDeck := readDeckFromFile("my_cards")
	loadedDeck.shuffle()
	loadedDeck.print()

	hand, remainingCards := deal(loadedDeck, 5)
	hand.print()
	remainingCards.print()
}
