/*
*	Title:				Exercise2Section5/wordSearchNoCase/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-02
*	Version:			1.0
*	Description:		Takes the word search program and allows the user to enter in any search
*						word regardless of case and get their results back.
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if strings.ToLower(q) == strings.ToLower(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				break queries
			}
		}
	}
}
