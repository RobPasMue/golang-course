package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cS := range cardSuits {
		for _, cV := range cardValues {
			cards = append(cards, cV+" of "+cS)
		}
	}

	return cards
}

// Receiver functions

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) shuffle() {
	// Properly randomize
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	sizeDeck := len(d) - 1
	for i := range d {
		newPosition := r.Intn(sizeDeck)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// Normal functions

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(filename string) error {
	deckAsString := d.toString()
	return os.WriteFile(filename, []byte(deckAsString), 0666)
}

func newDeckFromFile(filename string) (deck, error) {
	// Read the file into a byte slice
	bs, err := os.ReadFile(filename)

	// Check for error... empty Deck
	if err != nil {
		// Option 1 - Return a new deck + error
		// return newDeck(), err
		// Option 2 - Exit program with error code 1
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Split the byte slice into a string slice
	deckAsStringSlice := strings.Split(string(bs), ",")

	// Return results
	return deck(deckAsStringSlice), nil
}
