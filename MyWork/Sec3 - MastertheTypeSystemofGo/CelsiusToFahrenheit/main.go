/*
*	Title:				CelsiusToFahrenheit/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-17
*	Version:			1.0
*	Description:		This is the main go file for a small exercise program called CelsiusToFahrenheit.
*						The program will take a command line arguement and assume it is in Celsius. It will
* 						convert the argument from a string to a float and perform the necessary calculations
*						to convert it to Fahrentheit. it will then print the result to the command line.
 */

// Main package to hold our program.
package main

// necessary imports.
import (
	"fmt"
	"os"
	"strconv"
)

// Main, entry point of the program.
func main() {
	// obtain the value we want to convert from the second index of the os.Args slice.
	arg := os.Args[1]

	// Convert the value from a string type to a float64. We could alos have chosen a float32
	// by passing 32 rather than 64. we discard the error value in this case as this program is
	// small and unlikely to currentl yeild an error at this time.
	celsius, _ := strconv.ParseFloat(arg, 64)

	//  Formula calculation to convert celsius to a fahrenheit value.
	fahrentheit := (celsius * 1.8) + 32

	// Using formatted text and escape sequences we can get the correct print statement.
	fmt.Printf("The temperature %.1f\u00B0C is %.1f\u00B0F", celsius, fahrentheit)
}
