/*
*	Title:				Exercise1Section8/assignSlices/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-11
*	Version:			1.2
*	Description:		Take the code from the previous exercise, nilSlices, and assign values using a
*						slice literal.
 */

package main

import "fmt"

func main() {
	names := []string{"Zach", "Alexa", "Kyle"}
	distances := []int{1, 2, 3, 4, 5}
	data := []uint8{1, 2, 3, 4, 5}
	ratios := []float64{1.2, 2.3}
	status := []bool{true, true, false, false}

	fmt.Printf("name: \t\t%T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: \t%T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: \t\t%T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: \t%T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("status: \t%T %d %t\n", status, len(status), status == nil)

	if len(distances) == len(data) {
		fmt.Println("The length of the distances and data slices are the same.")
	}
}
