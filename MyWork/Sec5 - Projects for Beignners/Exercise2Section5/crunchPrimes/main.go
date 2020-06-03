/*
*	Title:				Exercise2Section5/crunchPrimes/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-02
*	Version:			1.0
*	Description:		gather a list of numbers from the commnad line, ignoring non-numbers,
*						print out only the prime numbers.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
main:
	for _, n := range os.Args[1:] {
		num, err := strconv.Atoi(n)
		if err != nil {
			continue
		}

		switch {
		case num == 2, num == 3:
			fmt.Printf("%d ", num)
		case num == 1, num%2 == 0, num%3 == 0:
			continue
		default:
			i := 5
			w := 2
			for (i * i) <= num {
				if num%i == 0 {
					continue main
				}
			}
			i += w
			w = 6 - w
			fmt.Printf("%d ", num)
		}
	}
}
