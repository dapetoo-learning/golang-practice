package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName   string
	LastName    string
	Age         int
	PhoneNumber int
}

func main() {

	var p1 Person
	p1 = Person{
		FirstName:   "John",
		LastName:    "Doe",
		Age:         30,
		PhoneNumber: 1234567890,
	}

	var p2 Person
	p2 = Person{
		FirstName:   "Jane",
		LastName:    "Doe",
		Age:         25,
		PhoneNumber: 1234567890,
	}

	b, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))
	b, err = json.Marshal(p2)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(b))

}
