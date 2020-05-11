/*
*	Title:				Exercise1/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-11
*	Version:			1.0
*	Description:		This is the main.go file for the first exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
 */

package main

import "fmt"

func main() {
	// 1. Print the literals - Print a few values using the literals.
	fmt.Println(17, 12, 22, 28, 29)
	fmt.Println(3.2, 1., 2.5, -5.07)
	fmt.Println(true, false)
	fmt.Println("hello", "gopher")

	// 2. Print Hexes - Print Numbers in Hexadecimal
	fmt.Println(0x11)
	fmt.Println(0x22)
	fmt.Println(0x33)
	fmt.Println(0x44)
	fmt.Println(0x55)
}
