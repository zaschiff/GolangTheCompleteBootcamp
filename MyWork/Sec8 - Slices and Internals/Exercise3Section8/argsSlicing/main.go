/*
*	Title:				Exercise3Section8/argsSlicing/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-19
*	Version:			1.0
*	Description:		Using arguments, return the request sliced value form the
*						original slice.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// slice literal supplied from cours work
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}
	var (
		start int
		end   int
		err   error
	)

	if len(os.Args) == 1 || len(os.Args) > 3 {
		fmt.Println(ships, "\n\nProvide only the [starting] and [stopping] positions.")
		fmt.Println("\nError: we expect one or two arguments.")
		return
	}

	if len(os.Args) == 3 {
		start, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: Not a valid starting position.")
			return
		}
		end, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Not a valid ending position.")
			return
		}
	} else if len(os.Args) == 2 {
		start, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error: Not a valid starting position.")
			return
		}
	}

	switch {
	case start < 0:
		fallthrough
	case end < 0:
		fallthrough
	case start > len(ships) || end > len(ships):
		fallthrough
	case start > end && end != 0:
		fmt.Println("Wrong Positions.")
		return
	}

	if len(os.Args) == 3 {
		fmt.Println(ships[start:end])
	} else {
		fmt.Println(ships[start:])
	}

}
