/*
*	Title:				Exercise1Section6/assignArray/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		using array assignment, create a books array, use the
*						books array to populate two other arrays. Once they
*						are populated, use string methods to convert them to
*						all upper case and all lower case.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	books := [3]string{"Kafka's Revenge", "Stay Golden", "Everythingship"}
	upper := books
	lower := books

	for i, v := range upper {
		upper[i] = strings.ToUpper(v)
	}

	for i, v := range lower {
		lower[i] = strings.ToLower(v)
	}

	fmt.Printf("books: %q\n", books)
	fmt.Printf("upper: %q\n", upper)
	fmt.Printf("lower: %q\n", lower)
}
