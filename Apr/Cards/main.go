package main

import "fmt"

func main() {
	cards := newCards()
	cards.Print()
	// hand, remainingCards := cards.deal(6)
	// hand.print()
	// remainingCards.print()
	// fmt.Printf("%+v\n", []byte("krtishna"))
	// fmt.Printf("%T %T\n", hand, remainingCards)
	// fmt.Printf("%T %T\n", hand.toString(), remainingCards.toString())
	// fmt.Printf("%v %v\n", hand.toString(), remainingCards.toString())
	// cards.saveToFile("myCards")
	// newDeck := newDeckFromFile("myCards1")
	// newDeckFromFile("myCards")
	// fmt.Printf("%+v\n", newDeck)
	// fmt.Printf("newDeck: %+v\n", newDeck)
	// fmt.Printf("%+v\n", "We are back to mainc")
	fmt.Println()
	cards.Shuffle()
	cards.Print()
}
