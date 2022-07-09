package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := strings.Fields(strings.ToLower(text));
	word_map := map[string]uint8{}

	for _, word := range words {
		fmt.Println("Processing : " + word)
		count, ok := word_map[word]
		if !ok {
			word_map[word] = 1
		} else {
			word_map[word] = count + 1
		}
	}

	fmt.Println(word_map)
}
