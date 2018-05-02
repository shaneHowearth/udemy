package main

import (
	"os"
	"testing"
)

// The deck should have 52 cards
func TestNewDeck(t *testing.T) {
	nd := newDeck()

	if len(nd) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(nd))

	}

	// The first card should be the
	if nd[0] != "Ace of Spades" {
		t.Errorf("Expected last card to be Ace of Spades but got %v", nd[len(nd)-1])
	}

	// The last card should be the
	if nd[len(nd)-1] != "King of Diamonds" {
		t.Errorf("Expected last card to be King of Diamonds but got %v", nd[len(nd)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
