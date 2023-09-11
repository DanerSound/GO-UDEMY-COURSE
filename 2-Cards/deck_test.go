package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	fmt.Println(" Sto Testanando 1")
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestCheckFirst(t *testing.T) {
	fmt.Println(" Sto Testanando 2")
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected ace of spades, but got %v", d[0])
	}
}

func TestCheckLast(t *testing.T) {
	fmt.Println(" Sto Testanando 3")
	d := newDeck()
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	fmt.Println(" Sto Testanando 4")
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
