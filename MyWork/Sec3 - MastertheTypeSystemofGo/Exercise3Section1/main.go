/*
*	Title:				Exercise3/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-13
*	Version:			1.0
*	Description:		This is the main.go file for the third exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
 */

package main

import "fmt"

func main() {
	// 1. Using short variable declaration, declare and print 4 variables
	i, f, s, b1 := 314, 3.14, "hello", true
	fmt.Println("i:", i, "f:", f, "s:", s, "b:", b1)

	// 2. Using short declaration, declare and print two variables
	a1, b2 := 14, true
	fmt.Println(a1, b2)

	// 3. Using short declaration, declare and print the following variables:
	// a = 42
	// c = "good"
	a2, c := 42, "good"
	fmt.Println(a2, c)

	// 4. Using short declaration, declare and print a variable who's value is
	// 		the sum of the expression 27 + 3.5
	sum := 27 + 3.5
	fmt.Println(sum)

	// 5. Using short declaration, declare two boolean variables.
	//		Print the first variable and discard the second.
	b3, _ := true, true
	fmt.Println(b3)

	// 6. Using short variable declaration. Decalre two variables called age and yourAge.
	//		then declare a new variable called ration and redeclar age to 42
	age, yourAge := 16, 20
	age, ratio := 42, 3.14
	fmt.Println(age, yourAge, ratio)
}
