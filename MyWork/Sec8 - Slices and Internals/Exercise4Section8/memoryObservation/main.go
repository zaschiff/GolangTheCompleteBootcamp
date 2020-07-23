/*
*	title:				Exercise4Section8/memoryObservation/main.go
*	course (optional):	GO (Golang): The Complete Bootcamp By Jose Portilla and Inanc Gumus - Udemy Online
*	program author:		zach schiff
*	date:				2020-07-23
*	version:			1.0
*	description:		This is a simple memory obsrvation exercise to see the memory size when making an
*						array declaration
 */

package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const size = 1e7

func main() {
	// don't worry about this code.
	// it stops the garbage collector: prevents cleaning up the memory.
	// see the link if you're curious:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	// run the program to see the initial memory usage.
	report("initial memory usage")

	// 1. allocate an array with 10 million int elements
	//    the array's size will be equal to ~80MB
	//    hint: use the `size` constant above.
	var myArray [size]int

	// 2. print the memory usage (use the report func).
	report("After array allocation.")

	// 3. copy the array to a new array.
	newArray := myArray

	// 4. print the memory usage
	report("After copying initial array.")

	// 5. pass the array to the passArray function
	passArray(myArray)

	// 6. convert one of the arrays to a slice
	slice1 := myArray[:]

	// 7. slice only the first 1000 elements of the array
	slice2 := myArray[:100]

	// 8. slice only the elements of the array between 1000 and 10000
	slice3 := myArray[1000:10000]

	// 9. print the memory usage (report func)
	report("After slicing the array.")

	// 10. pass the one of the slices to the passSlice function
	passSlice(slice1)

	// 11. print the sizes of the arrays and slices
	//     hint: use the unsafe.Sizeof function
	//     see more here: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Println()
	fmt.Printf("Array's size : %d bytes.\n", unsafe.Sizeof(myArray))
	fmt.Printf("Array2's size: %d bytes.\n", unsafe.Sizeof(newArray))
	fmt.Printf("Slice1's size: %d bytes.\n", unsafe.Sizeof(slice1))
	fmt.Printf("Slice2's size: %d bytes.\n", unsafe.Sizeof(slice2))
	fmt.Printf("Slice3's size: %d bytes.\n", unsafe.Sizeof(slice3))
}

// passes [size]int array — about 80MB!
//
// observe that passing an array to a function (or assigning it to a variable)
// affects the memory usage dramatically
func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

// only passes 24-bytes of slice header
//
// observe that passing a slice doesn't affect the memory usage
func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

// reports the current memory usage
// don't worry about this code
func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
