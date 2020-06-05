/*
*	Title:				Exercise2Section6/currencyConverter/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-04
*	Version:			1.0
*	Description:		a simple program to take in an argument from the command
*						line, assume it a USD amount, and convert it to all
*						available currencies in a stored array. The array should
*						use keyed values.
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	eur = iota
	gbp
	jpy
)

func main() {

	currencyRatio := [3]float64{
		eur: .88,
		gbp: .78,
		jpy: 113.02,
	}

	if len(os.Args) != 2 {
		fmt.Println("Please enter a dollar amount to convert.")
		return
	}

	amt, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Printf("%q is not a valid dollar amount.", os.Args[1])
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amt, amt*currencyRatio[eur])
	fmt.Printf("%.2f USD is %.2f GBP\n", amt, amt*currencyRatio[gbp])
	fmt.Printf("%.2f USD is %.2f JPY\n", amt, amt*currencyRatio[jpy])
}
