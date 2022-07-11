// Calling GitHub API
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Name     string `json:"name"`
	NumRepos int `json:"public_repos"`// from public_repos in JSON
}

// userInfo return information on github user
func userInfo(login string) (*User, error) {
	url := "https://api.github.com/users/" + login
	

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: can't call " + url)
		return nil, err
	}
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("error: No Body Found at " + url)
		return nil, err
	}

	// Decode request
	dec := json.NewDecoder(resp.Body)

	var user User
	if err := dec.Decode(&user); err != nil {
		log.Fatalf("error: can't decode - %s", err)
		return nil, err
	}

	return &user, nil
}

func main() {
	user, err := userInfo("tebeka")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Printf("%#v\n", user)
}
