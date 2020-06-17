/*
*	Title:				Exercise2Section8/parseDataAverages/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author:		Zach Schiff
*	Date:				2020-06-16
*	Version:			1.1
*	Description:		Take a constant of data. Parse the data, and pretty print it to the command line
*						followed by the averages.
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

	var (
		location []string
		size     []int
		beds     []int
		baths    []int
		price    []int

		sizeSum  float64
		bedSum   float64
		bathSum  float64
		priceSum float64
	)

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
		fmt.Printf("%-15s %-15d %-15d %-15d %-15d\n", location[i], size[i], beds[i], baths[i], price[i])
		sizeSum += float64(size[i])
		bedSum += float64(beds[i])
		bathSum += float64(baths[i])
		priceSum += float64(price[i])
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 15*len(columNames)))
	fmt.Printf("%-15s %-15.2f %-15.2f %-15.2f %-15.2f\n", "", (sizeSum / 4), (bedSum / 4), (bathSum / 4), (priceSum / 4))
	fmt.Println()
}
