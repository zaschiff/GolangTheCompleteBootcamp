/*
*	Title:				Exercise10/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-17
*	Version:			1.0
*	Description:		This is the main.go file for the tenth exercise in section three of the course.
*						The code is for the exercise labeled Constants Exercise. To run each code
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
	"time"
)

func main() {
	// 1. Declare minutesInDay and daysInWeek constants and print the
	//		total of minutes in 2 weeks.
	const (
		minutesInDay1 = 24 * 60
		daysInWeek1   = 7
	)
	fmt.Printf("There are %d minutes in %d weeks.\n", (minutesInDay1 * (daysInWeek1 * 2)), 2)

	// 2. Create three constants hoursInDay, daysInWeek, hoursPerWeek.
	//		Use the previous two constants to create the third and print
	//		the total of hours in 5 weeks.
	const (
		hoursInDay2  = 24
		daysInWeek2  = 7
		hoursInWeek2 = daysInWeek2 * hoursInDay2
	)
	fmt.Printf("There are %d hours in %d weeks.\n", 5*hoursInWeek2, 5)

	// 3. Declare two constants one to the string Milky Way Galaxy,
	//		the other to the length of the first. Print them.
	const (
		home3   = "Milky Way Galaxy"
		length3 = len(home3)
	)
	fmt.Printf("There are %d characters inside %q\n", length3, home3)

	// 4. Fix the following code to print the value of Tau
	//		const pi, tau = 3.14159265358979323846264, pi * 2
	// the above code does not work as it is trying to assign tau before pi
	//		is initialized and defined.
	const (
		pi4  = 3.14159265358979323846264
		tau4 = pi4 * 2
	)
	fmt.Printf("Tau = %.15f\n", tau4)

	// 6. Fix the code
	// const (
	// 	width  int16 = 25
	// 	height int32 = width * 2
	// )

	// fmt.Printf("area = %d\n", width*height)
	// Another way to fix this is by removing the types in the constant declaration
	//		and just having them as typeless contants.
	const (
		width  int16 = 25
		height int32 = int32(width) * 2
	)

	fmt.Printf("area = %d\n", int32(width)*height)

	// 7. Fix the code and explain why it doesn't work
	// const later int = 10

	// hours, _ := time.ParseDuration("1h")

	// fmt.Printf("%s later...\n", hours*later)
	// it doesnt work as later is of type int and hours is a duration.
	//		To fix this without using conversion, we need to make later typeles
	const later = 10

	hours, _ := time.ParseDuration("1h")

	fmt.Printf("%s later...\n", hours*later)

	// 8. Using Iota, delcare three constants that increase in value.
	const (
		Nov = 11 - iota
		Oct = 11 - iota
		Sep = 11 - iota
	)
	fmt.Println(Sep, Oct, Nov)

	// 8. Using iota iniliaze three constants for the first three months.
	const (
		_   = iota
		jan = iota
		feb = iota
		mar = iota
	)
	fmt.Println(jan, feb, mar)

	// 9. Using iota, declare constants for the seasons where they begin.
	const (
		spring = (iota + 1) * 3
		summer
		fall
		winter
	)
	fmt.Println(spring, summer, fall, winter)
}
