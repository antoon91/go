package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("cards.txt")
	// cards = newDeckFromFile("cards.txt")
	cards.print()
}
