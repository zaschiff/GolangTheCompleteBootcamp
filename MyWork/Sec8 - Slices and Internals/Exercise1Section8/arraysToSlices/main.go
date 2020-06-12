/*
*	Title:				Exercise1Section8/arraysToSlices/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-12
*	Version:			1.0
*	Description:		change the commented code from arrays to slices
 */

package main

import "fmt"

func main() {
	// names := [3]string{"Einstein", "Tesla", "Shepard"}
	// distances := [...]int{50, 40, 75, 30, 125}
	// data := [5]byte{'H', 'E', 'L', 'L', 'O'}
	// ratios := [1]float64{3.14145}
	// alives := [...]bool{true, false, true, false}
	// zero := [0]byte{}

	names := []string{"Einstein", "Tesla", "Shepard"}
	distances := []int{50, 40, 75, 30, 125}
	data := []byte{'H', 'E', 'L', 'L', 'O'}
	ratios := []float64{3.14145}
	alives := []bool{true, false, true, false}
	zero := []byte{}

	fmt.Printf("names    : %[1]T %[1]q\n", names)
	fmt.Printf("distances: %[1]T %[1]d\n", distances)
	fmt.Printf("data     : %[1]T %[1]d\n", data)
	fmt.Printf("ratios   : %[1]T %.2[1]f\n", ratios)
	fmt.Printf("alives   : %[1]T %[1]t\n", alives)
	fmt.Printf("zero     : %[1]T %[1]d\n", zero)
}
