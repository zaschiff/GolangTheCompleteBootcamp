/*
*	title:				Exercise4Section8/sortMidBackArray/main.go
*	course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	program author:		zach schiff
*	date:				2020-07-22
*	version:			1.0
*	description:		Sort the middle three itmes of the backing array
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}
	fmt.Println("Original:", items)

	// My code to soprt the middle three items
	m := len(items) / 2
	midItems := items[m-1 : m+2]
	sort.Strings(midItems)

	fmt.Println()
	fmt.Println("Sorted  :", items)
}
