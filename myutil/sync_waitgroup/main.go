package main

import (
	"fmt"
	"sync"
)

func callMe(i int, wg *sync.WaitGroup) {
	defer wg.Done() //this we use to tell the wait group that this goroutine is done, we can only write wg.DOne in last before completiojn of function, we are doing this to tell the sync wait group that how many tasks are done
	fmt.Printf("Goroutine started for %d\n", i)
	fmt.Printf("Goroutine ended for %d\n", i)
}

func main() {

	//now we want ki our main function should wait to execute itself till the time all the goroutines are not
	//completed. for that we have syncwait group
	var wg sync.WaitGroup //now this is how we declare the wait group

	for i := 1; i <= 3; i++ {
		wg.Add(1) //this we are using to tell the wg that how many goroutines are added to track or check
		go callMe(i, &wg)
	}

	wg.Wait() //this we use to tell the main function to wait till the time all the goroutines are not completed

	fmt.Println("Main function ended")
}
