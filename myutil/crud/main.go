package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// get request
func performGetRequest() {
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

func performPostRequest() {
	todo := Todo{
		UserId:    1,
		Title:     "Tarun Tiwari",
		Completed: true,
	}

	//now we have to convert our object into json
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data := string(jsonData)

	//for posting the string data we have to convert our data into reader
	dataReader := strings.NewReader(data)

	//now we can post our request the specific endpoint
	myUrl := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myUrl, "application/json", dataReader)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()

	finalData, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Data from response", string(finalData))
	fmt.Println("Response status:", res.Status)
}

func performUpdateMethod() {
	todo := Todo{
		UserId:    1,
		Title:     "Tarun Tiwari",
		Completed: true,
	}

	//now we have to convert our object into json
	jsonData, err := json.Marshal(todo)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	data := string(jsonData)

	//for posting the string data we have to convert our data into reader
	dataReader := strings.NewReader(data)

	//now we can post our request the specific endpoint
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myUrl, dataReader)

	if err != nil {
		fmt.Println("Error found,", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	//send the request
	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error found,", err)
		return
	}

	defer res.Body.Close()

	dataa, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Data from response", string(dataa))
	fmt.Println("Response status:", res.Status)

}

func performDeleteMethod() {
	myUrl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodDelete, myUrl, nil)

	if err != nil {
		fmt.Println("Error found,", err)
		return
	}

	//send the request
	client := http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Error found,", err)
		return
	}

	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)

}

func main() {
	fmt.Println("Learning CRUD opertionas")

	//performGetRequest()  //get request learnt how an we do gert the data from the server

	//performPostRequest() //post request learnt how we can post the data to the server

	//performUpdateMethod() //put/patch request learnt how we can update the data to the server

	performDeleteMethod() //delete request learnt how we can delete the data from the server
}
