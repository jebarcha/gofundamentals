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
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		//go checkLink(l, c)

		go func() {
			//time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()

	}

	for range links {
		<-c
	}

	// // for l := range c {
	// // 	go func() {
	// // 		time.Sleep(5 * time.Second)
	// // 		checkLink(l, c)
	// // 	}()
	// // }

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link //"Might be down I think"
		return
	}

	fmt.Println(link, "is up!")
	c <- link //"Yep it's up!"
}
