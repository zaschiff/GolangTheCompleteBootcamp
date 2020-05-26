/*
*	Title:				Exercise2Section4/LeapYear/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-25
*	Version:			1.0
*	Description:		This program will check if the command line argument is  valid year and is a leap year.
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

	LeapYear:
	A simple program to check if the command line argument is 
	a valid year and if it is also a leap year
	
	main.go [value]
	
	[value]	the number that the program will check
	`
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: No year entered. Please enter a year.")
		fmt.Println(usage)
		return
	}

	if len(os.Args[1]) != 4 {
		fmt.Println("This doesn't seem like a year. Please enter a valid Year.")
		return
	}

	year, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("This isn't a valid number or year. Please enter a four digit year.")
		return
	}

	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Printf("The year %d is a leap year.\n", year)
	} else {
		fmt.Printf("The year %d is not a leap year.\n", year)
	}
}
