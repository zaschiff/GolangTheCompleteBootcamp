/*
*	Title:				Exercise5Section4/sum/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-30
*	Version:			1.0
*	Description:		A simple go progam to loop and coutn the numbers 1 to 10.
*						The exepected output should be 55
 */

package main

import "fmt"

func main() {
	var sum int

	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}
