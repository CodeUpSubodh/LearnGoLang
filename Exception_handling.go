package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Define the file name and content
	fileName := "data.txt"
	content := "Hello, this is a file written by a Go program!"

	// ---------- Write to file ----------
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Data written to", fileName)

	// ---------- Read from file ----------
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}
	fmt.Println("Data read from file:")
	fmt.Println(string(data))
}
