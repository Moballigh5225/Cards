package main





func main(){
	cards :=  newDeck()
	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	cards.toString()
	
	cards.saveToFile("myCards")
	(newDeckFromFile("my")).print()

	cards.shuffle()
	cards.print()
}



