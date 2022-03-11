package main

func main() {
	// var card string = "Ace of Spades"

	cards := newDeck()
	cards.shuffle()
	// cards.saveToFile("my_cards")
	// cards := newDeckFromFile("my_cards")
	cards.print()
}
