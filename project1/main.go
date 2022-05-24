package main // executable type package 

import (

//    "fmt"
//      "ioutil"
)    

func main() {

//	var card string = "Ace of Spades"
//	cards := newDeck()
//	hand, remainingCards := deal(cards, 5)
//
//	hand.print()
////	remainingCards.print()
//	greeting := "Hi there!"
//	fmt.Println([]byte(greeting))

    cards := newDeck()
    cards.saveToFile("my_cards")
}



