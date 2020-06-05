/*
*	Title:				Exercise2Section6/wizardPrinter/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		using mutlidimensional arrays. Store several scientists'
*						first names, last name, and their nicknames.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	scientists := [6][3]string{
		{"Firstname", "Lastname", "Nickname"},
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "Apple"},
		{"Stephen", "Hawking", "Blackhole"},
		{"Marie", "Curie", "Radium"},
		{"Charles", "Darwin", "Fittest"},
	}

	for i, sci := range scientists {
		fmt.Printf("%-25s%-25s%s\n", sci[0], sci[1], sci[2])
		if i == 0 {
			fmt.Println(strings.Repeat("=", (len(sci[0]) + len(sci[1]) + len(sci[2]) + 50)))
		}
	}
}
