/*
*	Title:				Exercise3Section4/richterScale2/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-28
*	Version:			1.0
*	Description:		takes in a command line argument and prints out the
*						corresponding richter scale number
 */

package main

import (
	"fmt"
	"os"
)

const errNoMag = `Error: No Magnitude given.
Please give me a mgnitude in human terms.`

func main() {
	if len(os.Args) == 1 {
		fmt.Println(errNoMag)
		return
	}

	var desc string

	if len(os.Args) == 2 {
		desc = os.Args[1]
	} else {
		desc = os.Args[1] + " " + os.Args[2]
	}

	var mag string

	switch desc {
	case "micro":
		mag = "Less tha  2.0"
	case "very minor":
		mag = "2.0 - 2.9"
	case "minor":
		mag = "3 - 3.9"
	case "light":
		mag = "4 - 4.9"
	case "moderate":
		mag = "5 - 5.9"
	case "strong":
		mag = "6 - 6.9"
	case "major":
		mag = "7 - 7.9"
	case "great":
		mag = "8 - 10.0"
	case "massive":
		mag = "10.0 or more"
	default:
		mag = "not a magnitude in human terms"
	}

	fmt.Printf("%s's richter scale is %s.\n", desc, mag)
}
