/*
*	Title:				Exercise5Section5/infiniteKill/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-30
*	Version:			1.0
*	Description:		A infinite loop that will run until interupted by the keyboard.
*						While running it will print a random symbol from a group of 4 symbols
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var c string
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for {
		switch r.Intn(4) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}

		fmt.Printf("\r%s Please wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}
