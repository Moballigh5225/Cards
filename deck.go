package main

import (
	"fmt"
	"strings"
)


type deckCS []string




func newDeck() deckCS{
	cards := deckCS{}

	cardSuits := []string{"Spades","Hearts","Diamonds","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits{
		for _,value := range cardValues{
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}


// to print
func (d deckCS) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}


// to shuffle
func deal (d deckCS, handSize int) (deckCS, deckCS){
	return d[:handSize], d[handSize:]
}


// deck to byte string
func (d deckCS) toString() string {
	return	strings.Join([]string(d), ",")
}