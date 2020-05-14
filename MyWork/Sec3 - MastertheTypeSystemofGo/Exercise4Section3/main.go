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

import (
	"fmt"
	"path"
)

func main() {
	// 1.  Change the variable color form green to blue
	color1 := "green"
	fmt.Println(color1)
	color1 = "blue"
	fmt.Println(color1)

	// 2. Assign the variable color to dark green using the previous value held by the variable and concatenation
	color2 := "green"
	fmt.Println(color2)
	color2 = "dark " + color2
	fmt.Println(color2)

	// 3. Assign the variable n using an expression for the value
	n1 := 0.
	fmt.Println(n1)
	n1 = 3.14 * 2
	fmt.Println(n1)

	// 4. Find the rectangle's perimeter when the width is 5 and height is 6
	var (
		perimeter     int
		width, height = 5, 6
	)
	perimeter = (2 * width) + (2 * height)
	fmt.Println(perimeter)

	// 5. using multi assign, assign "go" to the variable lang at the same time as assigning 2 to the variable version
	var (
		lang    string
		version int
	)
	lang, version = "go", 2
	fmt.Println(lang, "version", version)

	// 6. assign the correct values to the correct variables to print the following statement
	// Air is good on Mars
	// It's true
	// It is 19.5 degrees
	var (
		planet string
		isTrue bool
		temp   float64
	)

	planet, isTrue, temp = "Mars", true, 19.5
	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("it is", temp, "degrees")

	// 7. Use the function multi and short declaration to print out only the second result from the function.
	_, b1 := multi()
	fmt.Println(b1)

	// 8. change the values of the two variables at the same time.
	color3, color4 := "red", "blue"
	fmt.Println(color3, color4)
	color3, color4 = "orange", "green"
	fmt.Println(color3, color4)

	// 9. Complete a full swap with the two variables
	red, blue := "red", "blue"
	fmt.Println(red, blue)
	red, blue = blue, red
	fmt.Println(red, blue)

	// 10. Print only the directory using path.Split
	dir, _ := path.Split("secret/filte.txt")
	fmt.Println(dir)
}

func multi() (int, int) {
	return 5, 4
}
