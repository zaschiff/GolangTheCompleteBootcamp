/*
*	Title:				Exercise1Section6/getSetArray/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.2
*	Description:		we are taking the array declaration from the previous
*						problem, empty array, and setting the individual
*						elements.
 */

package main

import "fmt"

func main() {
	name := [...]string{"Kyle", "Alexa", "Jack"}
	distances := [...]int{10, 2, 4, 80, 25}
	data := [...]byte{2, 3, 4, 5, 6}
	ratio := [...]float32{1.2}
	status := [...]bool{true, true, false, false}
	zero := [...]byte{}

	// for i := 0; i < len(name); i++ {
	// 	fmt.Println(name[i])
	// }

	// for i := 0; i < len(distances); i++ {
	// 	fmt.Println(distances[i])
	// }

	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i])
	// }

	// for i := 0; i < len(ratio); i++ {
	// 	fmt.Println(ratio[i])
	// }

	// for i := 0; i < len(status); i++ {
	// 	fmt.Println(status[i])
	// }

	// for i := 0; i < len(zero); i++ {
	// 	fmt.Println(zero[i])
	// }

	for _, v := range name {
		fmt.Println(v)
	}

	for _, v := range distances {
		fmt.Println(v)
	}

	for _, v := range data {
		fmt.Println(v)
	}

	for _, v := range ratio {
		fmt.Println(v)
	}

	for _, v := range status {
		fmt.Println(v)
	}

	for _, v := range zero {
		fmt.Println(v)
	}

}
