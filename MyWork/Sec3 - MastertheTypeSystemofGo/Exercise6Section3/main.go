/*
*	Title:				Exercise4/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-14
*	Version:			1.0
*	Description:		This is the main.go file for the sixth exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. print the count of command line arguments if the input is bilbo balbo bungo
	// count1 := len(os.Args) - 1
	// fmt.Printf("there are %d names.\n", count1)

	// 2. Print the path of the running program using os.Args
	// fmt.Println(os.Args[0])

	// 3. Print your name when it is passed as command line argument
	// fmt.Println(os.Args[1])

	// 4. Print out greetings to 3 people that are passed as command ine arguments.
	// 		store the information in variables
	// length := len(os.Args) - 1
	// name1 := os.Args[1]
	// name2 := os.Args[2]
	// name3 := os.Args[3]

	// fmt.Println("there are", length, "people!")
	// fmt.Println("Hello greate", name1+"!")
	// fmt.Println("Hello Greate", name2+"!")
	// fmt.Println("Hello Greate", name3+"!")
	// fmt.Println("Nice to meet you all.")

	// 5. Greet 5 people but do not use variables
	fmt.Println("there are", len(os.Args)-1, "people!")
	fmt.Println("Hello greate", os.Args[1]+"!")
	fmt.Println("Hello greate", os.Args[2]+"!")
	fmt.Println("Hello greate", os.Args[3]+"!")
	fmt.Println("Hello greate", os.Args[4]+"!")
	fmt.Println("Hello greate", os.Args[5]+"!")
	fmt.Println("Nice to meet you all.")
}
