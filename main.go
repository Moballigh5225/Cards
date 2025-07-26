package main

import "fmt"





func main(){
	
	oddEven := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, number := range oddEven{
		if number % 2 == 0 {
			fmt.Println(number, "is even")
		}else{
			fmt.Println(number, "is odd")
		}
	}

	// cards :=  newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	// cards.toString()
	
	// cards.saveToFile("myCards")
	// (newDeckFromFile("my")).print()

	// cards.shuffle()
	// cards.print()
}

type oddEven []int

