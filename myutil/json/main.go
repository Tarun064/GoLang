package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"person"`   // we are creating the object but then we are saying it that you are a json kind data
	Age     int    `json:"age"`      // as we know that in json keys are always in string, so we are telling them
	IsAdult bool   `json:"is_adult"` // how to store them (their name) with json tags
	//agr hmm yeah json and key name nhi denge toh bhi chlega bss frk itna hoga ki sbse pehle voh json
	// wala name dekhta hai agr voh nhi mila toh voh struct field name dekhta hai aur usse use krr lega

}

func main() {
	fmt.Println("learning Go")
	person := Person{Name: "Tarun", Age: 24, IsAdult: true}
	jsonData, err := json.Marshal(person) //marshal means encoding/convert krna in json
	//jsonData will hold slice of bytes, so later we have to encode it in string
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))

	//decoding (Unmarshalling)
	var personData Person
	err = json.Unmarshal(jsonData, &personData) //unmarshal means decoding/convert krna from json
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}
	fmt.Println("Decoded Person Struct:", personData)

}
