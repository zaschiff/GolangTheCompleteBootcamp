/*
*	Title:				Exercise1/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-09
*	Version:			1.0
*	Description:		This is the main.go file for the first exercise in section two of the course above.
*						To run each code separately, comment and uncomment accordingly. Each exercise
*						problem is labled by number.
 */
package main

import "fmt"

func main() {
	// 1. Print your name and you best friend's name
	fmt.Println("Zach")
	fmt.Println("Alexa")

	// 4. Call Print instead of Println
	// fmt.Print("Zach")
	// fmt.Print("Alexa")

	// 5. Call Print or Println with mutiple values while separating them with a comma
	// fmt.Println("Zach", "Daniel", "Stephanie")
	// fmt.Println("Alexa", "Bo")

	// 6. Remove double quotes from string literal -> thorws and error "zach is undefined"
	// fmt.Println(Zach)
}

// 7. Move package and import statements to bottom of file -> throws an error "expected 'package', found 'func'"
// package main

// import "fmt"
