package main

import (
	"fmt"
	"os"
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

// save to file using new package os in place of deprecated package io/util
func (d deckCS) saveToFile(filename string) error{
		return os.WriteFile(filename, []byte(d.toString()), 0666)	
}