/*
*	Title:				Exercise7/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-17
*	Version:			1.0
*	Description:		This is the main.go file for the seventh exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
*
*						Please make sure to either comment out the line under problem 10 or compile the
*						program before running. if you use the go run command you can pass arguments but it
*						gets confusing with the number of command line values.
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. Print your age using Printf
	fmt.Printf("I'm %d years old\n", 27)

	// 2. Print your first name and last name using Printf and variables
	firstName := "Zach"
	lastName := "Schiff"
	fmt.Printf("My name is %s and my last name is %s\n", firstName, lastName)

	// 3. Use Printf to print the gollowing variable
	// tf := false
	tf := false
	fmt.Printf("These are %t claims\n", tf)

	// 4. Print the temperature with Printf but do not use %v
	// 		and do not give it a large amount of precision (trailing zeroes)
	temp4 := 29.5
	fmt.Printf("Temperature is %.1f degrees.\n", temp4)

	// 5. Print hellow world with double quotes.
	fmt.Println("\"hello world\"")

	// 6. Print the type and value of 3 using Printf
	fmt.Printf("Type of %v is %T.\n", 3, 3)

	// 7. Print the type and value of 3.14 using Printf
	fmt.Printf("Type of %v is %T.\n", 3.14, 3.14)

	// 8. Print the type and value of "hello" using Printf
	fmt.Printf("Type of %v is %T.\n", "hello", "hello")

	// 9. Print the type and value of true using Printf
	fmt.Printf("Type of %v is %T.\n", true, true)

	// 10. using Printf, get your first name and last name from the command line
	// and print them to the screen
	fmt.Printf("You name is %s, and your last name is %s.\n", os.Args[1], os.Args[2])
}
