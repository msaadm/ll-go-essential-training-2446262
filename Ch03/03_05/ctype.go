// Writing a function that return Content-Type header
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	ctype, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}

// contentType will return the value of Content-Type header returned by making an
// HTTP GET request to url
func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "Incorrect Error or Site is Down!", err
	}
	defer resp.Body.Close()
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		return "Some Error", err2
	}
	return http.DetectContentType(body), nil
}
