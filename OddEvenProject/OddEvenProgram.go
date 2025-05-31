package main

import "fmt"

type intslice []int

func main() {
	var length int

	fmt.Print("Lenght of slice ")
	fmt.Scan(&length)

	intslice := intslice{length}

	for i := 0; i < length; i++ {
		var value int
		fmt.Scan(&value)
		intslice = append(intslice, value)
	}

	fmt.Println(intslice)
	for _, value := range intslice {
		if value%2 == 0 {
			fmt.Println("Even", value, "Value")
			fmt.Printf("Even %v Value\n", value) //Same Work

		} else {
			fmt.Println("Odd Value", value)
		}
	}
}
