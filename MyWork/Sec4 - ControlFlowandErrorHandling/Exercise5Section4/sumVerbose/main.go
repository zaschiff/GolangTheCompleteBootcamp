/*
*	Title:				Exercise5Section4/sumVerbose/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-30
*	Version:			1.0
*	Description:		Take the last problem and print out the sum along with the actual addition.
 */

package main

import "fmt"

func main() {
	var sum int
	for i := 1; i <= 10; i++ {
		if i == 10 {
			sum += i
			fmt.Printf("%d = %d\n", i, sum)
		} else {
			sum += i
			fmt.Printf("%d + ", i)
		}
	}
}
