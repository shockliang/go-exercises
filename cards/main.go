package main

func main() {
	//cards := newDeck()
	//cards.saveToFile("my-cards")
	cards := newDeckFromFile("my-cards")
	cards.print()
}

