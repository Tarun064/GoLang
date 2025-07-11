package main

import (
	"fmt"
	"start/myutil" // Importing the myutil package
)

func main() {
	fmt.Println("Hello, World!")
	myutil.PrintMessage("This is a message from myutil package.")

	var x int = 10
	fmt.Println("The value of x is:", x)

	var y = "Hello, Go!"
	fmt.Println("The value of y is:", y)
}
