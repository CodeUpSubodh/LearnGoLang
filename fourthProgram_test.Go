package main

import "testing"

// before writing any test case in go run command go mod init {directory name like LearnGoLoand}- This will create a file of go.mod after which you can execute go test
func TestDivide(t *testing.T) {
	a, b := 2.4, 5.2
	value := divide(a, b)

	if value == nil {
		t.Errorf("Float Value Expected")

	}

}
