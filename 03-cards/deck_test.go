package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spdes, but got %v", d[0])
	}
	lastIndex := len(d) - 1
	if d[lastIndex] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[lastIndex])
	}
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expect 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")
}
