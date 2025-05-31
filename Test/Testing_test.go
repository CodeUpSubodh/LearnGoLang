package main

import (
	"fmt"
	"testing"
)

// before writing any test case in go run command go mod init {directory name like LearnGoLang}- This will create a file of go.mod after which you can execute go test
func TestDivide(t *testing.T) {
	a, b := 2.4, 5.2 //Success Case
	// a, b := 0.0, 0.0 //Failure Case
	value, err := divide(a, b)
	fmt.Println(value)
	if err != nil {
		t.Errorf("0 cannot be the denominator %v", err) //to make use of dynamic value we use %v in go

	}

	if value != a/b {
		t.Errorf("Getting Invalid Values") // These if are acting as assert but in go it does not know how any Testcases are passed or faield

	}

}
