/*
*	Title:				Exercise5Section4/breakUp/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-30
*	Version:			1.0
*	Description:		Another version of evenOnly but an infinite loop that kills at the max value
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	errNoNum   string = "Please give me a min and max to use."
	noValidInt string = "Something is wrong. Probably wrong numbers."
	usage      string = "[min] [max] ==> min is the start of the loop max is the end.\n\t\t    Please make sure min is less than max."
)

func main() {
	var (
		sum int
		i   int
	)

	// check for the right amount of inputs
	if len(os.Args) != 3 {
		fmt.Println(errNoNum)
		fmt.Println(usage)
		return
	}

	// check to make sure first argument is correct
	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min > max {
		fmt.Println(noValidInt)
		fmt.Println(usage)
	}

	// sum the numbers verbose style
	i = min
	for {
		if i%2 != 0 {
			i++
			continue
		} else if i == max {
			sum += i
			fmt.Printf("%d = %d\n", i, sum)
			break
		} else {
			sum += i
			fmt.Printf("%d + ", i)
			i++
		}
	}
}
