/*
*	Title:				Exercise2Section4/OddOrEven/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-25
*	Version:			1.0
*	Description:		This program will take a single command line argument and check whether it is even
*						or odd and if it is divisible by 8.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	usage string = `
	Usage:
	---------------------

	OddOrEven:
	A simple program to check if the command line argument is 
	even or odd and if it is divisible by 8.
	
	main.go [value]
	
	[value]	the number that the program will check
	`
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("\nError: No argument given. Please enter a number.")
		fmt.Println(usage)
		return
	}

	val, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error: %q is not a number.\n", args[1])
		return
	}

	if (val%2 == 0) && (val%8 == 0) {
		fmt.Printf("The value %d is even and divisible by 8.\n", val)
	} else if val%2 == 0 {
		fmt.Printf("The value %d is even.\n", val)
	} else {
		fmt.Printf("The value %d is odd.\n", val)
	}

}
