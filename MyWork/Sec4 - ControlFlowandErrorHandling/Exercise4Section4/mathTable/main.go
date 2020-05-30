/*
*	Title:				Exercise4Section4/mathTable/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-30
*	Version:			1.0
*	Description:		Utilizing a for loop, get two command line arguments. These are arguments
*						are 1) a mathmatical operation, 2) size of the table. Parse the informaiton
* 						and handle any errors that could occur
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// check the number of arguments entered at the command line
	if len(os.Args) != 3 {
		fmt.Println("Give me an operation followed by size.")
		return
	}

	// validate the operation
	op := os.Args[1]
	switch op {
	case "*":
		fallthrough
	case "+":
		fallthrough
	case "/":
		fallthrough
	case "-":
		fallthrough
	case "%":
		break
	default:
		fmt.Printf("%q is not an operation.\n", op)
		return
	}

	// validate the size of the table
	max, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("%q is not a number.\n", os.Args[2])
		return
	}
	if max < 0 {
		fmt.Println("Wrong size.")
		return
	}

	// create the header
	fmt.Printf("%5s", op)
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	// build the table based on the operation.
	switch op {
	case "*":
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", i*j)
			}
			fmt.Println()
		}
	case "+":
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", i+j)
			}
			fmt.Println()
		}
	case "/":
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= max; j++ {
				if j == 0 {
					fmt.Printf("%5d", 0)
				} else {
					fmt.Printf("%5d", i/j)
				}
			}
			fmt.Println()
		}
	case "-":
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", i-j)
			}
			fmt.Println()
		}
	case "%":
		for i := 0; i <= max; i++ {
			fmt.Printf("%5d", i)

			for j := 0; j <= max; j++ {
				fmt.Printf("%5d", i%j)
			}
			fmt.Println()
		}
	}
}
