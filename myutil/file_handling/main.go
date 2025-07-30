package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//   1.   this is for creating a file and writing to it
	/*
		file, err := os.Create("sample.txt") // Create a new file named sample.txt
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close() // Ensure the file is closed after all operations are done
		content := "Hello World, this is a sample file."
		_, errors := io.WriteString(file, content) // Write content to the file amd instead of _ we can use byte bcz it tells us how much byte we are taking
		if errors != nil {
			fmt.Println("Error writing to file:", errors)
			return
		}
	*/

	//   2.    now lets see how can we read the fileby two methods buffer and function
	/*
		buffer := make([]byte, 1024) //we create the slice of bytes with a capacity of 1024

		for {
			//reading the file by buffer
			n, err := file.Read(buffer) // Read the file into the buffer
			if err == io.EOF {
				break // Exit the loop if we reach the end of the file
			}
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			fmt.Println(string(buffer[:n])) // Print the content read from 0 to n bytes
		}
	*/
	// 3. reading the file by function

	content, err := ioutil.ReadFile("sample.txt") // Read the entire file content into memory
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(content)) // Print the content of the file as a string

}
