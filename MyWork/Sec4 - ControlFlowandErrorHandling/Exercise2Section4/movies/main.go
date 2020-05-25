/*
*	Title:				Exercise2Section4\movies\main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-25
*	Version:			1.0
*	Description:		Simple program to check the age of a person and verify what movie ratings they can see.
*						It will take one command line arument and assume it is age and print the appropriate
*						rating group
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	// Program Usage
	usage string = `
		Usage:
		---------------------
		A small program to check what group rating a person can see based on an age 
		that is entered at the command line

		movies [age]

		[age]	age of the user that needs to be checked. Must be a number and 
				can not be blank or negative
		`
	// Ratings
	r    string = "R-Rated movies."
	pg13 string = "PG-13 movies."
	pg   string = "PG movies."

	// Message
	ratingMsg string = "This person can see "
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("\n%q Requires Age\n", args[0])
		fmt.Println(usage)
		return
	}

	age, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error: %q is not a number.\n", args[1])
		return
	}

	if age <= 0 {
		fmt.Printf("Wrong age: %d\n", age)
	} else if age > 17 {
		fmt.Println(ratingMsg + r)
	} else if age > 12 {
		fmt.Println(ratingMsg + pg13)
	} else {
		fmt.Println(ratingMsg + pg)
	}
}
