package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	// Without go routines/channels
	//
	// One at a time... slow!
	for _, link := range links {
		checkLink(link)
	}

	// Let's do it concurrently with channels
	c := make(chan string)
	for _, link := range links {
		// Generate independent processes (linked through channel)
		go checkLinkWithChannel(link, c)
	}
	// We have to wait for all routines to send over the information
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

	// Continuous checker -- comment out to see how it works
	for _, link := range links {
		go checkLinkWithChannelReturnLink(link, c)
	}
	// for {
	// 	go checkLinkWithChannelReturnLink(<-c, c)
	// }
	//
	// Alternative syntax for previous loop with
	// time management (1secs between requests for each link)
	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLinkWithChannelReturnLink(link, c)
	// 	}(l)
	// }

	// With go routines alone.. we will see no output since
	// child processes are executed unlinked to the main process...
	// This will cause our main program to finish earlier - before
	// child process completion
	for _, link := range links {
		go checkLink(link)
		// That is why we need channels!
	}

}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("Link %v might be down\n", link)
	} else {
		fmt.Printf("Link %v is up\n", link)
	}
}

func checkLinkWithChannel(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "(Using channels) Link " + link + " might be down"
	} else {
		c <- "(Using channels) Link " + link + " is up"
	}
}

func checkLinkWithChannelReturnLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Printf("(Using channels) Link %v might be down\n", link)
		c <- link
	} else {
		fmt.Printf("(Using channels) Link %v is up\n", link)
		c <- link
	}
}
