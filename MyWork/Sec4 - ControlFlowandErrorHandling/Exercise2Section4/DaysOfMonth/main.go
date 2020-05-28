/*
*	Title:				Exercise2/DaysOfMonth/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-26
*	Version:			1.0
*	Description:		This program will take in a command line argument, check it for a month name, and
*						print out the number of days in that month.
 */

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	usage string = `
	Usage:
	---------------------

	DaysOfMonth:
	A simple program to check for a month name from command line input and
	print out the numebr of days in that month
	
	main.go [value]
	
	[value]	the month the program will chekc for to determine how many days are in the month
	`
	noMonth string = "Error: No month given. Please enter a month name."
)

func main() {
	var monthNames = []string{"January", "February", "March", "April", "May", "June", "July",
		"August", "September", "October", "November", "December"}
	var monthsWith31 = []string{"January", "March", "May", "July", "August", "October", "December"}
	var monthsWith30 = []string{"April", "June", "September", "November"}
	if len(os.Args) != 2 {
		fmt.Println(noMonth)
		fmt.Println(usage)
		return
	}

	if !contains(monthNames, strings.Title(strings.ToLower(os.Args[1]))) {
		fmt.Printf("%q is not a month.\n", os.Args[1])
	} else if contains(monthsWith31, strings.Title(strings.ToLower(os.Args[1]))) {
		fmt.Printf("%q has 31 days.\n", os.Args[1])
	} else if contains(monthsWith30, strings.Title(strings.ToLower(os.Args[1]))) {
		fmt.Printf("%q has 30 days.\n", os.Args[1])
	} else {
		var year int = time.Now().Year()
		if (year%400 == 0) || (year%4 == 0 && year%100 == 0) {
			fmt.Printf("%q has 29 days.\n", os.Args[1])
		} else {
			fmt.Printf("%q has 28 days.\n", os.Args[1])
		}
	}

}

func contains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
