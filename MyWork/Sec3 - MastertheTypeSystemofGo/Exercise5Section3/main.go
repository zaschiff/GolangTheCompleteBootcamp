/*
*	Title:				Exercise4/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-14
*	Version:			1.0
*	Description:		This is the main.go file for the fourth exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
 */

package main

import "fmt"

func main() {
	// 1. Convert and fix the flllowing code
	// a,b := 10,5.5
	//fmt.Println(a + b)
	a1, b1 := 10, 5.5
	fmt.Println(float64(a1) + b1)

	// 2. Convert and fix
	// a, b := 10, 5.5
	// a = b
	// fmt.Println(a + b)
	a2, b2 := 10, 5.5
	a2 = int(b2)
	fmt.Println(float64(a2) + b2)

	// 3. Convert and fix the following code
	// fmt.Println(int(5.5))
	fmt.Println(5.5)

	// 4. Convert and fix the following code
	// age := 2
	// fmt.Println(int(7.5) + int(age))
	age := 2
	fmt.Println(7.5 + float64(age))

	// 5. Convert and fix
	// min := int8(127)
	// max := int16(1000)
	// fmt.Println(int8(max) + min)
	min := int8(127)
	max := int16(1000)
	fmt.Println(max + int16(min))
}
