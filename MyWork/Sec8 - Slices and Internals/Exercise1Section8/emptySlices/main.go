/*
*	Title:				Exercise1Section8/emptySlices/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-11
*	Version:			1.1
*	Description:		Take the code from the previous exercise, nilSlices, and make them empty slices
*						instead.
 */

package main

import "fmt"

func main() {
	names := []string{}
	distances := []int{}
	data := []uint8{}
	ratios := []float64{}
	status := []bool{}

	fmt.Printf("name: \t\t%T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: \t%T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: \t\t%T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: \t%T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("status: \t%T %d %t\n", status, len(status), status == nil)
}
