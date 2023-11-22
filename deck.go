package main

import (
	"fmt"
	"strings"
)

type deck []string 

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits{ // We use := inside the for loop because this variables doesent exist yet, different of C and C++ that whe would need declare empty variable before
		for _, value := range cardValues{
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string { // declaration functin, atribuitin function property to deck, function name, return type of function 
	return strings.Join([]string(d), ",")
}
