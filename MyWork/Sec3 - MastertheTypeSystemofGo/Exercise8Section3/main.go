/*
*	Title:				Exercise8/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-17
*	Version:			1.0
*	Description:		This is the main.go file for the eighth exercise in section three of the course.
*						The code is for the exercise labeled Data Types Exercise. To run each code
*						separately, comment and uncomment accordingly. Each exercise problem is labled
*						by number.
*
*						Please make sure to either comment out the line under problem 10 or compile the
*						program before running. if you use the go run command you can pass arguments but it
*						gets confusing with the number of command line values.
 */

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// 1. Print the following calculations
	// sum of 50 and 25
	// difference of 50 and 15.5
	// product of 50 and 0.5
	// quotient of 50 and 0.5
	// remainder of 25 and 3
	// negation of 5 + 2
	fmt.Println(50 + 25)
	fmt.Println(50 - 15.5)
	fmt.Println(50 * .5)
	fmt.Println(50 / .5)
	fmt.Println(25 % 3)
	fmt.Println(-(5 + 2))

	// 2. Fix the following coed to print 2.5 rather than 2
	// 		x := 5 / 2
	// 		fmt.Println(x)
	x2 := 5. / 2
	fmt.Println(x2)

	// 3. fix the code to print the desired value
	// 10 + 5 - 5 -10
	// should print 20
	fmt.Println(10 + 5 - (5 - 10))
	// -10 + 0.5 - 1 + 5.5
	// should print -16
	fmt.Println((-10 + 0.5) - (1 + 5.5))
	// 5 + 10 * 2 - 5
	// should print -25
	fmt.Println(5 + 10*(2-5))
	// 0.5 * 2 - 1
	// should print 0.5
	fmt.Println(0.5 * (2 - 1))
	// 3 + 1 / 2 * 10 + 4
	// should print 24
	fmt.Println(((3+1)/2)*10 + 4)
	// 10 / 2 * 10 % 7
	// should print 15
	fmt.Println((10 / 2) * (10 % 7))
	// 100 / 5 /2
	// should print 40
	fmt.Println(100 / (5. / 2.))

	// 4. Increase and decrease the follwing variables by the desired amount
	// counter = 45 increment 5 times
	// factor = 0.5 decrement 2 times
	counter4, factor4 := 45, 0.5
	counter4 += 5
	factor4 -= 2
	fmt.Println(float64(counter4) * factor4)

	// 5. Manipulate a counter variable by the following instructions
	// simplest way to increase by 1
	var counter5 int
	counter5++
	fmt.Println(counter5)
	// simplest way to decrease by
	counter5--
	fmt.Println(counter5)
	//simplest way to increase by 5
	counter5 += 5
	fmt.Println(counter5)
	// simplest way to multiply by 10
	counter5 *= 10
	fmt.Println(counter5)
	// simplest way to divide by 5
	counter5 /= 5
	fmt.Println(counter5)

	// 6. Refactor the follwing code using only incdec and assignment
	// width, height := 10, 2
	// width = width + 1
	// width = width + height
	// width = width - 1
	// width = width - height
	// width = wdith * 20
	// width = width / 25
	// width = width % 5
	width6, height6 := 10, 2
	fmt.Println(width6, height6)
	width6++
	fmt.Println(width6, height6)
	width6 += height6
	fmt.Println(width6, height6)
	width6--
	fmt.Println(width6, height6)
	width6 -= height6
	fmt.Println(width6, height6)
	width6 *= 20
	fmt.Println(width6, height6)
	width6 /= 25
	fmt.Println(width6, height6)
	width6 %= 5
	fmt.Println(width6, height6)

	// 7. Calculate the radius of a circle with the given radius of 10
	//		Use math.Pi or pi
	var (
		radius7 = 10.
		area7   float64
	)
	area7 = math.Pi * math.Pow(radius7, 2)
	fmt.Printf("radius: %.0f -> area: %.2f\n", radius7, area7)

	// 8. Using a command line argument, assuming it's a radius, calculate
	// 		the surface area of a sphere using the formula a=4*pi*r^2
	// ****** Uncomment the code below and comment the code in problem 9 to run ******
	// var (
	// 	radius8, area8 float64
	// )
	// arg := os.Args[1]
	// radius8, _ = strconv.ParseFloat(arg, 64)
	// area8 = 4 * math.Pi * math.Pow(radius8, 2)
	// fmt.Printf("radius: %.0f -> area: %.2f\n", radius8, area8)

	// 9. Using a command line argument, assuming it's radius, calculate
	// 		the volume of a sphere
	// ****** Uncomment the code below and comment the code in problem 8 to run ******
	var (
		radius9, volume9 float64
	)
	arg := os.Args[1]
	radius9, _ = strconv.ParseFloat(arg, 64)
	volume9 = (4. / 3.) * math.Pi * (math.Pow(radius9, 3))
	fmt.Printf("radius: %.1f -> volume: %.2f\n", radius9, volume9)
}
