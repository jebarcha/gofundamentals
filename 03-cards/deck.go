package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Dimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// (d deck) is called: "receiver"
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	//["red","yellow","blue"] -> "red,yellow,blue"
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and return a call to newDeck()
		//->Option 2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

//OLD: no need to do the NewSource and UnixNano thing anymore
// func (d deck) shuffle1() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)
// 	for i := range d {
// 		newPosition := r.Intn(len(d) - 1)
//			d[i], d[newPosition] = d[newPosition], d[i]
//		}
//	}

// with a receiver of a deck
func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
