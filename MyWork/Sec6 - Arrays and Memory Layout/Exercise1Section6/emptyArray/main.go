/*
*	Title:				Exercise1Section5/emptyArray/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		A small program exercise to work with creating arrays.
 */

package main

import "fmt"

func main() {
	var (
		name      [3]string
		distances [5]int
		data      [5]byte
		ratios    [1]float32
		status    [4]bool
		zero      [0]byte
	)

	fmt.Printf("names    : %#v\n", name)
	fmt.Printf("distances: %#v\n", distances)
	fmt.Printf("data     : %#v\n", data)
	fmt.Printf("ratio    : %#v\n", ratios)
	fmt.Printf("alive    : %#v\n", status)
	fmt.Printf("zero     : %#v\n", zero)

	fmt.Println()

	fmt.Printf("name    : %T\n", name)
	fmt.Printf("distance: %T\n", distances)
	fmt.Printf("data    : %T\n", data)
	fmt.Printf("ratio   : %T\n", ratios)
	fmt.Printf("alive   : %T\n", status)
	fmt.Printf("zero    : %T\n", zero)

	fmt.Println()

	fmt.Printf("name     : %q\n", name)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data     : %d\n", data)
	fmt.Printf("ratio    : %.2f\n", ratios)
	fmt.Printf("alive    : %t\n", status)
	fmt.Printf("zero     : %d\n", zero)
}
