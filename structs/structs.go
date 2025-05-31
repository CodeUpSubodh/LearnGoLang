// stuct in go means collection of related properties

package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// subodh := person{"Subodh", "Srivastava"} //Depends on assigning the value of the basis of var defined in struct
	subodh := person{firstName: "Subodh", lastName: "Srivastava"} //Same work but different syntax
	fmt.Println(subodh)

}
