/*
*	title:				Exercise4Section8/fixBackArray/main.go
*	course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	program author:		zach schiff
*	date:				2020-07-22
*	version:			1.0
*	description:		Fix the code so that that backing array is not modified when changing elements in
*						the original array
 */

package main

import "fmt"

func main() {
	// DON'T TOUCH THE FOLLOWING CODE
	nums := []int{56, 89, 15, 25, 30, 50}

	// ----------------------------------------
	// WRITE CODE HERE SO THAT:
	//
	// Ensure that nums slice never changes even though
	// the mine slice changes.
	mine := []int{}
	mine = append(mine, nums[:3]...)
	// ----------------------------------------

	// DON'T TOUCH THE FOLLOWING CODE
	//
	// This code changes the elements of the nums
	// slice.
	//
	mine[0], mine[1], mine[2] = -50, -100, -150

	fmt.Println("Mine         :", mine[:3])
	fmt.Println("Original nums:", nums[:3])
}
