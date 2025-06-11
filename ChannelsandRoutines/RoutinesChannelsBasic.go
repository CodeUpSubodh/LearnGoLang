package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{"http://google.com", "http://facebook.com", "http://stackoverflow.com", "http://golang.org", "http://amazon.com"}
	c := make(chan string) //Creating Channel

	for _, link := range links {
		//checkLink(link) //There is delay at the time of execution because go is waiting for response and then proceeding to other execution process
		go checkLink(link, c) //Creates Go Routine and passing channel reference in function
	}

	// fmt.Println(<-c) //Reciving a value through channel even after this implementaion we will recive data from only one go routine
	// fmt.Println(<-c)

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLink(<-c, c) //Repeting Rotuines
	// }
	for l := range c {
		time.Sleep(5 * time.Second) // Sleep for 5 seconds
		go checkLink(l, c)          //Repeting Rotuines
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Might be down!")
		c <- "Might be Down" //Sending data through channel
	}
	fmt.Println(link, "is up!")
	// c <- "Its Up"
	c <- link
}
