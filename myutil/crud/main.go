package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD opertionas")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close() //it will close the body after the main function is executed

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", res.StatusCode)
		return
	}

	//generic way of getting or showing the data
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// fmt.Println("Data", string(data))

	//iss upr waale kaam se hmme direct json data mol rha hai toh aage jaake agr mujhe usse struct mein store
	// krvana hai toh uske unmarshal vgrh krna pdega, toh number of lines of code bddh jaate, bss vhi same kaam
	//unmarshal vgrh vala newdecoder function directly krke de rha hai hmme

	//one more to directly getting the data
	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo) //we are directly decoding the data into the struct

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Todo:", todo)
}
