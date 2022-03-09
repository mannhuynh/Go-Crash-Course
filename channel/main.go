package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"http://amazon.com",
		"http://google.com",
		"http://golang.org",
		"http://facebook.com",
		"http://twitter.com",
	}
	c := make(chan string)

	for _, link := range links {
		go getLink(link, c)
	}

	for l := range c {
		go getLink(l, c)
	}
	// The above for loop is exactly same as below:
	// for {
	// 	go getLink(<-c, c)
	// }

}

// getLink will send links to channel c
func getLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is DOWN!")
		c <- link
		return
	}

	fmt.Println(link, "is UP")
	c <- link
}
