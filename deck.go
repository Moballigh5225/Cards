package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
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



func newDeckFromFile(filename string) deckCS {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("ErrorBabu:",err)
		os.Exit(1)
	}
	s :=strings.Split(string(bs), ",")
	return deckCS(s)
}	

// my random approch
func randomSeedGen() int64{
	return time.Now().UnixNano() + int64(time.Now().Second()*42)
}
/*
// but this will effect the global rand
func (d deckCS) shuffle() {
	seed := randomSeedGen()
	rand.Seed(seed)

	for i := range d {
		newPosition := rand.Intn(len(d))
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
*/


func (d deckCS) shuffle(){
	source := rand.NewSource(randomSeedGen())
	r := rand.New(source)


	
	for i := range d{
		newPosition := r.Intn(len(d) - 1)
		d[i] , d[newPosition] = d[newPosition],d[i]
	}
}
