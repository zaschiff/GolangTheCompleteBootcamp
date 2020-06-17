/*
*	Title:				Exercise2Section8/appendFix/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-16
*	Version:			1.1
*	Description:		Fix the code given by the instructor
 */

package main

import "fmt"

func main() {
	// makethe slice's element type a string
	toppings := []string{"black olives", "green peppers"}

	// make pizza a slice rather than a three elemnt array
	// var pizza []string

	// move the elipsis to the end so it loops through the slice.
	// pizza = append(pizza, toppings...)
	// pizza = append(toppings, "onions")
	// toppings = append(pizza, "extra cheese")

	// BONUS: to simplify
	toppings = append(toppings, "onions", "extra cheese")
	fmt.Printf("toppings: %s\n", toppings)
}
