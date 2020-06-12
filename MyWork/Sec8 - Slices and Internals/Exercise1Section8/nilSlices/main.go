/*
*	Title:				Exercise1Section8/nilSlices/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-11
*	Version:			1.0
*	Description:		A small exercise in declaring slices with nil values.
 */

package main

import "fmt"

func main() {
	var names []string
	var distances []int
	var data []uint8
	var ratios []float64
	var status []bool

	fmt.Printf("name: \t\t%T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: \t%T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data: \t\t%T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios: \t%T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("status: \t%T %d %t\n", status, len(status), status == nil)
}
