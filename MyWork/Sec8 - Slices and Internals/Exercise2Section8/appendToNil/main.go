/*
*	Title:				Exercise2Section8/appendToNil/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-16
*	Version:			1.0
*	Description:		append to nil slices.
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	var pizza []string
	var departure []time.Time
	var gradYears []int
	var lights []bool

	pizza = append(pizza, "pepperoni", "olives", "pineapple")
	departure = append(departure, time.Now(), time.Now().Add(time.Hour*24), time.Now().Add(time.Hour*48))
	gradYears = append(gradYears, 1998, 2005, 2018)
	lights = append(lights, true, false, true)

	fmt.Printf("Pizza:   %s\n", pizza)
	fmt.Printf("Departure:   %s\n", departure)
	fmt.Printf("Graduation Years: %d\n", gradYears)
	fmt.Printf("lights:   %t\n", lights)
}
