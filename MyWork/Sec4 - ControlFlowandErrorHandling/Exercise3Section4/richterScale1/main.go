/*
*	Title:				Exercise3Section4/richeterScale1/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-28
*	Version:			1.0
*	Description:		Gets the argument from the command line and
*						prints the appropraite richter scale message
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: No magnitude given. Please enter a magnitude")
		return
	}

	mag, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("%q is not a magnitude. Please enter a real magnitude.\n",
			os.Args[1])
		return
	}

	switch {
	case mag < 2.0:
		fmt.Printf("%.2f is a micro earthquake.\n", mag)
	case mag > 2.0 && mag < 3.0:
		fmt.Printf("%.2f is a very minor earthquake.\n", mag)
	case mag > 3.0 && mag < 4.0:
		fmt.Printf("%.2f is a minor earthquake.\n", mag)
	case mag > 4.0 && mag < 5.0:
		fmt.Printf("%.2f is a light earthquake.\n", mag)
	case mag > 5.0 && mag < 6.0:
		fmt.Printf("%.2f is a moderate earthquake.\n", mag)
	case mag > 6.0 && mag < 7.0:
		fmt.Printf("%.2f is a strong earthquake.\n", mag)
	case mag > 7.0 && mag < 8.0:
		fmt.Printf("%.2f is a major earthquake.\n", mag)
	case mag > 8.0 && mag < 10.0:
		fmt.Printf("%.2f is a great earthquake.\n", mag)
	default:
		fmt.Printf("%.2f is a massive earthquake.\n", mag)
	}
}
