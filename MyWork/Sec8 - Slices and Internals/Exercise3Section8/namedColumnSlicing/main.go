/*
*	Title:				Exercise3Section8/namedColumnSlicing/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-07-15
*	Version:			1.0
*	Description:		take in argument from the user and slice column accordingly
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	separator = ","
)

func main() {
	var (
		start int
		end   int
	)
	columnNames := strings.Split(strings.Split(data, "\n")[0], separator)
	cityData := strings.Split(data, "\n")[1:]

	if len(os.Args) > 3 {
		fmt.Println("Error: Too many arguments.")
		return
	}

	switch len(os.Args) {
	case 1:
		start = 0
		end = len(columnNames)
	case 2:
		for i, v := range columnNames {
			if string(v) == strings.Title(strings.ToLower(os.Args[1])) {
				start = i
			}
		}
		end = len(columnNames)
	case 3:
		for i, v := range columnNames {
			if string(v) == strings.Title(strings.ToLower(os.Args[1])) {
				start = i
			}
		}

		for i, v := range columnNames {
			if string(v) == strings.Title(strings.ToLower(os.Args[2])) {
				end = i
			}
		}
	}

	for i := start; i < end; i++ {
		fmt.Printf("%-15s", columnNames[i])
	}
	fmt.Println()
	fmt.Println()

	for _, v := range cityData {
		values := strings.Split(v, separator)
		for j := start; j < end; j++ {
			fmt.Printf("%-15s", values[j])
		}
		fmt.Println()
	}

}
