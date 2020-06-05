/*
*	Title:				Exercise2Section6/numberSort/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		Using a bubble sort algorithm and commnad line input,
*						sort the numbers.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 6 {
		fmt.Println("Please give me some numbers. (Maximum of 5 numbers)")
		return
	}

	var nums [5]float64

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			continue
		}
		nums[i-1] = n
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < (len(nums) - i - 1); j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}

		}
	}

	fmt.Println(nums)
}
