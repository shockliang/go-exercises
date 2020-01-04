package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Arrange
	expected := 52

	// Act
	d := newDeck()

	// Assert
	if len(d) != expected {
		t.Errorf("Expected deck length of %v, but get %v", expected, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFrromFile(t *testing.T) {
	// Arrange
	filename := "_decktesting"
	os.Remove(filename)

	d := newDeck()

	// Act
	d.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	// Assert
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
