package main

import (
	"fmt"
	"os"
)

func main() {

	name := os.Args[1]
	fmt.Println("Hello great", name+"!")

	name, age := "Gandalf", 2019

	fmt.Println("My name is", name)
	fmt.Println("My age is", age)
	fmt.Println("BTW, ou shall apss")
}
