/*
*	Title:				Exercise2/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-09
*	Version:			1.0
*	Description:		This is the main.go file for the second exercise in section two of the course above.
*						To run each code separately, comment and uncomment accordingly. Each exercise
*						problem is labled by number.
 */

package main

func main() {
	// 1. Shy Semicolons - Observe how Go Fixes semicolons for you
	// The soltion below is correc, tbut with the functionality of Vs Code and go,
	// each time the file is saves, it will remove the semicolons, and place the statements
	// on their own line
	//fmt.Println("Hellow");fmt.Println("Gopher");fmt.Println("Bye")

	// 2. Naked Expression - Observe what happens when you use an expression wihtout a statement
	// This throws an error "'Hello' evaluated but not used"
	//"Hello"

	// 3. Operators Combine - try using operators to combine epxressions
	// prints out helloropher
	//fmt.Println("hello" + "gopher")

	// 4. Print Go Version - use a package from Go Standard Library to print the current version of
	//		Go on your system
	//fmt.Println(runtime.Version())

	// 5. Comment out - Leanr how to comment out.
	// the below code uses both single line and multiline comments
	//fmt.Println("Hello")
	/*
		fmt.Println("Gopher")
		fmt.Println("!")
	*/

}

// 6. Use the GoDoc - try Godoc yourself
// ran in terminal using the command go doc runtime NumCPU
