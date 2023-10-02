package main

func main() {

	cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("mycards")

	// cards, _ := newDeckFromFile("mycards")

	// cards.print()
	cards.shuffle()
	cards.print()
}
