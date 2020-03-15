package main

import (
	"os"
	"testing"
)

var deckLenAssert int = 28

func TestNewDeck(t *testing.T) {
	firstCardOnTheDeck := "Ace of Clubs"
	lastCardOnTheDeck := "Seven of Diamonds"
	d := newDeck()

	if len(d) != deckLenAssert {
		t.Errorf("Deck length should be %v, but got %v", deckLenAssert, len(d))
	}

	if d[0] != firstCardOnTheDeck {
		t.Errorf("First card on the deck must be %s, but got %s", firstCardOnTheDeck, d[0])
	}

	if d[len(d)-1] != lastCardOnTheDeck {
		t.Errorf("Last card on the deck must be %s, but got %s", lastCardOnTheDeck, d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	f := "_decktesting"
	os.Remove(f)

	d := newDeck()
	d.saveToFile(f)

	df := newDeckFromFile(f)

	if len(df) != deckLenAssert {
		t.Errorf("Deck from file length should be %v, but got %v", deckLenAssert, len(df))
	}

	os.Remove(f)
}

func TestDeal(t *testing.T) {
	cardsToDeal := 5
	d := newDeck()
	hand, d := deal(d, cardsToDeal)

	if len(d) != deckLenAssert-cardsToDeal {
		t.Errorf("Deck length after deal should be %v, but got %v", deckLenAssert-cardsToDeal, len(d))
	}

	if len(hand) != cardsToDeal {
		t.Errorf("Hand length should be %v, but got %v", cardsToDeal, len(hand))
	}
}

func TestToString(t *testing.T) {
	toStringReturn := "Ace of Clubs,Two of Clubs"
	cardsToDeal := 2
	d := newDeck()
	hand, _ := deal(d, cardsToDeal)

	if hand.toString() != toStringReturn {
		t.Errorf("toString method must return %s, but got %s", toStringReturn, hand.toString())
	}
}
