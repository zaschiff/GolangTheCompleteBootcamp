/*
*	Title:				Exercise2Section8/parseData/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-16
*	Version:			1.0
*	Description:		Take a constant of data. Parse the data, and pretty print it to the command line
 */

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	header = "Location,Size,Beds,Baths,Price"
	data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

	separator = ","
)

func main() {

	var location []string
	var size []int
	var beds []int
	var baths []int
	var price []int

	fmt.Println()
	rows := strings.Split(data, "\n")
	columNames := strings.Split(header, separator)
	for _, cn := range columNames {
		fmt.Printf("%-15s", cn)
	}
	fmt.Println()
	fmt.Println(strings.Repeat("=", 15*len(columNames)))

	for _, r := range rows {
		dSlice := strings.Split(r, separator)
		for i, d := range dSlice {
			switch i {
			case 0:
				location = append(location, d)
			case 1:
				num, err := strconv.Atoi(d)
				if err != nil {
					fmt.Println("Error: %d is not a number\n", d)
					return
				}
				size = append(size, num)
			case 2:
				num, err := strconv.Atoi(d)
				if err != nil {
					fmt.Println("Error: %d is not a number\n", d)
					return
				}
				beds = append(beds, num)
			case 3:
				num, err := strconv.Atoi(d)
				if err != nil {
					fmt.Println("Error: %d is not a number\n", d)
					return
				}
				baths = append(baths, num)
			case 4:
				num, err := strconv.Atoi(d)
				if err != nil {
					fmt.Println("Error: %d is not a number\n", d)
					return
				}
				price = append(price, num)
			}
		}
	}

	for i := 0; i < len(rows); i++ {
		fmt.Printf("%-14s %-14d %-14d %-14d %-14d\n", location[i], size[i], beds[i], baths[i], price[i])
	}

	fmt.Println()
}
