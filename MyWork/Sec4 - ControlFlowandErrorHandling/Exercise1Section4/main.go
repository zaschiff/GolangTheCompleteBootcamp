/*
*	Title:				Exercise1/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-24
*	Version:			1.0
*	Description:		First Program in section 4. SHows the use of the if statement and conditional checks
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. Using in if statemnt print the follwoing information:
	// if age is greater than 60 print Getting Older
	// if age is greater than 30 print Getting Wiser
	// if age is greater than 20 print Adulthood
	// if age si greater than 10 print Young blood
	// otherwise print Booting up
	age1 := 65
	if age1 > 60 {
		fmt.Println("Getting Older")
	} else if age1 > 30 {
		fmt.Println("Getting Wiser")
	} else if age1 > 20 {
		fmt.Println("Adulthood")
	} else if age1 > 10 {
		fmt.Println("Young Blood")
	} else {
		fmt.Println("Booting Up")
	}

	// 2. Simplify the following code:
	// isSphere, radius := true, 200
	// var big bool
	// if radius >= 50 {
	// 	if radius >= 100 {
	// 		if radius >= 200 {
	// 			big = true
	// 		}
	// 	}
	// }
	// if big != true {
	// 	fmt.Println("I don't know.")
	// } else if !(isSphere == false) {
	// 	fmt.Println("It's a big sphere.")
	// } else {
	// 	fmt.Println("I don't know.")
	// }
	isSphere2, radius2 := true, 100
	var big2 bool
	if radius2 >= 200 {
		big2 = true
	} else {
		big2 = false
	}

	if big2 && isSphere2 {
		fmt.Println("It's big sphere.")
	} else {
		fmt.Println("I don't know")
	}

	// 3. Print an output dependent on the number of arguments given
	//		at the commnd line
	if len(os.Args) == 2 {
		fmt.Printf("There is one: %s\n", os.Args[1])
	} else if len(os.Args) == 3 {
		fmt.Printf("There are two: \"%s %s\"\n", os.Args[1], os.Args[2])
	} else if len(os.Args) > 3 {
		fmt.Printf("There are %d arguments\n", len(os.Args))
	} else {
		fmt.Println("Give me Args!")
	}

	// 4. Detect if the command line input is a vowel or consonent
	if len(os.Args) < 2 || len(os.Args[1]) != 1 {
		fmt.Println("Please enter a letter")
		return
	}

	if strings.IndexAny(os.Args[1], "aeiou") != -1 {
		fmt.Printf("%q is a vowel", os.Args[1])
	} else if strings.IndexAny(os.Args[1], "wy") != -1 {
		fmt.Printf("%q is sometimes a vowel and sometimes not.", os.Args[1])
	} else {
		fmt.Printf("%q is a consonant", os.Args[1])
	}

}
