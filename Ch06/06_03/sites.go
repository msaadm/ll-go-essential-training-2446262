// Get content type of sites (using channels)
package main

import (
	"fmt"
	"net/http"
)

func returnType(url string, output chan string) {
	resp, err := http.Get(url)
	if err != nil {
		output <- fmt.Sprintf("%s -> error: %s\n", url, err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	output <- fmt.Sprintf("%s -> %s", url, ctype)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	// Create response channel
	ch := make(chan string)
	
	for _, url := range urls {
		go returnType(url, ch)
	}
	// close(ch)

	// Wait using channel
	for range urls { // Run number of Urls times
		output := <-ch
		fmt.Println(output)
	}
}
