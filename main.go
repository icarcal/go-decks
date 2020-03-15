package main

func main() {
	deck := newDeck()
	deck.shuffle()

	handPlayer1, deck := deal(deck, 5)
	handPlayer2, deck := deal(deck, 5)
	handPlayer3, deck := deal(deck, 5)

	handPlayer1.print()
	handPlayer2.print()
	handPlayer3.print()

	deck.print()
}
