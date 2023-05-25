package main

import (
	"fmt"
	"strings"
)

// count how many times each word appears in a text

// to split text to works, use the "Fields" function from the "strings" package to convert all words to lowercase

func main() {
	text := `
	This is the sample text to use for my program. 
	My program will count the occurrence of 
	each word that is seen in the text.
	`
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	wordCt := map[string]int{}

	for _, word := range words {
		wordCt[word] = 0
		for _, wordCheck := range words {
			if word == wordCheck {
				wordCt[word]++
			}
		}
	}

	fmt.Printf("final count: %v\n", wordCt)
}
