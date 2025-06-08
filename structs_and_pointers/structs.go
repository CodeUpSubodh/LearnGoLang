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
	var codeupsubodh person
	//Go assign default for strings its "" int- 0 float - 0 and bool - false
	fmt.Println(codeupsubodh) //empty struct
	fmt.Println(subodh)       //{Subodh Srivastava}
	fmt.Printf("%+v", subodh) //{firstName:Subodh lastName:Srivastava}%

}

// If attribute in struct has a value of type interface then it means that that atrribute's values shouls satify the functions
