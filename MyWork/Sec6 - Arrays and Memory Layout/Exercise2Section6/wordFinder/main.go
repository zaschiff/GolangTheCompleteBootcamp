/*
*	Title:				Exercise2Section6/wordFinder/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-05
*	Version:			1.0
*	Description:		Using arrays and other tools search a corpus for key
*						words passed at the command line by the user.
 */
package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again since the beginning this was very important"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give me a word to search.")
		return
	}

	filter := [...]string{
		"and",
		"or",
		"was",
		"the",
		"since",
		"very",
	}

	query := os.Args[1:]
	words := strings.Fields(corpus)

queries:
	for _, q := range query {

		for _, f := range filter {
			if strings.ToLower(q) == strings.ToLower(f) {
				continue queries
			}
		}

		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}
