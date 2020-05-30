/*
*	Title:				multiplicationTable/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-29
*	Version:			1.0
*	Description:		A small coding challenge to create a dynamic
*						multiplication table based ona size passed in by the
*						user at the command line.
*
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// check for user input
	if len(os.Args) != 2 {
		fmt.Println("Give me a size for a multiplication table.")
		return
	}

	// check if user input is valid input.
	max, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%q is not a number.\n", os.Args[1])
		return
	}
	if max < 0 {
		fmt.Println("Wrong size.")
		return
	}

	// Print the header
	fmt.Printf("%5s", "X")
	for i := 0; i < max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// build the table
	for i := 0; i < max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j < max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
