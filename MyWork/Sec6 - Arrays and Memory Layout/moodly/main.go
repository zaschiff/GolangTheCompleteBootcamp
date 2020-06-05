/*
*	Title:				moodly/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-03
*	Version:			1.1
*	Description:		a small coding challenge to show use of array's. The different moods should be in an array.
*						and should be randomly selected. the code will be modified to support a second argument
*						from the command line.
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please enter your name.")
		return
	}

	name := os.Args[1]

	var side int
	if os.Args[2] != "positive" {
		side = 1
	}

	moods := [...][3]string{
		{"happy", "good", "awesome"},
		{"sad", "bad", "terrible"},
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(name + " is feeling " + moods[side][rand.Intn(len(moods[side]))])

}
