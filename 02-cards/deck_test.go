package main

import (
	"io"
	"os"
	"testing"
)

// captureStdout calls a function f and returns its stdout side-effect as string
func captureStdout(f func()) string {
	// return to original state afterwards
	// note: defer evaluates (and saves) function ARGUMENT values at definition
	// time, so the original value of os.Stdout is preserved before it is
	// changed further into this function.
	defer func(orig *os.File) {
		os.Stdout = orig
	}(os.Stdout)

	r, w, _ := os.Pipe()
	os.Stdout = w
	f()
	w.Close()
	out, _ := io.ReadAll(r)

	return string(out)
}

func TestNewDeckLength(t *testing.T) {
	myDeck := newDeck()

	// Check the length of the Deck is as expected
	deckLength := len(myDeck)
	if deckLength != 16 {
		t.Errorf("Deck length is not 16 (expected value) but %v", deckLength)
	}

}

func TestNewDeckCards(t *testing.T) {
	myDeck := newDeck()

	// Check the first and last cards
	firstCard := myDeck[0]
	if firstCard != "Ace of Spades" {
		t.Errorf("The first card is not 'Ace of Spades' (expected value) but %v", firstCard)
	}

	lastCard := myDeck[len(myDeck)-1]
	if lastCard != "Four of Clubs" {
		t.Errorf("The last card is not 'Four of Clubs' (expected value) but %v", lastCard)
	}
}

func TestDealHand(t *testing.T) {
	myDeck := newDeck()
	numCardsInHand := 5

	handDealed, myDeckDealed := deal(myDeck, numCardsInHand)

	if len(myDeckDealed) != (len(myDeck) - numCardsInHand) {
		t.Errorf("Deck length is not %v (expected value) but %v", (len(myDeck) - numCardsInHand), len(myDeckDealed))
	}

	if len(handDealed) != (numCardsInHand) {
		t.Errorf("Deck length is not %v (expected value) but %v", (numCardsInHand), len(handDealed))
	}
}

func TestSaveLoadDeck(t *testing.T) {

	// Create a Deck
	myDeck := newDeck()

	// Shuffle deck and store first and last cards
	myDeck.shuffle()

	expFirstCard := myDeck[0]
	expLastCard := myDeck[len(myDeck)-1]

	// Save to file
	file := t.TempDir() + "/deck_save_test.txt"
	myDeck.saveToFile(file)

	// Load from file
	myDeckLoaded, _ := newDeckFromFile(file)

	// Check the first and last cards
	firstCard := myDeckLoaded[0]
	if firstCard != expFirstCard {
		t.Errorf("The first card is not '%v' (expected value) but '%v'", expFirstCard, firstCard)
	}
	lastCard := myDeckLoaded[len(myDeckLoaded)-1]
	if lastCard != expLastCard {
		t.Errorf("The last card is not '%v' (expected value) but '%v'", expLastCard, lastCard)
	}

	// Check the entire deck is the same
	myDeckStr := captureStdout(myDeck.print)
	myDeckLoadedStr := captureStdout(myDeckLoaded.print)
	if myDeckStr != myDeckLoadedStr {
		t.Errorf("The initial deck (expected value)\n%v  is not like the loaded deck \n%v", myDeckStr, myDeckLoadedStr)
	}

}
