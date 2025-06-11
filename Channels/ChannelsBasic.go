package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"}
	for _, link := range links {
		checkLink(link) //There is delay at the time of execution

	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
	}
	fmt.Println(link, "is up!")
}
