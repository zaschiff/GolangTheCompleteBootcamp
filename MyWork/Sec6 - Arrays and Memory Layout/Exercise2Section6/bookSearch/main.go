/*
*	Title:				Exercise2Section6/bookSearch/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		a simple example of a search engine that will
*						search for keywords in the book titles collected
*						inside an array.
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	books := [4]string{
		"kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}

	if len(os.Args) != 2 {
		fmt.Println("No Search Query. Program Exiting.")
		return
	}

	query := strings.ToLower(os.Args[1])
	count := 0

	fmt.Println("Search Results:")
	for i, b := range books {
		if strings.Contains(strings.ToLower(b), query) {
			count++
			fmt.Printf("\t+ %s\n", b)
		}

		if i == (len(books)-1) && count == 0 {
			fmt.Printf("\tWe do not have a book with %q in the title.\n", query)
		}
	}
}
