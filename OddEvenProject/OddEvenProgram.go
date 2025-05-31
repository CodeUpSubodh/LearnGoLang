package main

import "fmt"

type intslice []int

func main() {

	intslice := intslice{1, 58, 89, 24, 67, 61}
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
