/*
*	Title:				Sec7 - Retro Clock/ main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-05
*	Version:			1.0
*	Description:		using everyhting learned to this point create a retro
*						style clock to be printed on the commend line every second.
 */

// █░

package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	for shift := 0; ; shift++ {
		screen.Clear()
		screen.MoveTopLeft()

		t := time.Now()
		h := t.Hour()
		m := t.Minute()
		s := t.Second()
		//ssec := t.Nanosecond() / 1e8

		//a := s%10==0

		clock := [...]placeholder{
			digits[h/10], digits[h%10],
			colon,
			digits[m/10], digits[m%10],
			colon,
			digits[s/10], digits[s%10],
			/*
				dot,
				digits[ssec%10],
			*/
		}

		for line := range clock[0] {
			l := len(clock)

			start, end := shift%l, l
			/*
				if a {
					clock = alarm
				}
			*/

			if shift%(l*2) >= l {
				start, end = 0, start
			}

			for j := 0; j < (l - end); j++ {
				fmt.Print("   ")
			}

			for index := start; index < end; index++ {
				next := clock[index][line]
				if clock[index] == colon && s%2 == 0 {
					next = "  "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
		screen.Clear()
		screen.MoveTopLeft()
	}
}
