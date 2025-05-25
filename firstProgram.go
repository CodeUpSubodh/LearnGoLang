package main

import (
	"fmt"
)

func main() {
	//var declaration
	var card string = "Ace of spades"
	card2 := "Kind of spades"
	card3 := newCard()
	card4 := newSlice()

	// Custom Type
	card5 := newType()
	card6 := deck{"UNO Reverse", "UNO Skip"}

	card6.pritn()

	fmt.Println(anotherFilefunction())
	fmt.Println(card5)
	fmt.Println(card4)
	fmt.Println(card3)
	fmt.Println(card2)
	fmt.Println(card)
	fmt.Println("Hello World")
}

// funcation decalaration
func newCard() string {
	return "five of Diamaonds"
}

// arrays know as slices for go
func newSlice() []string {
	newSlice := []string{"New Card", "Old Card"}
	anotherNewSlice := append(newSlice, "New New Card")

	for i, card := range anotherNewSlice {
		fmt.Println(i, card)
	}
	return anotherNewSlice
}

type deck []string

// Custome Type Declaration
func newType() deck {
	card := deck{"king", "queen"}
	print(card)
	return card

}

// Receiver Function

func (d deck) pritn() deck {
	for i, card := range d {
		fmt.Println(i, card)
	}
	return d
}
