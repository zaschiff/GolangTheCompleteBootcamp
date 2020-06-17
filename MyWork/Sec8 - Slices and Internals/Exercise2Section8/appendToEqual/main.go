/*
*	Title:				Exercise2Section8/appendToEqual/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-16
*	Version:			1.0
*	Description:		using pre-written code, append to a slcide to make it equal to another slice.
 */

package main

import "fmt"

func main() {
	// this is the pre-written code. It declares two slices
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	header = append(header, 'P', 'N', 'G')

	equal := false

	if len(png) == len(header) {
		equal = true
	}

	for _, letter := range png {
		for _, compare := range header {
			if letter == compare {
				equal = true
			} else {
				equal = false
			}
		}
	}

	fmt.Printf("Are they equal: %t.\n", equal)
}
