/*
*	Title:				Exercise1Section8/fixTheCode/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-12
*	Version:			1.0
*	Description:		the code below will be th3e fixed variation.
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	// -----------------------------------
	books := []string{
		"Stay Golden",
		"Fire",
		"Kafka's Revenge",
	}

	sort.Strings(books)

	// -----------------------------------
	// this time, do not change the nums array to a slice
	nums := []int{5, 1, 7, 3, 8, 2, 6, 9}

	// use the slicing expression to change the nums array to a slice below
	sort.Ints(nums)

	// -----------------------------------
	// Here: Use the strings.Join function to join the names
	//       (see the expected output)
	fmt.Printf("%q\n", names)

	fmt.Printf("%q\n", books)
	fmt.Printf("%d\n", nums)
}
