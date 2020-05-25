/*
*	Title:				Passme/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-24
*	Version:			1.0
*	Description:		This is a simple program designed to take in a username and password for the command
*						prompt and verify if it is a valid user with a valid password. This is only a simple
*						program and is by no means secure. It will also only validate a single user.
 */

package main

import (
	"fmt"
	"os"
)

const (
	usage string = `Usage:
	go run main.go [username] [password]
	
	[username] - username to check for validation.
	[password] - password to check for validation
	
	please enter a username and password to check for validation.`
	deniedUser string = "Invalid user. Ending program."
	errPass    string = "Incorrect password for %q.\n"
	accessOk   string = "Access granted for %q.\n"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(usage)
		return
	}

	if os.Args[1] != "Jack" {
		fmt.Println(deniedUser)
	} else if os.Args[2] != "1888" {
		fmt.Printf(errPass, os.Args[1])
	} else {
		fmt.Printf(accessOk, os.Args[1])
	}
}
