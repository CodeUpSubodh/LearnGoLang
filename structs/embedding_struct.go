// stuct in go means collection of related properties

package main

import "fmt"

type contactInfo struct { //Embedding Structs
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// contactInfo  same working but the assigning key/variable will be also contactInfo
	contact contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 11005,
		},
	}

	fmt.Printf("%+v", jim)

}
