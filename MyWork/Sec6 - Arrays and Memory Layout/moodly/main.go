/*
*	Title:				moodly/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-03
*	Version:			1.0
*	Description:		a small coding challenge to show use of array's. The different moods should be in an array.
*						and should be randomly selected
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter your name.")
		return
	}

	name := os.Args[1]
	moods := [...]string{
		"sad",
		"bad",
		"terrible",
		"happy",
		"good",
		"awesome",
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println(name + " is feeling " + moods[rand.Intn(len(moods))])

}
