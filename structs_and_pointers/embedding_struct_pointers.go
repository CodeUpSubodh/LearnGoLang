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

	// fmt.Printf("%+v", jim)
	jimPointer := &jim //Pointers in Go to Update a struct value
	// & is used to assgin memory address of a variable
	jimPointer.updateName("jimmy")
	// jim.updateName("jimmy") //This will also work but the function should have reciver type as pointer
	jim.print()
	// This case of creating a copy and functions can only modifies the copies and not the original values
	// without pointer is for integer, string, float and struct but in case of slices pointers are not
	// needed to change a value using a function

}

// Reciver funcation and function in structs

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	// * is used to assgin value for a memory address
	(*pointerToPerson).firstName = newFirstName //Pointers in Go to Update a struct value
}
