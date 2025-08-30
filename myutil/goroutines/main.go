package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("Hello again from Goroutine after 2 seconds")
}

func sayHi() {
	fmt.Println("Hi Tarun from Goroutine")
	time.Sleep(1000 * time.Millisecond)
}

func main() {
	fmt.Println("Learning Goroutines in Go")

	//sequential execution check the output
	// sayHello()
	// sayHi()

	// if i make one function as goroutine then it will execute concurrently
	// go sayHello()
	// sayHi()

	// wait for goroutines to finish
	//time.Sleep(3000 * time.Millisecond) // this time is the wait time that main function should wait
	//like what happens is we make sayHello as goroutine now the sayHi will execute concurrently, so first sayHi will execute and
	// for getting the output of function that we make as goroutine we need to wait for some time for its completion
	//check output without giving time.

	//now lets see if we make both the function as goroutine then how the output willl look like for getting the output form the
	//goroutine functions we need to wait in the main function as after both the function there is nothing to do in the main function
	//so it will terminate the function without waiting for goroutines to finish. to get the output from goroutine function we need to wait for them to complete

	go sayHello()
	go sayHi()

	time.Sleep(3000 * time.Millisecond) // this time is the wait time that main function should wait

}
