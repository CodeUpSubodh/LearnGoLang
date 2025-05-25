// Created to test go commands for multiple file execution {go run firstProgram.go multipleFileExecution.go}
package main

import "fmt"

func anotherFilefunction() string {
	fmt.Println("Hello from another file")
	return "Hello from another file"

}
