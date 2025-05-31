// Error Handling

package main

import (
	"errors"
	"fmt"
)

// divide function returns the result of a / b and an error if any
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	var num1, num2 float64

	fmt.Print("Enter numerator: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter denominator: ")
	fmt.Scanln(&num2)

	result, err := divide(num1, num2)
	if err != nil {
		// Error handling
		fmt.Println("Error:", err)
	} else {
		// No error, print result
		fmt.Printf("Result: %.2f\n", result)
	}
}
