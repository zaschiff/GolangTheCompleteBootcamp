/*
*	Title:				PassMulti/Main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-24
*	Version:			1.0
*	Description:		Same as the program Passme, only this can take and validate two different users one at a time.
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
	accessOk   string = "Access granted for %q.\n"
	errPass    string = "Incorrect password for %q.\n"
	deniedUser string = "Invalid user. Ending program."

	user, user2 = "Jack", "Zach"
	pass, pass2 = "1888", "0220"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if (u != user) && (u != user2) {
		fmt.Println(deniedUser)
	} else if (u == user) && (p == pass) {
		fmt.Printf(accessOk, u)
	} else if (u == user2) && (p == pass2) {
		fmt.Printf(accessOk, u)
	} else {
		fmt.Printf(errPass, u)
	}

}
