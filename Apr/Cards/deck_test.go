package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newCards()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
	if d[0] != "Spades of Ace" {
		t.Errorf("excepted ace of spades, but got :( %v", d[0])
	}
	if d[len(d)-1] != "Clubs of Four" {
		t.Errorf("Expected Four of Clubes but got :( %v", d[len(d)-1])
	}
}

func TestSaveToAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newCards()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if (len(loadedDeck)) != 161 {
		t.Errorf("Expected deck length of 16, but got %d", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
