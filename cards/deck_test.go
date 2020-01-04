package main

import "testing"

func TestNewDeck(t *testing.T) {
	// Arrange
	expected := 52
	d := newDeck()

	// Act

	// Assert
	if len(d) != expected {
		t.Errorf("Expected deck length of %v, but get %v", expected, len(d))
	}
}
