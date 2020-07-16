/*
*	Title:				Exercise3Seciton8/numberSlice/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-18
*	Version:			1.0
*	Description:		using a prewritten declaration. slice
*						the array after converting it from string to ints.
*						Comments will describe the slicing structure.
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// prewritten data string
	data := "2 4 6 1 3 5"
	var nums []int

numbers:
	for _, n := range strings.Split(data, " ") {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue numbers
		}
		nums = append(nums, num)
	}

	fmt.Printf("%-12s %d\n", "numbers:", nums)
	fmt.Printf("%-12s %d\n", "evens:", nums[:3])
	fmt.Printf("%-12s %d\n", "odds:", nums[3:])
	fmt.Printf("%-12s %d\n", "middle", nums[2:4])
	fmt.Printf("%-12s %d\n", "First 2:", nums[:2])
	fmt.Printf("%-12s %d\n", "Last 2:", nums[len(nums)-2:])
	fmt.Printf("%-12s %d\n", "Last Even:", nums[2:3])
	fmt.Printf("%-12s %d\n", "Odds Last 2:", nums[4:])
}
