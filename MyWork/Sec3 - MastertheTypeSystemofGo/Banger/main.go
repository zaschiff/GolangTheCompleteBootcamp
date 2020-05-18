/*
*	Title:				Banger/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-17
*	Version:			1.0
*	Description:		This is the main go file for a small exercise program called Banger. Banger will
*							take a command line argument that is a string, count the number of characters
*							in the string and add that number of exclamation points to the beginning and
*							the end of the string.
*
 */

// Main package for the program
package main

// Necessary import statements
import (
	"fmt"
	"os"
	"strings"
)

// Main Function, the entry point of the program
func main() {
	// get the string from the command line
	msg := os.Args[1]

	// get the length of the string
	l := len(msg)

	// create the string of the exclamation points
	bang := strings.Repeat("!", l)

	// concatenate the strings with the message and convert it to upper case
	s := bang + msg + bang
	s = strings.ToUpper(s)

	// print the result.
	fmt.Println(s)
}
