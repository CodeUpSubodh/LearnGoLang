package main

import (
	"fmt"
	"strings"
)

func main() {
	//var declaration
	var card string = "Ace of spades"
	card2 := "Kind of spades"
	card3 := newCard()
	card4 := newSlice()

	// Custom Type
	card5 := newType()

	deck := deck{"UNO Reverse", "UNO Skip"}
	deck.print()
	first_hand, second_hand := deck.deal(1)

	fruitSlice := newSliceTut()

	//Byte Conversion
	byteconversion()

	fmt.Println(first_hand)
	fmt.Println(second_hand)
	fmt.Println(fruitSlice)
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

func newSliceTut() []string {
	fruits := []string{"Apple", "Banana", "Grapes"}
	fmt.Println(fruits[0])
	fmt.Println(fruits[0:2]) //This means from 0 to excluding 2
	fmt.Println(fruits[:2])
	fmt.Println(fruits[2:])
	return fruits
}

type deck []string

// Custome Type Declaration
func newType() deck {
	card := deck{"king", "queen"}
	print(card)
	return card

}

// Receiver Function d
// This Working is same like how in OOPs related we create a class and use keywords like self and this to create the funcation that are related to that class

func (d deck) print() deck {
	for i, card := range d {
		fmt.Println(i, card)
	}
	return d
}

// Multiple Return Values func declaration

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Byte Slice
func byteconversion() {
	byte_value := []byte("Hi There")
	fmt.Println(byte_value)
	//Customer Type Conversion
	deck := deck{"UNO Reverse", "UNO Skip"}
	deckToSlice := []string(deck)
	fmt.Println(deckToSlice)
	sliceToString := strings.Join(deckToSlice, ",")
	fmt.Println((sliceToString))

}
