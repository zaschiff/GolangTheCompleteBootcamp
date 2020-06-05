/*
*	Title:				Exercise2Section6/averages/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		an average calculator based on a max of 5 numbers entered at the command line
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 6 {
		fmt.Println("Please give me some numbers.(maximum of 5")
		return
	}

	var (
		nums  [5]float64
		sum   float64
		total float64
	)

	for i := 1; i < len(os.Args); i++ {
		n, err := strconv.ParseFloat(os.Args[i], 64)
		if err != nil {
			continue
		}

		total++
		nums[i-1] = n
		sum += n
	}

	fmt.Println("Your Numbers: ", nums)
	fmt.Println("Average: ", sum/total)

}
