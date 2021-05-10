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

func newCards() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "Four"}
	for _, v := range cardSuits {
		for _, d := range cardValues {
			cards = append(cards, v+" of "+d)
		}
	}
	return cards
}

func (d deck) Print() {
	for i, v := range d {
		fmt.Printf("%v %+v\n", i, v)
	}
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Inside error block:\n%+v\n", err)
		os.Exit(1)
	}
	// fmt.Printf("%+v\n", string(data))
	strslice := strings.Split(string(data), ",")
	// fmt.Printf("%+v %T\n", strslice, strslice)
	return deck(strslice)
}
func (d deck) Shuffle() {
	// for i := range d {
	// 	newPosition := rand.Intn(len(d) - 1)
	// 	d[i], d[newPosition] = d[newPosition], d[i]
	// }
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
