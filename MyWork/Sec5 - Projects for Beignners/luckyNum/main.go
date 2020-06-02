/*
*	Title:				luckyNum/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-05-30
*	Version:			1.0
*	Description:		a simple lucky number guessing game. Game is played at the command line.
 */

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns int    = 5
	guessMax int    = 10
	usage    string = "Welcome to the Lucky Number Game! \U0001F340\n\n" +
		`The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	// this is the third problem. I have set two variables to be
	// created and assigned. This allows the player to enter two
	// numbers into the game.
	guess1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	var guess2 int
	if len(os.Args) == 3 {
		guess2, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	for turn := 0; turn <= maxTurns; turn++ {
		n := rand.Intn(guessMax + 1)

		// This is porblem # for the luck number challenge.
		// This meesage will only print hsould the user get it on the first try.
		if (n == guess1 || n == guess2) && turn == 0 {
			fmt.Println("\U0001f631 WOW you guessed it on the first try!!! Your really good!!")
			return
		} else if n == guess1 || n == guess2 {

			// this is the second problem in the luck number challenge.
			// The rand ackage ill genereate a number from 0 to 5 and
			// print the corresponding message.
			msg := rand.Intn(5)

			switch msg {
			case 0:
				fmt.Println("\U0001F389 You won!!")
			case 1:
				fmt.Println("Do you read minds? Cause you read mine!")
			case 2:
				fmt.Println("How did you know?!?")
			case 3:
				fmt.Println("Good job! wanna play again?")
			case 4:
				fmt.Println("Did you look in my cpu? Cause thats cheating!")
			}
			return
		}
	}

	fmt.Println("\U0001F480 You lost....Try again?")
}
