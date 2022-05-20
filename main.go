package main

import (
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

// chan<- means a (send only) channel
func hitUrl(url string, c chan<- result) {
	res, err := http.Get(url)
	status := "OK"
	if err != nil || res.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- result{url: url, status: status}
}

func main() {

	// establish a channel
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
	}

	for _, url := range urls {
		go hitUrl(url, c)

	}

	// consuming data from the channel
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}

}
