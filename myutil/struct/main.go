package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	house  int
	street string
	city   string
	state  string
}

type Contact struct {
	Email string
	Phone string
}

type Employee struct {
	person_details  Person
	address_details Address
	contact_details Contact
}

func main() {
	employee1 := Employee{
		person_details: Person{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		},
		address_details: Address{
			house:  123,
			street: "Main St",
			city:   "Springfield",
			state:  "IL",
		},
		contact_details: Contact{
			Email: "john@doe.com",
			Phone: "123-456-7890",
		},
	}

	fmt.Println("Employee Details:", employee1)
}
