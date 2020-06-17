/*
*	Title:				Exercise2Section8/appendSort/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-16
*	Version:			1.0
*	Description:		get numbers from the command line. Convert them form the
*						string slice to an int slice and sort them.
 */

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var nums []int

	if len(os.Args) == 1 {
		fmt.Println("Provide a few numbers.")
		return
	}

numsConvert:
	for _, n := range os.Args {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue numsConvert
		}
		nums = append(nums, num)
	}
	sort.Ints(nums)
	fmt.Println(nums)
}
