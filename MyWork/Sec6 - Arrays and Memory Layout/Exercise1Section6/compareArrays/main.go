/*
*	Title:				Exercise1Section6/compareArrays/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.1
*	Description:		fix the commented code
 */

package main

import "fmt"

func main() {
	// week := [...]string{"Monday", "Tuesday"}
	// wend := [4]string{"Saturday", "Sunday"}

	// fmt.Println(week != wend)

	week := [4]string{"Monday", "Tuesday"}
	wend := [4]string{"Saturday", "Sunday"}

	fmt.Println(week != wend)

	// evens := [...]int{2, 4, 6, 8, 10}
	// evens2 := [...]int32{2, 4, 6, 8, 10}

	// fmt.Println(evens == evens2)
	evens := [...]int32{2, 4, 6, 8, 10}
	evens2 := [...]int32{2, 4, 6, 8, 10}

	fmt.Println(evens == evens2)

	// Use     : uint8 for one of the arrays instead of byte here.
	// Remember: Aliased types are the same types.
	image := [5]byte{'h', 'i'}
	var data [5]byte

	fmt.Println(data == image)
}
