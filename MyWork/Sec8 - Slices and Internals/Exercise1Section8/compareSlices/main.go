/*
*	Title:				Exercise1Section8/compareSlices/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-12
*	Version:			1.0
*	Description:		take the variables created in the project folder, split the strings and compare
*						the resulting slice to the one that was created.
 */

package main

import (
	"fmt"
	"strings"
)

func main() {

	namesA := "Da Vinci, Wozniak, Carmack"
	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}

	namesASlice := strings.Split(namesA, ", ")

	compared := false

	if len(namesASlice) == len(namesB) {
		compared = true
	}

	for _, s1 := range namesB {
		for _, s2 := range namesASlice {
			if s1 == s2 {
				compared = true
			}
		}
	}

	fmt.Printf("Are the slices equal: %t\n", compared)
}
