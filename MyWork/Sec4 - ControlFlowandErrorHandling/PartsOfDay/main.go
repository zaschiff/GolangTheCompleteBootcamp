/*
*	Title:				PartsOfDay
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-28
*	Version:			1.0
*	Description:		a small program to display a message based on hours of the day
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	switch hr := time.Now().Hour(); {
	case hr > 18:
		fmt.Println("Good Night!")
	case hr > 15:
		fmt.Println("Good Evening!")
	case hr > 12:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Morning!")
	}

	// One Solution:
	// switch {
	// 		case hr >= 18:
	// 			fmt.Println("Good Evening!")
	// 		case hr >= 12:
	// 			fmt.Println("Good afternoon!")
	// 		case hr >= 6:
	// 			fmt.Println("Good Morning!")
	// 		default:
	// 			fmt.Println("Good Night!")
	// }
}
