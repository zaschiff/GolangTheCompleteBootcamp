/*
*	Title:				Exercise9/main.go
*	Course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	Program Author: 	Zach Schiff
*	Date: 				2020-05-17
*	Version:			1.0
*	Description:		This is the main.go file for the ninth exercise in section three of the course.
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
)

func main() {
	// 1. Find the optimal types for the following data descriptions
	// an english letter
	var char byte = 'A'
	//a non english character
	var unicode rune = '£'
	// a four digit year
	var year uint16 = 2040
	// a month number
	var month uint8 = 12
	// speed of light
	// I disagree with the answer as a 32 bit signed
	//		value can hold up over 2 billion values.
	//		since the speed of light is the fastest speed
	//		in the world and will not be changing, the
	//		optimum would be the smallest opton available
	var light int32 = 670616629
	// angle of a circle
	var angle float32 = 45.3
	// pi
	var pi float64 = 3.141592653589793
	fmt.Println(char, unicode, year, month, light, angle, pi)

	// 2. fix teh type problems in the following code
	// var (
	// 	width  uint8
	// 	height uint16
	// )

	// width, height = 255, 265
	// width += 10
	// fmt.Printf("width: %d height: %d\n", width, height)
	var (
		width  uint16
		height uint16
	)

	width, height = 255, 265
	width += 10
	fmt.Printf("width: %d height: %d\n", width, height)

	// 3. Convert command line arguments to the various int sizes
	// val8, _ := strconv.ParseInt(os.Args[1], 10, 8)
	// fmt.Println("int8 value is:", int8(val8))

	// val16, _ := strconv.ParseInt(os.Args[2], 10, 16)
	// fmt.Println("int16 value is:", int16(val16))

	// val32, _ := strconv.ParseInt(os.Args[3], 10, 32)
	// fmt.Println("int32 value is:", int32(val32))

	// val64, _ := strconv.ParseInt(os.Args[4], 10, 64)
	// fmt.Println("int64 value is:", int64(val64))

	// valBin, _ := strconv.ParseInt(os.Args[5], 2, 8)
	// fmt.Printf("%s is: %d\n", os.Args[5], valBin)

	// 4. take a command line argument and multiply it by the t
	//		value declared below. You will need to convert it to int64
	// t, _ := time.ParseDuration("1h30m")

	// m, _ := strconv.ParseInt(os.Args[1], 10, 64)

	// t *= time.Duration(m)
	// fmt.Println(t)

	// 5. Create a new types called feet and meters.
	//		Get the feet value from the command line and convert it to meters
	// type (
	// 	Feet  float64
	// 	Meter float64
	// )

	// var (
	// 	feet  Feet
	// 	meter Meter
	// )

	// val, _ := strconv.ParseFloat(os.Args[1], 64)
	// feet = Feet(val)

	// meter = Meter(feet * .3048)
	// fmt.Printf("%g feet is %g meters\n", feet, meter)

	// 6. Write the correct conversions to get the conversion
	//		function work using the custom declared types and values
	type (
		Temperature float64
		Celsius     Temperature
		Fahrenheit  Temperature
	)

	var (
		celsius       Celsius     = 15.5
		fahr          Fahrenheit  = 59.9
		celsiusDegree Temperature = 10.5
		fahrDegree    Temperature = 2.5
		factor                    = 2.
	)

	celsius *= Celsius(celsiusDegree * Temperature(factor))
	fahr *= Fahrenheit(fahrDegree * Temperature(factor))

	fmt.Println(celsius, fahr)

}
