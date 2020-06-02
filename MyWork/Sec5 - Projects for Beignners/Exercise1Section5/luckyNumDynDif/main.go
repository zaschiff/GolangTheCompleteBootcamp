/*
*	Title:				luckyNum/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-02
*	Version:			1.3
*	Description:		a simple lucky number guessing game. Game is played at the command line.
*						Now has a verbose flag option. The game now allows the user to enter in a
*						number, and should that number be greater than 10, the number will become
*						the number generation limit for the computer. Game no includes dynamic
*						difficulty by incresing the number of turns the computer can take.
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

	var (
		verbose  bool
		guess    int
		err      error
		guessMax int = 10
	)

	if len(os.Args) == 3 && os.Args[1] == "-v" {
		verbose = true
		guess, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
	} else {
		guess, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Not a number.")
			return
		}
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	} else if guess > 10 {
		guessMax = guess
	}

	for turn := (maxTurns + guess/4); turn > 0; turn-- {
		n := rand.Intn(guessMax + 1)

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			fmt.Println("\U0001f631 WOW You guessed it! Your really good!!")
			return
		}
	}

	fmt.Println("\U0001F480 You lost....Try again?")
}
