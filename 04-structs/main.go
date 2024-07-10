package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	ringo := person{
		firstName: "Ringo",
		lastName:  "Starr",
		contactInfo: contactInfo{
			email:   "ringo.starr@beatles.com",
			zipCode: 10000,
		},
	}

	//ringo.print()

	ringo.updateName("Ringin")
	ringo.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
