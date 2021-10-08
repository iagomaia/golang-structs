package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

func (p person) printName() {
	fmt.Printf("Name: %s %s\n", p.firstName, p.lastName)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
