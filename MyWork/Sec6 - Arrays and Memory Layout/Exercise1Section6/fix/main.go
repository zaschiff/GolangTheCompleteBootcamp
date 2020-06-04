/*
*	Title:				Exercise1Section6/fix/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.1
*	Description:		the bellow code must be fixed
 */

package main

import "fmt"

func main() {
	// var names [3]string = [3]string{
	// 	"Einstein" "Shepard"
	// 	"Tesla"
	// }

	// var books [5]string = [5]string{
	// 	"Kafka's Revenge",
	// 	"Stay Golden",
	// 	"",
	// 	"",
	// 	""
	// }

	// fmt.Printf("%q\n", names)
	// fmt.Printf("%q\n", books)

	names := [3]string{
		"Einstein",
		"Shepard",
		"Tesla",
	}

	books := [5]string{
		"Kafka's Revenge",
		"Stay Golden",
	}

	fmt.Printf("%q\n", names)
	fmt.Printf("%q\n", books)
}
