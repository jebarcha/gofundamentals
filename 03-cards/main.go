package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
	//cards.print()

	//cards := newDeck()
	//fmt.Println(cards.toString())
	//greeting := "Hi there!"
	//fmt.Println([]byte(greeting))
	//cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards1")
	// cards.print()

	cards2 := newDeck()
	cards2.shuffle()
	cards2.print()

}

//func newCard() string {
//	return "Five of Diamonds"
//}

//take our deck type and turn it into a slice of string
//and join it all together in a one single string
//then turn it in to a byte slice: []byte
