/*
*	Title:				Exercise2/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-12
*	Version:			1.0
*	Description:		This is the main.go file for the second exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
 */

package main

import "fmt"

// 9. Declare a pakcage level variable and observe what happens when it is unused
var check bool

func main() {
	// 1. Declare an int - declare and print a int variable with name of height
	var height int
	fmt.Println(height)

	// 2. Declare a bool - declare and print a boolean variable with a name of isOn
	var isOn bool
	fmt.Println(isOn)

	// 3. Declare a float - declare and print a float variable with a name of brightness
	var brightness float64
	fmt.Println(brightness)

	// 4. Declare a string - declare and print a string variable
	// instructions ask use to use the following code eventhough we have not covered the printf function
	// %T prints the type of of the value
	// %q prints an empty string
	var s string
	fmt.Printf("s (%T): %q\n", s, s)

	// 5. Declare the following variables and observe the output
	// 3speed
	// !speed
	// spe?ed
	// var
	// func
	// package
	/*
		var 3speed int
		var !speed int
		var spe?ed int
		var var int
		var func int
		var package int
		fmt.Println(3speed)
		fmt.Println(!speed)
		fmt.Println(spe?ed)
		fmt.Println(var)
		fmt.Println(func)
		fmt.Println(package)
	*/
	// the above code throws variuos errors from naming convention
	//		issues to trying to ue keywords as variable names.

	// 6. Declare wwith Bits- declare variables using the follwoing types
	// int
	// int8
	// int16
	// int32
	// int64
	// float32
	// float64
	// complex64
	// complex128
	// bool
	// string
	// rune
	// byte
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var f32 float32
	var f64 float64
	var c64 complex64
	var c128 complex128
	var bo bool
	var s2 string
	var r rune
	var by byte
	fmt.Println(i, i8, i16, i32, i64, f32, f64, c64, c128, bo, s2, r, by)

	// 7. Delcare multiple - using mutiple variable declaration
	// declare a variable named active of type bool and a variable named delta of type int
	var (
		active bool
		delta  int
	)
	fmt.Println(active, delta)

	// 8. Declare mutiple variables of the same type. They should be named firstname and lastname
	var firstName, lastName string
	fmt.Println(firstName, lastName)

	// 9. Declare an unused varible named isLiquid using the blank identifier
	var isLiquid bool
	_ = isLiquid

	// 11. Attempt to use a varible before delaring it and observe what happnes
	/*
		fmt.Println(beforeDec)
		var beforeDec bool
	*/
	// the above code throws an arguement of being undefined.
	//		You can use a variable before declaring it.
}
