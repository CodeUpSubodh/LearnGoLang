package main

import "fmt"

type bot interface { //Interface Tyep - For which we can not create value directly
	getGreeting() string
}
type englishBot struct{} //Concrete Type - For which we can create value directly
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (englishBot) getGreeting() string { // If we are not using reciver value we can write englishBot in place of eb englishBot
	//Custom Logic
	return "Hello There"
}

func (spanishBot) getGreeting() string {
	//Custom Logic
	return "Holla!"
}
