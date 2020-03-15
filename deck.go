package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	deckToByte := []byte(d.toString())
	return ioutil.WriteFile(filename, deckToByte, 0666)
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Clubs", "Spades", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, fmt.Sprintf("%s of %s", values, suits))
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	byteDeck, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteDeck), ",")

	return deck(s)
}

func deal(d deck, numberOfCards int) (deck, deck) {
	return d[:numberOfCards], d[numberOfCards:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
