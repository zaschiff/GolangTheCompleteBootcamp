/*
*	Title:				FeeToMeters/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-25
*	Version:			1.0
*	Description:		Simple program to convert feet to meters based on command line input.
*						Continas error handling to accoutn for issue that may arrise.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const usage string = `
Usage:
---------------------
A small program to convert a given input from the command line assumed to be feet 
and prints the converted value in meters to the screen

FeetToMeters [feet]

[feet] 		value that will be converted.
`

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	feet, err := strconv.ParseFloat(args[1], 64)

	if err != nil {
		fmt.Printf("Error: The value entered is not a number.\nPlease enter a number.\n")
		fmt.Printf("Value entered: %q\n", args[1])
		fmt.Println("Here is some help with the program:")
		fmt.Println(usage)
		return
	}

	meters := feet * 0.3048

	fmt.Printf("%.2f feet is equal to %.2f meters.\n", feet, meters)
}
