/*
*	Title:				Exercise3Section4/stringManipulation/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-29
*	Version:			1.0
*	Description:		Using command line arguments, manipulate a string by getting the command and message
*						from the command line arguments
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	usage string = `
		Usage:
		---------------------

		go run main.go /go build main.go

		[command][message]
		
		[command] - desired manipulation to take place.
			Available commands are: lower, upper, and title
		[message] - the desired string to be changed.
	`
	errNoCmdMsg   string = "Error: No Command or Message."
	errMissingOne string = "Error: Either the Command or Message is missing."
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("\n", "\t"+errNoCmdMsg)
		fmt.Println(usage)
	}

	if len(os.Args) == 2 {
		fmt.Println("\n", "\t"+errMissingOne+"\n")
	}

	cmd := strings.ToLower(os.Args[1])
	words := os.Args[2:]
	var msg string

	for _, w := range words {
		msg = msg + " " + w
	}

	switch cmd {
	case "lower":
		msg = strings.ToLower(msg)
	case "upper":
		msg = strings.ToUpper(msg)
	case "title":
		msg = strings.Title(msg)
	default:
		fmt.Printf("Unknown Command: %q\n", cmd)
		return
	}

	fmt.Println(msg)
}
